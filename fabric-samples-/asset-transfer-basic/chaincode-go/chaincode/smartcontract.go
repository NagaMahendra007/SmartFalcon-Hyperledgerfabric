package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	DEALERID       string    `json:"DEALERID"`
	MSISDN         string    `json:"MSISDN"`
	MPIN           int       `json:"MPIN"`
	BALANCE        int       `json:"BALANCE"`
	STATUS         string    `json:"STATUS"`
	DEALER         string    `json:"DEALER"`
	TRANSAMOUNT    int       `json:"TRANSAMOUNT"`
	TRANSTYPE      string    `json:"TRANSTYPE"`
	REMARKS        string    `json:"REMARKS"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{DEALERID: "DEALER_1", MSISDN:"+910000000001", MPIN:1000, BALANCE:1000, STATUS:"ACTIVE", DEALER:"DEALER1", TRANSAMOUNT:10, TRANSTYPE:"TO", REMARKS:"ADD REMARKS HERE"},
		  {DEALERID: "DEALER_2", MSISDN:"+91000000002", MPIN:1001, BALANCE:1000, STATUS:"ACTIVE", DEALER:"DEALER2", TRANSAMOUNT:10, TRANSTYPE:"TO", REMARKS:"ADD REMARKS HERE"},
  {DEALERID: "DEALER_3", MSISDN:"+910000000003", MPIN:1003, BALANCE:1000, STATUS:"ACTIVE", DEALER:"DEALER3", TRANSAMOUNT:10, TRANSTYPE:"TO", REMARKS:"ADD REMARKS HERE"},
		  {DEALERID: "DEALER_4", MSISDN:"+910000000004", MPIN:1004, BALANCE:1000, STATUS:"ACTIVE", DEALER:"DEALER4",TRANSAMOUNT:10,TRANSTYPE:"TO", REMARKS:"ADD REMARKS HERE"},
		  {DEALERID: "DEALER_5", MSISDN:"+910000000005", MPIN:1005, BALANCE:1000, STATUS:"ACTIVE", DEALER:"DEALER5", TRANSAMOUNT:10, TRANSTYPE:"TO", REMARKS:"ADD REMARKS HERE"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.DEALERID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerid string, msisdn string, mpin int, balance int, status_ string, dealer string, transamount int, transtype string, remarks string) error {
	exists, err := s.AssetExists(ctx, dealerid)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", dealerid)
	}

	asset := Asset{
		DEALERID:        dealerid,
		MSISDN:          msisdn,
		BALANCE:         balance,
		STATUS:          status_,
		DEALER:          dealer,
		TRANSAMOUNT:     transamount,
		TRANSTYPE:       transtype,
		REMARKS:         remarks,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerid, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, dealerid string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(dealerid)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", dealerid)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerid string, msisdn string, mpin int, balance int, status_ string, dealer string, transamount int, transtype string, remarks string) error {
	exists, err := s.AssetExists(ctx, dealerid)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", dealerid)
	}

	// overwriting original asset with new asset
	asset := Asset{
		DEALERID    :     dealerid,
		MSISDN      :     msisdn,
		BALANCE     :     balance,
		STATUS      :     status_,
		DEALER      :     dealer,
		TRANSAMOUNT :     transamount,
		TRANSTYPE   :     transtype,
		REMARKS     :     remarks,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerid, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, dealerid string) error {
	exists, err := s.AssetExists(ctx, dealerid)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", dealerid)
	}

	return ctx.GetStub().DelState(dealerid)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, dealerid string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(dealerid)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state, and returns the old owner.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, dealerid string, newDealer string) (string, error) {
	asset, err := s.ReadAsset(ctx, dealerid)
	if err != nil {
		return "", err
	}

	oldDealer := asset.DEALER
	asset.DEALER = newDealer

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(dealerid, assetJSON)
	if err != nil {
		return "", err
	}

	return oldDealer, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
