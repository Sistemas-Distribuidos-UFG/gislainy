package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type SocketMessage struct {
	Event   string                 `json:"event"`
	Message map[string]interface{} `json:"message"`
}

type SocketConn struct {
	conn net.Conn
	id   string
}

type UserData struct {
	UserName string `json:"userName"`
	RoomId   string `json:"roomId"`
	conn     net.Conn
	id       string
}
type Message struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
}

type SocketServer struct {
	users []UserData
	rooms map[string][]UserData
}

type Broadcast struct {
	roomId               string
	event                string
	message              map[string]interface{}
	includeCurrentSocket bool
}

func main() {
	controller := SocketServer{}
	controller.rooms = make(map[string][]UserData)

	service := ":9898"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go controller.onNewConnection(conn)

	}
}

func (server *SocketServer) onNewConnection(conn net.Conn) error {
	server.handleClient(conn)
	return nil
}

func (server *SocketServer) broadcast(broadcast Broadcast) {
	socketMessage := SocketMessage{Event: broadcast.event, Message: broadcast.message}
	for _, user := range server.rooms[broadcast.roomId] {
		server.sendMessage(user.conn, socketMessage)
	}
}

func (server *SocketServer) handleClient(conn net.Conn) {
	request := make([]byte, 128)
	defer conn.Close()
	socketConn := SocketConn{conn: conn, id: uuid.New().String()}
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			server.disconnect(socketConn)
			break
		}

		var socketMessage SocketMessage

		if err := json.Unmarshal(request[:read_len], &socketMessage); err != nil {
			fmt.Fprintf(os.Stderr, "Message invalid: %s", err.Error())
		}

		switch socketMessage.Event {
		case "joinRoom":
			var userData UserData
			mapstructure.Decode(socketMessage.Message, &userData)
			server.joinRoom(socketConn, userData)
		case "message":
			var message Message
			mapstructure.Decode(socketMessage.Message, &message)
			server.message(socketConn, message)
		}

	}

}
func (server *SocketServer) joinRoom(socketConn SocketConn, userData UserData) {
	fmt.Printf("%s joined\n", userData.UserName)

	userData.conn = socketConn.conn
	userData.id = socketConn.id
	server.rooms[userData.RoomId] = append(server.rooms[userData.RoomId], userData)
	server.users = append(server.users, userData)

	messageUpdateUsers := map[string]interface{}{
		`users`: server.rooms[userData.RoomId],
	}
	socketMessage := SocketMessage{"updateUsers", messageUpdateUsers}
	server.sendMessage(socketConn.conn, socketMessage)

	messageBroadcast := map[string]interface{}{
		`userName`: userData.UserName,
		`roomId`:   userData.RoomId,
	}

	b := Broadcast{event: "newUserConnected", message: messageBroadcast, roomId: userData.RoomId, includeCurrentSocket: false}

	server.broadcast(b)

}

func (server *SocketServer) message(socketConn SocketConn, message Message) {

	userData := server.getUser(socketConn)
	messageBroadcast := map[string]interface{}{
		`userName`: userData.UserName,
		`message`:  message.Message,
	}
	b := Broadcast{event: "message", message: messageBroadcast, roomId: userData.RoomId, includeCurrentSocket: false}
	server.broadcast(b)

}

func (server *SocketServer) getUser(socketConn SocketConn) UserData {
	var userData UserData

	for _, item := range server.users {

		if item.id == socketConn.id {
			userData = item
			break
		}
	}

	return userData
}

func (server *SocketServer) disconnect(socketConn SocketConn) {

	var index int
	var userData UserData

	userOnRoom := server.getUser(socketConn)
	fmt.Printf("%s left\n", userOnRoom.UserName)

	for i, item := range server.rooms[userOnRoom.RoomId] {

		if item.id == socketConn.id {
			index = i
			userData = item
			break
		}
	}

	server.rooms[userOnRoom.RoomId] = append(server.rooms[userOnRoom.RoomId][:index], server.rooms[userOnRoom.RoomId][index+1:]...)

	messageBroadcast := map[string]interface{}{
		`userName`: userData.UserName,
		`roomId`:   userData.RoomId,
	}

	b := Broadcast{event: "disconnectUser", message: messageBroadcast, roomId: userData.RoomId, includeCurrentSocket: false}
	server.broadcast(b)

}

func (server *SocketServer) sendMessage(conn net.Conn, socketMessage SocketMessage) {
	response, _ := json.Marshal(socketMessage)
	conn.Write(response)
	conn.Write([]byte(string('\n')))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
