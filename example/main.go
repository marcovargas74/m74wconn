package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/marcovargas74/m74wconn"
)

//------------------CONNECT-------------------------------

/*
Test
*/
func main() {
	//create new WhatsApp connection
	wac, err := whatsapp.NewConn(60 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	fmt.Println("================================================")

	err = m74wconn.Login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	<-time.After(3 * time.Second)

	//Cria a mensagem de teste para enviar
	msg := "Eu sou o BOT"
	numContact := "5548991119492@s.whatsapp.net"

	//ERIK
	//numContact := "554891175643@s.whatsapp.net"
	if err := m74wconn.SendMessages(msg, numContact, wac); err != nil {
		log.Fatalf("error send message: %v", err)
	}

	//ReceiveMessages(wac)
	if err := m74wconn.ReceiveMessages(wac); err != nil {
		log.Fatalf("error reveive message: %v", err)
	}

}
