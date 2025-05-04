package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	client := utils.CreateClient(conn, Manager, id, "online")
	Manager.AddClient(client)
	defer Manager.RemoveClient(client)

	if groups := models.GetClientGroups(id); len(groups) > 0 {
		go Manager.StoreGroups(groups, id)
	}

	for {
		msgType, payload, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("WebSocket closed unexpectedly:", err)
			}
			break
		}
		handleMessage(msgType, payload, r.Host)
	}
}

func handleMessage(msgType int, payload []byte, host string) {
	var err error
	var msg utils.Message
	var errMsg utils.Err

	switch msgType {
	case websocket.BinaryMessage:
		msg, err = utils.UploadMsgImg(payload)
		if err != nil {
			errMsg.Error = err.Error()
			broadcastError(errMsg, msg.Sender_id)
			return
		}
	case websocket.TextMessage:
		if err := json.Unmarshal(payload, &msg); err != nil {
			errMsg.Error = "failed to parse your data"
			broadcastError(errMsg, msg.Sender_id)
			log.Println("JSON Unmarshal error:", err)
			return
		}
	default:
		return
	}

	if msg.Reciever_id != 0 {
		broadcastPrivateMessage(msg, host)
	} else if msg.Group_id != 0 {
		if err := broadcastGroupMessage(msg, host); err != nil {
			errMsg.Error = err.Error()
			broadcastError(errMsg, msg.Sender_id)
		}
	}
}

func broadcastPrivateMessage(msg utils.Message, host string) {
	errMsg := utils.Err{}

	if ok, err := models.FriendsChecker(msg.Sender_id, msg.Reciever_id); err != nil || !ok {
		errMsg.Error = "you need to follow the receiver first"
		broadcastError(errMsg, msg.Sender_id)
		return
	}

	models.InsertMsg(msg)
	if msg.Filename != "" {
		msg.Filename = host + msg.Filename
	}
	broadcastMessage(msg.Reciever_id, msg)
	broadcastMessage(msg.Sender_id, msg)
}

func broadcastGroupMessage(msg utils.Message, host string) error {
	if !models.CheckSender(msg.Group_id, msg.Sender_id) {
		return fmt.Errorf("you need to be a group member first")
	}
	for _, receiverID := range Manager.Groups[msg.Group_id] {
		broadcastMessage(receiverID, msg)
	}
	models.InsertGroupMSG(msg)
	return nil
}

func broadcastMessage(receiverID int, msg utils.Message) {
	if connections, exists := Manager.UsersList[receiverID]; exists {
		for _, conn := range connections {
			if err := conn.Connection.WriteJSON(msg); err != nil {
				log.Println("WriteJSON failed:", err)
			}
		}
	}
}

func broadcastError(errMsg utils.Err, receiverID int) {
	if connections, exists := Manager.UsersList[receiverID]; exists {
		for _, conn := range connections {
			if err := conn.Connection.WriteJSON(errMsg); err != nil {
				log.Println("Error broadcast failed:", err)
			}
		}
	}
}

func BroadcastNotification(noti utils.Notification) {
	if targets, online := Manager.UsersList[noti.Target_id]; online {
		for _, conn := range targets {
			if err := conn.Connection.WriteJSON(noti); err != nil {
				log.Println("Notification broadcast failed:", err)
			}
		}
	}
}
