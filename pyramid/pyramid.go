package main

import (
	"fmt"
	"strings"
)

func main() {

	var n string //menampung inputan user
	fmt.Scan(&n) //memasukkan inputan user

	result := strings.Split(n, "") //proses pemisahan string percharacter

	for i := range result {
		fmt.Println(result[i] + strings.Repeat("0", len(n)-i-1))
		//strings.repeat untuk menambahkan angka 0 berdasarkan posisi index
	} //proses pengulangan dan penggabungan string yang ada di dalam result, dan
}
