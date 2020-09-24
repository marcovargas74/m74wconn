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

func init() {
	m74wconn.WconnCtrl.VersionSoftware = "20.09.24"
	/*var err error
	dt := time.Now()
	TesterCtrl.VersionSoftware = dt.Format("2006-01-02")
	//Configuracao do LOG
	TesterCtrl.LOOPMain = true
	TesterCtrl.LogProgEnable = true
	TesterCtrl.LogProgPrintLocal = true
	TesterCtrl.LogProgLevel = rlog.Debug | rlog.Local4
	TesterCtrl.LogProgIP = "172.31.11.162:514"

	appl.WorkDir, err = os.Getwd()
	appl.CheckErr(err)

	rlog.Clear()
	rlog.StartLogger(TesterCtrl.LogProgEnable, TesterCtrl.LogProgLevel, TesterCtrl.LogProgIP)
	rlog.SetPrintLocal(TesterCtrl.LogProgPrintLocal)
	rlog.AppSyslog(syslog.LOG_INFO, "%s ======== Start Mannager App Version %s\n", rlog.ThisFunction(), TesterCtrl.VersionSoftware)
	*/
}

/*
 */
func main() {
	//version := "20.07.28"
	//create new WhatsApp connection
	wac, err := whatsapp.NewConn(30 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	fmt.Println("================Starting================================", m74wconn.WconnCtrl.VersionSoftware)

	err = m74wconn.Login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	<-time.After(3 * time.Second)

	/*
		//Cria a mensagem de teste para enviar
		msg := "Eu sou o BOT"
		numContact := "554891119492@s.whatsapp.net"

		//ERIK
		//numContact := "554891175643@s.whatsapp.net"
		if err := m74wconn.SendMessages(msg, numContact, wac); err != nil {
			log.Fatalf("error send message: %v", err)
		}*/

	//ReceiveMessages(wac)
	if err := m74wconn.ReceiveMessages(wac); err != nil {
		log.Fatalf("error reveive message: %v", err)
	}

}
