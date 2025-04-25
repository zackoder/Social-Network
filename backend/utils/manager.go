package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

type GroupManager struct {
	Groups GroupMembers
	sync.RWMutex
}

type GroupMembers map[int][]int

type OnlineUsers map[int][]*Client

type Manager struct {
	UsersList OnlineUsers
	Groups    GroupMembers
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		UsersList: make(OnlineUsers),
		Groups:    make(GroupMembers),
	}
}

type Client struct {
	Connection *websocket.Conn
	Manager    *Manager
	Client_id  int
	Token      string
}

func (m *Manager) GetClient(id int) []*Client {
	client, ok := m.UsersList[id]
	if !ok {
		return nil
	}
	return client
}

func (m *Manager) CheckGroupMubers(id int) bool {
	_, exists := m.UsersList[id]
	return exists
}

func (m *Manager) AddClient(client *Client) {
	if c, ok := m.UsersList[client.Client_id]; !ok {
		m.UsersList[client.Client_id] = append(m.UsersList[client.Client_id], client)
	} else {
		m.UsersList[client.Client_id] = append(c, client)
	}
}

func (m *Manager) AddGroup(group_id, user_id int) {
	client := m.GetClient(user_id)
	if client != nil {
		if group, ok := m.Groups[group_id]; !ok {
			m.Groups[group_id] = append(group, user_id)
		} else {
			group = append(group, user_id)
		}
	}
}

func CreateClient(conn *websocket.Conn, m *Manager, id int, token string) *Client {
	return &Client{
		Connection: conn,
		Manager:    m,
		Client_id:  id,
		Token:      token,
	}
}

func (m *Manager) StoreGroups(groups []int, user_id int) {
	for _, group_id := range groups {
		if group, exists := m.Groups[group_id]; exists {
			m.Groups[group_id] = append(group, user_id)
		} else {
			m.Groups[group_id] = append(m.Groups[group_id], user_id)
		}
	}
}
