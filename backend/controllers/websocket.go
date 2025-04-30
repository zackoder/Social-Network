package controllers

import (
	"fmt"
	"net/http"
	"social-network/models"
	"social-network/utils"
	"strconv"

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

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Println(id)
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
		var code int
		var message utils.Message
		if messageType == websocket.BinaryMessage {
			message, err, code = utils.UploadMsgImg(pyload)
			if err != nil {
				utils.WriteJSON(w, map[string]string{"error": err.Error()}, code)
				return
			}
			message.Filename = r.Host + "/" + message.Filename
		}
		for _, clientConnectios := range Manager.UsersList {
			for _, conn := range clientConnectios {
				if messageType == websocket.TextMessage {
					conn.Connection.WriteJSON(map[string]string{"msg": string(pyload)})
				} else if messageType == websocket.BinaryMessage {
					conn.Connection.WriteJSON(message)
				}
			}
		}
	}
}
