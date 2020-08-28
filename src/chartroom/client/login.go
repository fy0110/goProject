package main

import (
	"chartroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(userID int, userPWD string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println(err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userID
	loginMes.UserPwd = userPWD
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes []byte
	binary.BigEndian.PutUint32(bytes[:], pkgLen)

	n, err := conn.Write(bytes)
	if n != 4 || err != nil {
		fmt.Println("发送长度失败！", err)
	}
	// conn.Write(data)
	return nil
}
