PROBLEM STATEMENT:(PROGRAMMING LANGUAGE : GO) 

Afinancial institution needs to implement a blockchain-based system to manage and track assets. 
The system should support creating assets, updating asset values, querying the world state to read assets, 
and retrieving asset transaction history. The assets represent accounts with specific attributes, 
such as DEALERID, MSISDN, MPIN, BALANCE, STATUS, TRANSAMOUNT, TRANSTYPE, and REMARKS.
The institution aims to ensure the security, transparency, and immutability of asset records, while also providing an efficient way to track and manage asset-related transactions and histories.
And i added an DEALER attribute for identification purpose 
INTRODUCTION: 
Hyperledger Fabric is an open source enterprise permissioned distributed ledger technology(DLT) platform
KEY CONCEPTS:
     -->SmartContract,Chaincode,Peer,Orderer Service, Pluggable Consensus Protocal,Permissioned Network,Transaction,Endorsement policy,Ledger,Channels etc

LEVEL 1:( SetUp HyperledgerFabric Test Network )
     -->After understanding the bussiness behind the Hyperledger fabric,I started install the prerequisites and setting up the development environment
        like GO,JQ,DOCKER,GIT,SoftHSM etc
     -->Next I  was cloned the fabric samples from repository hyperledger/fabric , CLI binary tools & APIs to support developing smart contracts(chaincode) in { GO ,JAVA,Node.js } 
     USING FABRIC TEST NETWORK:
         -->Deploy a test network by using scripts that are provided in fabric-samples- repository
         -->network.sh script used to create channel between the organizations
     --> deployCC subvcommand to install chaincodes
LEVEL 2:(Develop and test the smart contract for the above problem statement)
  
   -->fabric-samples-/assert-transfer-basic/chaincode-go/chaincode/smartcontract.go
        Have the code for the Smartcontract which is an logic used on ledger to read and write by using valid transactions
   --> Next using the commands Installation of chaincode package,Aprrove of chaincode definition, Commiting the chaincode  to  the definition of chaincode
   --> Invoking the chaincode and Upgrading it and identifiying using the Sequence number incremented by 2  

LEVEL 3:( Develop a rest api for invoking smart contract and deployed it into hyperledger fabric test network and creates a docker image for rest api)
     Fabric gateway  client APIs is a service which manages the following transaction steps:
     Evaluate a transaction proposal,Endorse a transaction proposal,Submit a transaction, wait for a commit status and recive chaincode events
    -->Run Fabric Application reference provides an way to Query and Updates the ledger
        1. Setting up a blockchain network
        2. Run the sample application to intercat with smart contracts
    -->creates a gRPC connection
    -->Query all assets, Create a new asset, Update an asset, Query the updated asset,Handle Transaction errors
      
