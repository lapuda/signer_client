package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/lapuda/signer_client/api"
	"github.com/lapuda/signer_client/client"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	c := client.NewClient(context.Background(), "http://127.0.0.1:9099")
	rsp, err := c.RequestList(api.ListRequest{})
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range rsp.Items {
		fmt.Printf("Address: %s \n", item)
	}
}

func TestNew(t *testing.T) {
	c := client.NewClient(context.Background(), "http://127.0.0.1:9099")
	rsp, err := c.RequestNew(api.NewRequest{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Address: %s \n", rsp.Address)
	fmt.Printf("PK     : %s \n", rsp.Pk)
}

func TestImport(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate ECDSA key: %v", err)
	}
	c := client.NewClient(context.Background(), "http://127.0.0.1:9099")
	rsp, err := c.RequestImport(api.ImportRequest{
		Pk: hex.EncodeToString(privateKey.D.Bytes()),
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Address: %s \n", rsp.Address)
}
