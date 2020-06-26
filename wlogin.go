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
	//file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	file, err := os.Open("whatsappSession.gob")
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
	//file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	file, err := os.Create("whatsappSession.gob")
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
			msg := <-qr
			qrcode.WriteFile(msg, qrcode.Medium, 256, "qrcode.png")
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
	fmt.Printf("Connected at the whatsApp Server WiD[%s]. ", session.Wid)
	return nil
}
