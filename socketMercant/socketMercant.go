package main

import "fmt"

func main() {
	var n int    //var banyak kaos kaki
	fmt.Scan(&n) //inputan dari user banyaknya kaos kaki

	socks := make([]int, 20) //array yang akan menampung inputan kaos kaki.

	for i := 0; i < n; i++ { //pengulangan yang bertujuan agar user terus mengisi kaos kaki
		var temp int    //variable untuk menampung inputan user.
		fmt.Scan(&temp) //inputan dari user untuk data/socks nya

		socks[temp]++ //memasukkan nilai inputan dari user kedalam index dari array.
		//jadi inputan user akan masuk langsung ke index nya, apabila user menginput
		// 1 maka nilai di index 1 akan bertambah menjadi 1. apabila input 1 lagi,
		// nilai di index 1 akan berubah menjadi 2.
	}

	count := 0 //nilai awal kaos kaki yg sepasang.

	for _, value := range socks {

		count = count + value/2
		// proses pengulangan untuk mendapatkan kaos kaki yang sepasang dalam suatu index,
	}

	fmt.Println("output = ", +count)
}
