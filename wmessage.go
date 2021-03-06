package m74wconn

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/Rhymen/go-whatsapp/binary/proto"
	//tw "github.com/marcovargas74/m74twitter"
)

//Twconn Type wconn Variaveis Globais usados em todo o wconn
type Twconn struct {
	LogProgEnable   bool
	VersionSoftware string //Versao da aplicação Default
}

//WconnCtrl Type wconn Variaveis Globais usados em todo o wconn
var WconnCtrl Twconn

//------------------RECEIVE Messages-------------------------------
type waHandler struct {
	c         *whatsapp.Conn
	startTime uint64
}

//ExecLinuxCmd Roda um comando em Linux
func ExecLinuxCmd(cmd string, file string) []byte {
	out, err := exec.Command(cmd, file).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("return: %s\n", out)
	return out
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

//var StartTime = uint64(time.Now().Unix()) // once you start the API we save the start time
//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (h *waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	//fmt.Printf("%v %v %v %v\n\t%v\n", message.Info.Timestamp, message.Info.Id, message.Info.RemoteJid, message.ContextInfo.QuotedMessageID, message.Text)
	//Discart old messages
	if message.Info.Timestamp < h.startTime {
		return
	}

	time.Sleep(500 * time.Millisecond)

	if message.Info.FromMe {

		//Verifica se está rodando
		if strings.Contains(strings.ToLower(message.Text), "bot run") {
			fmt.Printf("recebeu de %v \n", message.Info.RemoteJid)
			msg := fmt.Sprintf("BOT rodando!OK version %s", WconnCtrl.VersionSoftware)
			SendMessages(msg, message.Info.RemoteJid, h.c)
			//msg := "Eu sou o BOT" //		numContact := "554891119492@s.whatsapp.net"
			//fmt.Printf("recebeu de Unniti %v \n", message.Info.RemoteJid)
			//m74wconn.SendMessages(msg, numContact, wac)
			return
		}

		//Atualiza Bot
		if strings.Contains(strings.ToLower(message.Text), "bot update") {
			fmt.Printf("recebeu de %v \n", message.Info.RemoteJid)
			msg := "BOT sera atualizado!OK" //
			SendMessages(msg, message.Info.RemoteJid, h.c)
			ExecLinuxCmd("sh", "/data/whatsapp/updateW.sh")
			//msg := "Eu sou o BOT" //		numContact := "554891119492@s.whatsapp.net"
			//fmt.Printf("recebeu de Unniti %v \n", message.Info.RemoteJid)
			//m74wconn.SendMessages(msg, numContact, wac)
		}
		return

	}

	fmt.Printf("%v %v\n\t%v\n", message.Info.RemoteJid, message.ContextInfo.QuotedMessageID, message.Text)

	//Discart Message From Token Grupo Pista 554884923044-1486039747@g.us
	if strings.Contains(message.Info.RemoteJid, "-1397069327") || //Pista_limpa 1
		strings.Contains(message.Info.RemoteJid, "-1486039747") || //Pista_limpa 3
		strings.Contains(message.Info.RemoteJid, "-1473907385") { //Pista_limpa Estadual
		fmt.Printf("recebeu do #pista_limpa %v \n", message.Info.RemoteJid)
		//SendMessages(message.Text, numContact, h.c)
		/*if strings.Contains(message.Text, "#pista_limpa") {
		      //sendTwitter(message.Text)
				}*/

		return
	}

	//Token Grupo Comunicação 554899496824-1386712719@g.us
	if strings.Contains(message.Info.RemoteJid, "554899496824-1386712719") {

		if strings.Contains(strings.ToLower(message.Text), "unniti") {
			//Token Grupo Comunicação 554899496824-1386712719@g.us
			fmt.Printf("recebeu do #Grupo Comunicacao %v \n", message.Info.RemoteJid)
			msg := "Alguém falou de MIM? Eu sou o BOT Unniti" //
			SendMessagesEmoji(msg, message.Info.RemoteJid, h.c)
			//fmt.Printf("recebeu de Unniti %v \n", message.Info.RemoteJid)
			//m74wconn.SendMessages(msg, numContact, wac)
			return
		}

		if strings.Contains(strings.ToLower(message.Text), "felicidades") {
			msg := "Parabéns!!! Muitas Felicidades !!! Muita Saúde, SUCESSO!!" //
			SendMessages(msg, message.Info.RemoteJid, h.c)
			return
		}
		/*
			if strings.Contains(strings.ToLower(message.Text), "bom dia") {
				msg := "Bom Dia!" //
				SendMessages(msg, message.Info.RemoteJid, h.c)
			}*/

		if strings.Contains(strings.ToLower(message.Text), "boa tarde") {
			msg := "Boa Tarde!" //
			SendMessages(msg, message.Info.RemoteJid, h.c)
			return
		}

		if strings.Contains(strings.ToLower(message.Text), "boa noite") {
			msg := "Boa Noite!" //
			SendMessages(msg, message.Info.RemoteJid, h.c)

		}

		return
	}

	/* /Token Grupo Pista 554884923044-1486039747@g.us
	if strings.Contains(message.Text, "#pista_limpa") &&
		strings.Contains(message.Info.RemoteJid, "554884923044-1486039747") {
		//msg := "Eu sou o BOT" //
		//numContact := "554891119492@s.whatsapp.net"
		fmt.Printf("recebeu do #pista_limpa %v \n", message.Info.RemoteJid)
		//SendMessages(message.Text, numContact, h.c)
		//sendTwitter(message.Text)
		return
	}*/

	//Token
	if strings.Contains(strings.ToLower(message.Text), "pão de batata") {
		fmt.Printf("recebeu de %v \n", message.Info.RemoteJid)
		msg := "Alguém falou em Pão de Batata? Eu sou um BOT " //
		SendMessagesEmoji(msg, message.Info.RemoteJid, h.c)
		//msg := "Eu sou o BOT" //		numContact := "554891119492@s.whatsapp.net"
		//fmt.Printf("recebeu de Unniti %v \n", message.Info.RemoteJid)
		//m74wconn.SendMessages(msg, numContact, wac)
		return
	}

	/*
		if strings.Contains(strings.ToLower(message.Text), "bom dia") {
			msg := "Bom Dia!" //
			SendMessages(msg, message.Info.RemoteJid, h.c)
		}*/

}

//ReceiveMessages Receive Messages
func ReceiveMessages(wac *whatsapp.Conn) error {

	//Add handler
	//wac.AddHandler(&waHandler{wac})
	wac.AddHandler(&waHandler{wac, uint64(time.Now().Unix())})

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
	previousMessage := ""
	quotedMessage := proto.Message{
		Conversation: &previousMessage,
	}

	ContextInfo := whatsapp.ContextInfo{
		QuotedMessage:   &quotedMessage,
		QuotedMessageID: wac.Info.Wid, //"554891119492@s.whatsapp.net",
		Participant:     wac.Info.Wid, //"554891119492@s.whatsapp.net", //Whot sent the original message
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: contact,      //"5548991119492@s.whatsapp.net", //Erik 554891175643
			SenderJid: wac.Info.Wid, // "554891119492@s.whatsapp.net",
		},
		ContextInfo: ContextInfo,
		Text:        message,
	}

	msgID, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err) //		os.Exit(1)
	}

	fmt.Println("Message Sent -> ID : "+msgID+" Num: "+msg.Info.RemoteJid, " dst: "+wac.Info.Wid)
	return nil
}

//SendMessagesEmoji Send Messages using whatsApp Account
func SendMessagesEmoji(message string, contact string, wac *whatsapp.Conn) error {
	previousMessage := "😜"
	quotedMessage := proto.Message{
		Conversation: &previousMessage,
	}

	ContextInfo := whatsapp.ContextInfo{
		QuotedMessage:   &quotedMessage,
		QuotedMessageID: wac.Info.Wid, //"554891119492@s.whatsapp.net",
		Participant:     wac.Info.Wid, //"554891119492@s.whatsapp.net", //Whot sent the original message
	}

	msg := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: contact,      //"5548991119492@s.whatsapp.net", //Erik 554891175643
			SenderJid: wac.Info.Wid, // "554891119492@s.whatsapp.net",
		},
		ContextInfo: ContextInfo,
		Text:        message,
	}

	msgID, err := wac.Send(msg)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err) //		os.Exit(1)
	}

	fmt.Println("Message Sent -> ID : "+msgID+" Num: "+msg.Info.RemoteJid, " dst: "+wac.Info.Wid)
	return nil
}

/* TWITEER
 * /
func sendTwitter(messageFromWhats string) {
	fmt.Println("Go-Twitter Bot v0.01")
	/*creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("TOKEN >%+v\n", os.Getenv("ACCESS_TOKEN"))
	fmt.Printf("TOKEN_SECRET >%+v\n", os.Getenv("ACCESS_TOKEN_SECRET"))
	fmt.Printf("Consumer >%+v\n", os.Getenv("CONSUMER_KEY"))
	fmt.Printf("ConsumerSec >%+v\n", os.Getenv("CONSUMER_SECRET"))

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}* /

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	//fmt.Printf("%+v\n", client)

	client, err := tw.ConectTwitter()
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	//sendaTwitter(client)
	//tw.FindTwitter(client)
	tw.SendTwitter(client, messageFromWhats)

}
*/
