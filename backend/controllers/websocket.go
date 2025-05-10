package controllers

import (
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"

	"github.com/gorilla/websocket"
)

var Manager = utils.NewManager()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := 10
	client := utils.CreateClient(conn, Manager, id, "hello")
	Manager.AddClient(client)
	groups := models.GetClientGroups(id)
	if len(groups) != 0 {
		go Manager.StoreGroups(groups, id)
	}
	
	fmt.Println("groups", groups)
	defer conn.Close()
	for {
		messageType, pyload, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(err)
			}
			break
		}
		imagPath := ""
		if messageType == websocket.BinaryMessage {
			imagPath = r.Host + "/" + utils.UploadMsgImg(pyload)
		}
		for _, clientConnectios := range Manager.UsersList {
			for _, conn := range clientConnectios {
				if messageType == websocket.TextMessage {
					conn.Connection.WriteJSON(map[string]string{"msg": string(pyload)})
				} else if messageType == websocket.BinaryMessage {
					conn.Connection.WriteJSON(map[string]string{"image": imagPath})
				}
			}
		}
	}
}
