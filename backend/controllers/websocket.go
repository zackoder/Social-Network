package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

func Websocket(w http.ResponseWriter, r *http.Request, user_id int) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	cookie, _ := r.Cookie("token")
	client := utils.CreateClient(conn, Manager, user_id, cookie.Value)
	log.Println("client tres to log", client)
	go Manager.AddClient(client)
	defer func() {
		log.Println("1", client)
		go Manager.RemoveClient(client)
		log.Println("Client disconnected:", user_id)
	}()

	if groups := models.GetClientGroups(user_id); len(groups) > 0 {
		go Manager.StoreGroups(groups, user_id)
	}

	for {
		log.Println("dslkdfjqlskjflsjf")
		msgType, payload, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("WebSocket closed unexpectedly:", err)
			}
			break
		}

		log.Println("Message received type:", msgType)
		handleMessage(msgType, payload, r.Host, client)
	}
}

func handleMessage(msgType int, payload []byte, host string, client *utils.Client) {
	var (
		msg    utils.Message
		errMsg utils.Err
		err    error
	)

	switch msgType {
	case websocket.BinaryMessage:
		msg, err = utils.UploadMsgImg(payload)
		if err != nil {
			errMsg.Error = err.Error()
			Broadcast(client.Client_id, errMsg)
			return
		}

	case websocket.TextMessage:
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			errMsg.Error = "Failed to parse your data"
			Broadcast(client.Client_id, errMsg)
			log.Println("JSON unmarshal error:", err)
			return
		}

	default:
		log.Println("Unsupported message type:", msgType)
		return
	}

	msg.Sender_id = client.Client_id
	user, err := models.GetUserById(msg.Sender_id)
	if err != nil {
		errMsg.Error = "User not found"
		Broadcast(client.Client_id, errMsg)
		return
	}

	msg.Avatar = host + user.Avatar
	msg.First_name = user.FirstName
	msg.Last_name = user.LastName

	switch {
	case msg.Reciever_id != 0:
		BroadcastPrivateMessage(msg, host)

	case msg.Group_id != 0:
		err = BroadcastGroupMessage(msg, host)
		if err != nil {
			errMsg.Error = err.Error()
			Broadcast(msg.Sender_id, errMsg)
		}

	default:
		log.Println("Message has no recipient")
	}
}

func BroadcastPrivateMessage(msg utils.Message, host string) {
	errMsg := utils.Err{}
	ok, err := models.FriendsCheckerForMessages(msg.Sender_id, msg.Reciever_id)
	if err != nil || !ok {
		if msg.Filename != "" {
			_ = os.Remove("." + msg.Filename)
		}
		errMsg.Error = "You need to follow the receiver first"
		Broadcast(msg.Sender_id, errMsg)
		return
	}

	err = models.InsertMsg(msg)
	if err != nil {
		log.Println("Failed to save private message:", err)
		return
	}

	if msg.Filename != "" {
		msg.Filename = host + msg.Filename // host + msg.Filename
	}

	Broadcast(msg.Reciever_id, msg)
	Broadcast(msg.Sender_id, msg)
}

func BroadcastGroupMessage(msg utils.Message, host string) error {
	if !models.CheckSender(msg.Group_id, msg.Sender_id) {
		return fmt.Errorf("You need to be a group member first")
	}

	err := models.InsertGroupMSG(msg)
	if err != nil {
		log.Println("Failed to save group message:", err)
		return err
	}

	if msg.Filename != "" {
		msg.Filename = host + msg.Filename
	}

	for _, receiverID := range Manager.Groups[msg.Group_id] {
		Broadcast(receiverID, msg)
	}

	return nil
}

func Broadcast(receiverID int, msg any) {
	if connections, exists := Manager.UsersList[receiverID]; exists {
		for _, conn := range connections {
			if err := conn.Connection.WriteJSON(msg); err != nil {
				log.Println("WriteJSON failed for user", receiverID, ":", err)
			}
		}
	}
}
