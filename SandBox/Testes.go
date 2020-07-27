package main

import (
	"fmt"

	"github.com/marcovargas74/m74wconn"
)

//------------------CONNECT-------------------------------

/*
Test
*/
func main() {
	//create new WhatsApp connection
	/*wac, err := whatsapp.NewConn(30 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}*/

	fmt.Println("================Starting==TESTE============================")

	m74wconn.ExecLinuxCmd("sh", "updateW.sh")

	/*//cmd := exec.Command("sh updateW.sh").Output()
	//_, err := cmd.Output()

	if err != nil {
		//Println(err.Error())
		return
	}*/

	//Print(string(stdout))
	/*err = m74wconn.Login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}*/
}
