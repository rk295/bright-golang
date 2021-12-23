# `sample-client`

This is a quick eample application which uses the package.

Before running make sure to set the `BRIGHT_USERNAME` and `BRIGHT_PASSWORD` environment variables. No other configuration should be required.

You should be able to just `go run main.go` in this directory, it is set to `Trace` level logging so will be quite chatty:

```
$ go run main.go
DEBU[0000] token is empty, authenticating               
DEBU[0000] making auth request to url: https://api.glowmarkt.com/api/v0-1/auth 
<much output removed>
Current electricity usage: 3121W
```

If it throws an error or does something unexpected please send me the log and I'll look into it. It does not log your username or password but does log the JWT issued by the `/auth` endpoint, feel free to remove that line from the output before sending the log.