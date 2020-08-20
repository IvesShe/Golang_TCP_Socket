package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	//fmt.Println("conn 成功 = ", conn)
	fmt.Println("連接成功，您可以開始輸入訊息了!!")

	// 功能一：客戶端可以發送單行數據，然後就退出
	// os.Stdin 代表標準輸入[終端]
	reader := bufio.NewReader(os.Stdin)

	for {
		// 從終端讀取一行用戶輸入，並準備發送給服務器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nreadString err = \n", err)
		}

		// 如果用戶輸入的是 exit 就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			_, err = conn.Write([]byte("exit"))
			if err != nil {
				fmt.Println("conn.Write err = ", err)
			}
			fmt.Println("客戶端退出..")
			break
		}
		line += ("\n")
		// 再將line發送給服務器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err = ", err)
		}
		//fmt.Printf("客戶端發送了 %d 字節的數據，並退出!!", n)
		//fmt.Printf("客戶端發送了 %d 字節的數據...\n", n)
	}
}
