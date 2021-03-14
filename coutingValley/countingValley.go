package main

import "fmt"

func main() {

	// var n int
	// fmt.Scan(&n)

	// var steps []string

	// for i := 0; i < n; i++ {
	// 	var step string
	// 	fmt.Scanf("%s", &step)
	// 	steps = append(steps, step)
	// }

	//inputan
	var n = 10 // seharusny total inputan user

	steps := []string{"U", "D", "D", "U", "U", "D", "D", "U", "U", "D"} // seharusny inputan user
	var valley = 0                                                      //nilai awal valley
	var currentPos = 0                                                  //posisi awal daratan.

	for i := 0; i < n; i++ {
		if steps[i] == "U" {
			currentPos++ //posisi bertambah
		} else {
			currentPos-- //posisi berkurang
		}
		if currentPos == 0 && steps[i] == "U" { //ketika posisi di darat, dan dia habis dari dasar laut, maka total valley +1
			valley++
		}
	}
	fmt.Println(valley)
}
