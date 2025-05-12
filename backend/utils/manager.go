package utils

import (
	"fmt"
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
	m.RLock()
	defer m.RUnlock()
	client, ok := m.UsersList[id]
	if !ok {
		return nil
	}
	return client
}

func (m *Manager) CheckGroupMubers(id int) bool {
	m.RLock()
	m.RUnlock()
	_, exists := m.UsersList[id]
	return exists
}

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	m.Unlock()
	if c, ok := m.UsersList[client.Client_id]; !ok {
		m.UsersList[client.Client_id] = append(m.UsersList[client.Client_id], client)
	} else {
		m.UsersList[client.Client_id] = append(c, client)
	}
}

func (m *Manager) AddGroup(group_id, user_id int) {
	m.Lock()
	defer m.Unlock()
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
	m.Lock()
	defer m.Unlock()
	for _, group_id := range groups {
		if group, exists := m.Groups[group_id]; exists {
			if !m.CheckuserExistenc(user_id, group_id) {
				m.Groups[group_id] = append(group, user_id)
			}
		} else {
			m.Groups[group_id] = append(m.Groups[group_id], user_id)
		}
	}
}

func (m *Manager) CheckuserExistenc(user_id, group_id int) bool {
	m.RLock()
	defer m.RUnlock()
	for _, client := range m.Groups[group_id] {
		if client == user_id {
			return true
		}
	}
	return false
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	clients := m.UsersList[client.Client_id]
	for i, c := range clients {
		if c == client {
			m.UsersList[client.Client_id] = append(clients[:i], clients[i+1:]...)
			break
		}
	}
	fmt.Println(len(m.UsersList[client.Client_id]))
	// If no clients left, delete entry
	if len(m.UsersList[client.Client_id]) == 0 {
		delete(m.UsersList, client.Client_id)
		m.RemoveClientFromGroups(client.Client_id)
	}
}

func (m *Manager) RemoveClientFromGroups(id int) {
	for _, group := range m.Groups {
		for index, client_id := range group {
			if client_id == id {
				group = append(group[:index], group[index+1:]...)
			}
		}
	}
}
