package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"social-network/models"
	"social-network/utils"

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
		fmt.Println(string(payload))
		handleMessage(msgType, payload, r.Host, client)
	}
}

func handleMessage(msgType int, payload []byte, host string, client *utils.Client) {
	var err error
	var msg utils.Message
	var errMsg utils.Err

	switch msgType {
	case websocket.BinaryMessage:
		msg, err = utils.UploadMsgImg(payload)
		if err != nil {
			errMsg.Error = err.Error()
			Broadcast(client.Client_id, errMsg)
			return
		}
	case websocket.TextMessage:
		if err := json.Unmarshal(payload, &msg); err != nil {
			errMsg.Error = "failed to parse your data"
			Broadcast(client.Client_id, errMsg)
			log.Println("JSON Unmarshal error:", err)
			return
		}
	default:
		return
	}

	if msg.Reciever_id != 0 {
		BroadcastPrivateMessage(msg, host)
	} else if msg.Group_id != 0 {
		if err := BroadcastGroupMessage(msg, host); err != nil {
			errMsg.Error = err.Error()
			Broadcast(msg.Sender_id, errMsg)
		}
	}
}

func BroadcastPrivateMessage(msg utils.Message, host string) {
	errMsg := utils.Err{}

	if ok, err := models.FriendsChecker(msg.Sender_id, msg.Reciever_id); err != nil || !ok {
		if err := os.Remove("." + msg.Filename); err != nil {
			fmt.Println(err)
		}
		errMsg.Error = "you need to follow the receiver first"
		Broadcast(msg.Sender_id, errMsg)
		return
	}

	models.InsertMsg(msg)
	if msg.Filename != "" {
		msg.Filename = host + msg.Filename
	}
	Broadcast(msg.Reciever_id, msg)
	Broadcast(msg.Sender_id, msg)
}

func BroadcastGroupMessage(msg utils.Message, host string) error {
	if !models.CheckSender(msg.Group_id, msg.Sender_id) {
		return fmt.Errorf("you need to be a group member first")
	}
	for _, receiverID := range Manager.Groups[msg.Group_id] {
		Broadcast(receiverID, msg)
	}
	models.InsertGroupMSG(msg)
	return nil
}

func Broadcast(receiverID int, msg any) {
	if connections, exists := Manager.UsersList[receiverID]; exists {
		for _, conn := range connections {
			if err := conn.Connection.WriteJSON(msg); err != nil {
				log.Println("WriteJSON failed:", err)
			}
		}
	}
}
