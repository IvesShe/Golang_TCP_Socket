package main

import (
	"fmt"
	"net" // 作網路socket開發時，net包含需要的方法和函數
)

func process(conn net.Conn) {
	// 我們循環的介紹客戶端發送的數據

	// 關閉conn
	defer conn.Close()

	for {
		// 創建一個新的切片
		buf := make([]byte, 1024)

		// 1.等待客戶端通過conn發送信息
		// 2.如果客戶端沒有write[發送]，那麼協程就阻塞在這裡
		// fmt.Printf("\n服務器在等待客戶端%s 發送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if string(buf[:n]) == "exit" {
			fmt.Print(string(conn.RemoteAddr().String() + " : 已離線\n"))
			break
		}
		if err != nil {
			fmt.Println("服務器的Read err = ", err)
		}
		// 3.顯示客戶端發送的內容到服務器的終端
		fmt.Print(string(conn.RemoteAddr().String() + " : "))
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服務器開始監聽...")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}

	// 延時關閉listen
	defer listen.Close()

	// 循環等待客戶端來連接
	for {
		// 等待客戶端連接
		fmt.Println("等待客戶端來連接.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
		} else {
			fmt.Printf("Accept() suc con = %v\n", conn)
			fmt.Printf("客戶端ip = %v\n", conn.RemoteAddr().String())
		}

		// 這裡準備一個協程，為客戶端服務
		go process(conn)
	}

	//fmt.Printf("listen suc = %v\n", listen)
}
