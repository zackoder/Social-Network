package utils

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type (
	GroupMembers map[int][]int
	OnlineUsers  map[int][]*Client
)

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

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()


	// Initialize slice if it doesn't exist
	if m.UsersList[client.Client_id] == nil {
		m.UsersList[client.Client_id] = make([]*Client, 0)
	}

	m.UsersList[client.Client_id] = append(m.UsersList[client.Client_id], client)
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
			if !m.checkUserExistence(user_id, group_id) {
				m.Groups[group_id] = append(group, user_id)
			}
		} else {
			m.Groups[group_id] = append(m.Groups[group_id], user_id)
		}
	}
}

func (m *Manager) checkUserExistence(user_id, group_id int) bool {
	// This method should be called with lock already held
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

	log.Printf("Removing client %d", client.Client_id)

	clients := m.UsersList[client.Client_id]
	for i, c := range clients {
		if c == client {
			m.UsersList[client.Client_id] = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	log.Printf("Client %d removed. Remaining connections: %d", client.Client_id, len(m.UsersList[client.Client_id]))

	// If no clients left, delete entry
	if len(m.UsersList[client.Client_id]) == 0 {
		delete(m.UsersList, client.Client_id)
		m.removeClientFromGroups(client.Client_id)
	}
}

func (m *Manager) removeClientFromGroups(id int) {
	// This method should be called with lock already held
	for group_id, group := range m.Groups {
		for index, client_id := range group {
			if client_id == id {
				m.Groups[group_id] = append(group[:index], group[index+1:]...)
				break
			}
		}
	}
}
