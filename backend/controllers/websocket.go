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
			continue
		}
		if messageType == websocket.TextMessage {
			if err := json.Unmarshal(pyload, &message); err != nil {
				errormap["error"] = "faild to pars your data"
				fmt.Println(err)
			}
		}
		fmt.Println(message)

		if _, exist := errormap["error"]; exist {
			for _, claient := range Manager.UsersList[client.Client_id] {
				claient.Connection.WriteJSON(errormap)
			}
			continue
		}
		fmt.Println(Manager.Groups[message.Group_id])
		fmt.Println("sender id", message.Sender_id)
		if message.Reciever_id != 0 {
			BrodcatstPrivetMSG(message, messageType, r.Host)

		} else if message.Group_id != 0 {
			BrodcatstgroupMSG(message, messageType, r.Host)
		}
	}
}

func BrodcatstPrivetMSG(message utils.Message, messageType int, host string) {
	friends, err := models.FriendsChecker(message.Sender_id, message.Reciever_id)
	var ERR utils.Err
	if err != nil {
		fmt.Println("error", err)
		return
	}

	if !friends {
		ERR.Error = "you need to follow the reciever first"
		BrodcastError(ERR, message.Sender_id)
		return
	}

	if recieverConnectios, exists := Manager.UsersList[message.Reciever_id]; exists {
		for _, reciever := range recieverConnectios {
			if messageType == websocket.TextMessage {
				reciever.Connection.WriteJSON(message)
			} else if messageType == websocket.BinaryMessage {
				message.Filename = host + message.Filename
				reciever.Connection.WriteJSON(message)
			}
		}
	}
	models.InsertMsg(message)
}

func BrodcastError(err utils.Err, reciever_id int) {
	if senderConnectios, exists := Manager.UsersList[reciever_id]; exists {
		for _, sender := range senderConnectios {
			sender.Connection.WriteJSON(err)
		}
	}
}

func BrodcatstgroupMSG(message utils.Message, messageType int, host string) {
	fmt.Println(Manager.Groups[message.Group_id])
	for _, reciever := range Manager.Groups[message.Group_id] {
		if recieverConnectios, exists := Manager.UsersList[reciever]; exists {
			for _, reciever := range recieverConnectios {
				if messageType == websocket.TextMessage {
					reciever.Connection.WriteJSON(message)
				} else if messageType == websocket.BinaryMessage {
					message.Filename = host + message.Filename
					reciever.Connection.WriteJSON(message)
				}
			}
		}
	}
	models.InsertGroupMSG(message)
}

func BrodcastNoti(noti utils.Notification) {
	if target, online := Manager.UsersList[noti.Target_id]; online {
		for _, targetConnection := range target {
			targetConnection.Connection.WriteJSON(noti)
		}
	}
}
