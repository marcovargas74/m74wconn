package m74wconn

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/Rhymen/go-whatsapp"
	qrcode "github.com/skip2/go-qrcode"
)

//------------------CONNECT-------------------------------

//ReadSession Connection with whatsApp if this exist
func ReadSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

//WriteSession Connection with whatsApp if this exist
func WriteSession(session whatsapp.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

//Login in Whatsapp
func Login(wac *whatsapp.Conn) error {
	//load saved session
	session, err := ReadSession()
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v", err)
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			//terminal := qrcodeTerminal.New()
			//terminal.Get(<-qr).Print()
			msg := <-qr
			qrcode.WriteFile(msg, qrcode.Medium, 256, "qrcode.png")
			fmt.Printf("login , msg: %s\n", msg)
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	//save session
	err = WriteSession(session)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}
	return nil
}

/*
//------------------RECEIVE Messages-------------------------------

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
*/
