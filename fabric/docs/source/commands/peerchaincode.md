<!---
 File generated by help_docs.sh. DO NOT EDIT.
 Please make changes to preamble and postscript wrappers as appropriate.
 --->

# peer chaincode

The `peer chaincode` command allows users to invoke and query chaincode.

## Syntax

The `peer chaincode` command has the following subcommands:

  * invoke
  * query

The subcommands take a constructor flag (`-c` or `--ctor`) to pass arguments to a chaincode.
The value must be a JSON string that has either key 'Args' or 'Function' and 'Args'.
These keys are case-insensitive.

If the constructor JSON string only has the Args key, the key value is an array, where the
first array element is the target function to call, and the subsequent elements
are arguments of the function. If the JSON string has both 'Function' and
'Args', the value of Function is the target function to call, and the value of
Args is an array of arguments of the function. For instance,
`{"Args":["GetAllAssets"]}` is equivalent to
`{"Function":"GetAllAssets", "Args":[]}`.

Each peer chaincode subcommand is described together with its options in its own
section in this topic.

## Flags

Each `peer chaincode` subcommand has both a set of flags specific to an
individual subcommand, as well as a set of global flags that relate to all
`peer chaincode` subcommands. Not all subcommands would use these flags.
For instance, the `query` subcommand does not need the `--orderer` flag.

The individual flags are described with the relevant subcommand. The global
flags are

* `--cafile <string>`

  Path to file containing PEM-encoded trusted certificate(s) for the ordering
  endpoint

* `--certfile <string>`

  Path to file containing PEM-encoded X509 public key to use for mutual TLS
  communication with the orderer endpoint

* `--keyfile <string>`

  Path to file containing PEM-encoded private key to use for mutual TLS
  communication with the orderer endpoint

* `-o` or `--orderer <string>`

  Ordering service endpoint specified as `<hostname or IP address>:<port>`

* `--ordererTLSHostnameOverride <string>`

  The hostname override to use when validating the TLS connection to the orderer

* `--tls`

  Use TLS when communicating with the orderer endpoint

* `--transient <string>`

  Transient map of arguments in JSON encoding

Flags of type stringArray are to be repeated rather than concatenating their
values. For example, you will use `--peerAddresses localhost:9051
--peerAddresses localhost:7051` rather than `--peerAddresses "localhost:9051
localhost:7051"`.

## peer chaincode invoke
```
Invoke the specified chaincode. It will try to commit the endorsed transaction to the network.

Usage:
  peer chaincode invoke [flags]

Flags:
  -C, --channelID string               The channel on which this command should be executed
      --connectionProfile string       Connection profile that provides the necessary connection information for the network. Note: currently only supported for providing peer connection information
  -c, --ctor string                    Constructor message for the chaincode in JSON format (default "{}")
  -h, --help                           help for invoke
  -I, --isInit                         Is this invocation for init (useful for supporting legacy chaincodes in the new lifecycle)
  -n, --name string                    Name of the chaincode
      --peerAddresses stringArray      The addresses of the peers to connect to
      --tlsRootCertFiles stringArray   If TLS is enabled, the paths to the TLS root cert files of the peers to connect to. The order and number of certs specified should match the --peerAddresses flag
      --waitForEvent                   Whether to wait for the event from each peer's deliver filtered service signifying that the 'invoke' transaction has been committed successfully
      --waitForEventTimeout duration   Time to wait for the event from each peer's deliver filtered service signifying that the 'invoke' transaction has been committed successfully (default 30s)

Global Flags:
      --cafile string                       Path to file containing PEM-encoded trusted certificate(s) for the ordering endpoint
      --certfile string                     Path to file containing PEM-encoded X509 public key to use for mutual TLS communication with the orderer endpoint
      --clientauth                          Use mutual TLS when communicating with the orderer endpoint
      --connTimeout duration                Timeout for client to connect (default 3s)
      --keyfile string                      Path to file containing PEM-encoded private key to use for mutual TLS communication with the orderer endpoint
  -o, --orderer string                      Ordering service endpoint
      --ordererTLSHostnameOverride string   The hostname override to use when validating the TLS connection to the orderer
      --tls                                 Use TLS when communicating with the orderer endpoint
      --tlsHandshakeTimeShift duration      The amount of time to shift backwards for certificate expiration checks during TLS handshakes with the orderer endpoint
      --transient string                    Transient map of arguments in JSON encoding
```


## peer chaincode query
```
Get endorsed result of chaincode function call and print it. It won't generate transaction.

Usage:
  peer chaincode query [flags]

Flags:
  -C, --channelID string               The channel on which this command should be executed
      --connectionProfile string       Connection profile that provides the necessary connection information for the network. Note: currently only supported for providing peer connection information
  -c, --ctor string                    Constructor message for the chaincode in JSON format (default "{}")
  -h, --help                           help for query
  -x, --hex                            If true, output the query value byte array in hexadecimal. Incompatible with --raw
  -n, --name string                    Name of the chaincode
      --peerAddresses stringArray      The addresses of the peers to connect to
  -r, --raw                            If true, output the query value as raw bytes, otherwise format as a printable string
      --tlsRootCertFiles stringArray   If TLS is enabled, the paths to the TLS root cert files of the peers to connect to. The order and number of certs specified should match the --peerAddresses flag

Global Flags:
      --cafile string                       Path to file containing PEM-encoded trusted certificate(s) for the ordering endpoint
      --certfile string                     Path to file containing PEM-encoded X509 public key to use for mutual TLS communication with the orderer endpoint
      --clientauth                          Use mutual TLS when communicating with the orderer endpoint
      --connTimeout duration                Timeout for client to connect (default 3s)
      --keyfile string                      Path to file containing PEM-encoded private key to use for mutual TLS communication with the orderer endpoint
  -o, --orderer string                      Ordering service endpoint
      --ordererTLSHostnameOverride string   The hostname override to use when validating the TLS connection to the orderer
      --tls                                 Use TLS when communicating with the orderer endpoint
      --tlsHandshakeTimeShift duration      The amount of time to shift backwards for certificate expiration checks during TLS handshakes with the orderer endpoint
      --transient string                    Transient map of arguments in JSON encoding
```

## Example Usage

### peer chaincode invoke example

Here is an example of the `peer chaincode invoke` command:

  * Invoke the chaincode named `mycc` at version `1.0` on channel `mychannel`
    on `peer0.org1.example.com:7051` and `peer0.org2.example.com:9051` (the
    peers defined by `--peerAddresses`), requesting to move 10 units from
    variable `a` to variable `b`:

    ```
    peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc --peerAddresses peer0.org1.example.com:7051 --peerAddresses peer0.org2.example.com:9051 -c '{"Args":["invoke","a","b","10"]}'

    2018-02-22 16:34:27.069 UTC [chaincodeCmd] checkChaincodeCmdParams -> INFO 001 Using default escc
    2018-02-22 16:34:27.069 UTC [chaincodeCmd] checkChaincodeCmdParams -> INFO 002 Using default vscc
    .
    .
    .
    2018-02-22 16:34:27.106 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> DEBU 00a ESCC invoke result: version:1 response:<status:200 message:"OK" > payload:"\n \237mM\376? [\214\002 \332\204\035\275q\227\2132A\n\204&\2106\037W|\346#\3413\274\022Y\nE\022\024\n\004lscc\022\014\n\n\n\004mycc\022\002\010\003\022-\n\004mycc\022%\n\007\n\001a\022\002\010\003\n\007\n\001b\022\002\010\003\032\007\n\001a\032\00290\032\010\n\001b\032\003210\032\003\010\310\001\"\013\022\004mycc\032\0031.0" endorsement:<endorser:"\n\007Org1MSP\022\262\006-----BEGIN CERTIFICATE-----\nMIICLjCCAdWgAwIBAgIRAJYomxY2cqHA/fbRnH5a/bwwCgYIKoZIzj0EAwIwczEL\nMAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG\ncmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh\nLm9yZzEuZXhhbXBsZS5jb20wHhcNMTgwMjIyMTYyODE0WhcNMjgwMjIwMTYyODE0\nWjBwMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN\nU2FuIEZyYW5jaXNjbzETMBEGA1UECxMKRmFicmljUGVlcjEfMB0GA1UEAxMWcGVl\ncjAub3JnMS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDEa\nWNNniN3qOCQL89BGWfY39f5V3o1pi//7JFDHATJXtLgJhkK5KosDdHuKLYbCqvge\n46u3AC16MZyJRvKBiw6jTTBLMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAA\nMCsGA1UdIwQkMCKAIN7dJR9dimkFtkus0R5pAOlRz5SA3FB5t8Eaxl9A7lkgMAoG\nCCqGSM49BAMCA0cAMEQCIC2DAsO9QZzQmKi8OOKwcCh9Gd01YmWIN3oVmaCRr8C7\nAiAlQffq2JFlbh6OWURGOko6RckizG8oVOldZG/Xj3C8lA==\n-----END CERTIFICATE-----\n" signature:"0D\002 \022_\342\350\344\231G&\237\n\244\375\302J\220l\302\345\210\335D\250y\253P\0214:\221e\332@\002 \000\254\361\224\247\210\214L\277\370\222\213\217\301\r\341v\227\265\277\336\256^\217\336\005y*\321\023\025\367" >
    2018-02-22 16:34:27.107 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 00b Chaincode invoke successful. result: status:200
    2018-02-22 16:34:27.107 UTC [main] main -> INFO 00c Exiting.....

    ```

    Here you can see that the invoke was submitted successfully based on the log
    message:

    ```
    2018-02-22 16:34:27.107 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 00b Chaincode invoke successful. result: status:200

    ```

    A successful response indicates that the transaction was submitted for ordering
    successfully. The transaction will then be added to a block and, finally, validated
    or invalidated by each peer on the channel.

Here is an example of how to format the `peer chaincode invoke` command when the chaincode package includes multiple smart contracts.

  * If you are using the [contract-api](https://www.npmjs.com/package/fabric-contract-api), the name you pass to `super("MyContract")` can be used as a prefix.

    ```
    peer chaincode invoke -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{ "Args": ["MyContract:methodName", "{}"] }'

    peer chaincode invoke -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{ "Args": ["MyOtherContract:methodName", "{}"] }'

    ```

### peer chaincode query example

Here is an example of the `peer chaincode query` command, which queries the
peer ledger for the chaincode named `mycc` at version `1.0` for the value of
variable `a`:

  * You can see from the output that variable `a` had a value of 90 at the time of
    the query.

    ```
    peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'

    2018-02-22 16:34:30.816 UTC [chaincodeCmd] checkChaincodeCmdParams -> INFO 001 Using default escc
    2018-02-22 16:34:30.816 UTC [chaincodeCmd] checkChaincodeCmdParams -> INFO 002 Using default vscc
    Query Result: 90

    ```

<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.