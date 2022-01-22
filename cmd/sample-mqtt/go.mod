module github.com/rk295/bright-golang/cmd/sample-mqtt

go 1.17

require (
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/rk295/bright-golang/mqtt v0.0.0
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
)

replace github.com/rk295/bright-golang/mqtt v0.0.0 => /u01/home/robin/go/src/github.com/rk295/bright-golang/mqtt/
