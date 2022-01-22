## Clymene-gateway

Clymene-gateway can receive data through gRPC.

```
Clymene-gateway [flags]
```

### Options

```
--tdengine.dbname string                 Destination database (default "clymene")
--tdengine.hostname string               The host to connect to TDengine server. (default "127.0.0.1")
--tdengine.max-sql-length int            Number of SQLs that can be sent at one time (default 4096)
--tdengine.password string               The password to use when connecting to the server (default "taosdata")
--tdengine.server-port int               he HTTP port number to use for the connection to TDengine server (default 6041)
--tdengine.user string                   The TDengine user name to use when connecting to the server (default "root")
```

###### Auto generated by spf13/cobra on 22-Jan-2022