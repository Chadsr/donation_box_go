package main

import (
	"fmt"

	"github.com/joernweissenborn/eventual2go"
	"github.com/joernweissenborn/serialreactor"
)

func main() {
	fmt.Println("moin")
	sr, _ := serialreactor.New()
	sr.OnRead(onRead)
	for {}
}

func onRead(d eventual2go.Data) {
	raw := d.([]byte)
	fmt.Println(string(raw))

}

