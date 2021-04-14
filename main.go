package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection(urlList string) (*nats.Conn, *nats.EncodedConn) {
	natsConnection, err := nats.Connect(urlList, nats.MaxReconnects(3), nats.ReconnectWait(time.Second))
	if err != nil {
		fmt.Println("Error while connecting NATS:", err)
		os.Exit(1)
	}
	natsEncodedConnection, err := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)
	if err != nil {
		fmt.Println("Error while connecting NATS:", err)
		os.Exit(1)
	}
	return natsConnection, natsEncodedConnection
}

func main() {
	natsConnection, neConnection := NewNATSConnection("localhost:4222")
	defer natsConnection.Close()
	defer neConnection.Close()

	traceID, _ := uuid.NewV4()

	subject := "siman-util-sinkronisasi.t_loader_sakti.create." + traceID.String()
	fmt.Println("Subject:", subject)
	file, _ := os.Open("sample1.json")
	buf, _ := ioutil.ReadAll(file)
	jsonMessage := (json.RawMessage)(buf)

	if err := neConnection.Publish(subject, jsonMessage); err != nil {
		fmt.Println("Error while connecting NATS:", err)
		os.Exit(1)
	}
}
