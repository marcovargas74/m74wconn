package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/skip2/go-qrcode"
)

func main() {
	wac, err := whatsapp.NewConn(60 * time.Second)
	if err != nil {
		panic(err)
	}

	qr := make(chan string)
	go func() {
		//terminal := qrcodeTerminal.New()
		//terminal.Get(<-qr).Print()
		msg := <-qr
		qrcode.WriteFile(msg, qrcode.Medium, 256, "qr1.png")
		fmt.Printf("login , msg: %s\n", msg)
	}()

	fmt.Printf("login , myTOken: %v\n", qr)
	session, err := wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
	}
	fmt.Printf("login successful, session: %v\n", session)
}
