// You can edit this code!
// Click here and start typing.
//server.go
package main

import ( "fmt"
	"net/http"
	"github.com/gorilla/websocket"
	)

var upgraded := websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
	return true
	},
}

func wsHandler(w http.ResponseWriter, r * http.Request){

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	
	for {
		messageType, message, err:= conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		
		messageType, message, err:= conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println(err)
			return
		}	
	}
	
}

func main() {
	
	http.Handlefunc("/websocket", wsHandler)
	
	fmt.Println("Sever is listening")
	http.ListenAndServe(":8080",nil)
	fmt.Println("Hello, 世界")
}