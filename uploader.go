package main

import (
	"golang.org/x/net/websocket"
	//"io/ioutil"
	//"bytes"
	//"path/filepath"
	"fmt"
	"strings"
	

)


type uploader struct {
	dir    string
	socket *websocket.Conn
}






func (u *uploader)  UploadHandler(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	fmt.Println("yes")
	// first the open connection
	for {
		fmt.Println("loop")
		//if dat, err := ioutil.ReadAll(ws) ; err == nil {
		if num, err := ws.Read(buf); err == nil{
			msg := string(buf[:num])
			fmt.Printf("message : \n%s\n\n", msg)
			if strings.Compare(msg, "stop") == 0 {
				break
			}
			
			
			
			
		} else {
			fmt.Printf("Error : \n%s\n\n", err)
			break
		}
	}
	//u.socket.Close()
	fmt.Println("stop function")
	
}

