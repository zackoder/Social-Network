package controllers

import (
	"encoding/json"
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

		var message utils.Message
		var errormap = map[string]string{}
		fmt.Println("message type", messageType)
		if messageType == websocket.BinaryMessage {
			message, err = utils.UploadMsgImg(pyload)
			if err != nil {
				errormap["error"] = err.Error()
			}
			models.InsertMsg(message)
		}
		if messageType == websocket.TextMessage {
			if err := json.Unmarshal(pyload, &message); err != nil {
				errormap["error"] = "faild to pars your data"
			}
			models.InsertMsg(message)
		}
		for _, clientConnectios := range Manager.UsersList {
			fmt.Println("hello")
			for _, claient := range clientConnectios {
				if _, exist := errormap["error"]; exist {
					if claient.Client_id == message.Sender_id {
						claient.Connection.WriteJSON(errormap)
					}
				} else if messageType == websocket.TextMessage {
					claient.Connection.WriteJSON(message)
				} else if messageType == websocket.BinaryMessage {
					message.Filename = r.Host + message.Filename
					claient.Connection.WriteJSON(message)
				}
			}
		}
		if _, exist := errormap["error"]; exist {
			errormap = map[string]string{}
		}
	}
}
