/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

// Deterministic JSON.stringify()
const stringify  = require('json-stringify-deterministic');
const sortKeysRecursive  = require('sort-keys-recursive');
const { Contract } = require('fabric-contract-api');

class AssetTransfer extends Contract {

    async InitLedger(ctx) {
        const assets = [
            {
                   DEALERID :  "DEALER_1", 
		   MSISDN   :  "+910000000001",
		   MPIN     :  1000, 
		   BALANCE  :  1000, 
		   STATUS   :  "ACTIVE",
		   DEALER   :  "DEALER1",
		 TRANSAMOUNT:  10,
		   TRANSTYPE:  "TO",
                   REMARKS  :  "ADD REMARKS HERE",
            },
            {
            
                   DEALERID    :   "DEALER_2",                                                     MSISDN  :    "+91000000002",                                                 MPIN   :      1001,                                                      BALANCE   :   1000,                                                            STATUS :    "ACTIVE",                                                        DEALER:    "DEALER2",
                   TRANSAMOUNT :    10,
                   TRANSTYPE   :    "TO",
            },
            {
               DEALERID        :"DEALER_3",
		    MSISDN     :"+910000000003", 
		    MPIN       :1003,
		    BALANCE    :1000, 
		    STATUS     :"ACTIVE",
		    DEALER     :"DEALER3",
		    TRANSAMOUNT:10,
		    TRANSTYPE  :"TO",
		    REMARKS    :"ADD REMARKS HERE",
            },
            {
                    DEALERID   : "DEALER_4",
		    MSISDN     :"+910000000004", 
		    MPIN       :1004, 
		    BALANCE    :1000, 
		    STATUS     :"ACTIVE",
		    DEALER     :"DEALER4",
		    TRANSAMOUNT:10,
		    TRANSTYPE  :"TO",
		    REMARKS    :"ADD REMARKS HERE",
            },
            {
               DEALERID     : "DEALER_5",
		 MSISDN     :"+910000000005", 
		 MPIN       :1005,
		 BALANCE    :1000,
		 STATUS     :"ACTIVE",
		 DEALER     :"DEALER5",
		 TRANSAMOUNT:10,
		 TRANSTYPE  :"TO",
		 REMARKS    :"ADD REMARKS HERE",
            },
            
        ];

        for (const asset of assets) {
            asset.docType = 'asset';
            // example of how to write to world state deterministically
            // use convetion of alphabetic order
            // we insert data in alphabetic order using 'json-stringify-deterministic' and 'sort-keys-recursive'
            // when retrieving data, in any lang, the order of data will be the same and consequently also the corresonding hash
            await ctx.stub.putState(asset.DEALERID, Buffer.from(stringify(sortKeysRecursive(asset))));
        }
    }

    // CreateAsset issues a new asset to the world state with given details.
    async CreateAsset(ctx, dealerid, msisdn, mpin, balance,status_,dealer,transamount,transtype,remarks) {
        const exists = await this.AssetExists(ctx, dealerid);
        if (exists) {
            throw new Error(`The asset ${dealerid} already exists`);
        }

        const asset = {
           
           DEALERID:        dealerid,                                                  MSISDN:          msisdn,                                                    BALANCE:         balance,                                                   STATUS:          status_,                                                   DEALER:          dealer,                                                    TRANSAMOUNT:     transamount,                                               TRANSTYPE:       transtype,                                                 REMARKS:         remarks,
        };
        // we insert data in alphabetic order using 'json-stringify-deterministic' and 'sort-keys-recursive'
        await ctx.stub.putState(dealerid, Buffer.from(stringify(sortKeysRecursive(asset))));
        return JSON.stringify(asset);
    }

    // ReadAsset returns the asset stored in the world state with given id.
    async ReadAsset(ctx, dealerid) {
        const assetJSON = await ctx.stub.getState(dealerid); // get the asset from chaincode state
        if (!assetJSON || assetJSON.length === 0) {
            throw new Error(`The asset ${dealerid} does not exist`);
        }
        return assetJSON.toString();
    }

    // UpdateAsset updates an existing asset in the world state with provided parameters.
    async UpdateAsset(ctx, dealerid, msisdn, balance, status_,dealer,transamount,transtype,remarks) {
        const exists = await this.AssetExists(ctx, dealerid);
        if (!exists) {
            throw new Error(`The asset ${dealerid} does not exist`);
        }

        // overwriting original asset with new asset
        const updatedAsset = {
           DEALERID:        dealerid,                                                  MSISDN:          msisdn,                                                    BALANCE:         balance,                                                   STATUS:          status_,                                                   DEALER:          dealer,                                                    TRANSAMOUNT:     transamount,                                               TRANSTYPE:       transtype,                                                 REMARKS:         remarks,
        };
        // we insert data in alphabetic order using 'json-stringify-deterministic' and 'sort-keys-recursive'
        return ctx.stub.putState(dealerid, Buffer.from(stringify(sortKeysRecursive(updatedAsset))));
    }

    // DeleteAsset deletes an given asset from the world state.
    async DeleteAsset(ctx, dealerid) {
        const exists = await this.AssetExists(ctx, dealerid);
        if (!exists) {
            throw new Error(`The asset ${dealerid} does not exist`);
        }
        return ctx.stub.deleteState(dealerid);
    }

    // AssetExists returns true when asset with given ID exists in world state.
    async AssetExists(ctx, dealerid) {
        const assetJSON = await ctx.stub.getState(dealerid);
        return assetJSON && assetJSON.length > 0;
    }

    // TransferAsset updates the owner field of asset with given id in the world state.
    async TransferAsset(ctx, dealerid, newDealer) {
        const assetString = await this.ReadAsset(ctx, dealerid);
        const asset = JSON.parse(assetString);
        const oldOwner = asset.DEALER;
        asset.DEALER = newDEALER;
        // we insert data in alphabetic order using 'json-stringify-deterministic' and 'sort-keys-recursive'
        await ctx.stub.putState(dealerid, Buffer.from(stringify(sortKeysRecursive(asset))));
        return oldDealer;
    }

    // GetAllAssets returns all assets found in the world state.
    async GetAllAssets(ctx) {
        const allResults = [];
        // range query with empty string for startKey and endKey does an open-ended query of all assets in the chaincode namespace.
        const iterator = await ctx.stub.getStateByRange('', '');
        let result = await iterator.next();
        while (!result.done) {
            const strValue = Buffer.from(result.value.value.toString()).toString('utf8');
            let record;
            try {
                record = JSON.parse(strValue);
            } catch (err) {
                console.log(err);
                record = strValue;
            }
            allResults.push(record);
            result = await iterator.next();
        }
        return JSON.stringify(allResults);
    }
}

module.exports = AssetTransfer;
