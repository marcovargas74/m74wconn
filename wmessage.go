package m74wconn

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
)

//------------------RECEIVE Messages-------------------------------
type waHandler struct {
	c *whatsapp.Conn
}

//HandleError needs to be implemented to be a valid WhatsApp handler
func (h *waHandler) HandleError(err error) {

	if e, ok := err.(*whatsapp.ErrConnectionFailed); ok {
		log.Printf("Connection failed, underlying error: %v", e.Err)
		log.Println("Waiting 30sec...")
		<-time.After(30 * time.Second)
		log.Println("Reconnecting...")
		err := h.c.Restore()
		if err != nil {
			log.Fatalf("Restore failed: %v", err)
		}
	} else {
		log.Printf("error occoured: %v\n", err)
	}
}

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (*waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	//fmt.Printf("%v %v %v %v\n\t%v\n", message.Info.Timestamp, message.Info.Id, message.Info.RemoteJid, message.ContextInfo.QuotedMessageID, message.Text)
	fmt.Printf("%v %v\n\t%v\n", message.Info.RemoteJid, message.ContextInfo.QuotedMessageID, message.Text)
}

//ReceiveMessages Receive Messages
func ReceiveMessages(wac *whatsapp.Conn) error {

	//Add handler
	wac.AddHandler(&waHandler{wac})

	//verifies phone connectivity
	pong, err := wac.AdminTest()

	if !pong || err != nil {
		return fmt.Errorf("error pinging in: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	//Disconnect safe
	fmt.Println("Shutting down now.")
	session, err := wac.Disconnect()
	if err != nil {
		return fmt.Errorf("error disconnecting: %v", err)
	}
	if err := WriteSession(session); err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

//SendMessages Send Messages using whatsApp Account
func SendMessages(message string, contact string, wac *whatsapp.Conn) error {
	previousMessage := "😜"
	quotedMessage := proto.Message{
		Conversation: &previousMessage,
	}

	ContextInfo := whatsapp.ContextInfo{
		QuotedMessage:   &quotedMessage,
		QuotedMessageID: "554891175643@s.whatsapp.net",
		Participant:     "554891175643@s.whatsapp.net", //Whot sent the original message
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: contact, //"5548991119492@s.whatsapp.net", //Erik 554891175643
			SenderJid: "5548991119492@s.whatsapp.net",
		},
		ContextInfo: ContextInfo,
		Text:        message,
	}

	msgID, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err) //		os.Exit(1)
	}

	fmt.Println("Message Sent -> ID : " + msgID)
	return nil
}
