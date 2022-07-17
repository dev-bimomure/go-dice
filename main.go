package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jumlahPemain := 3
	jumlahDadu := 4
	mainDadu(jumlahPemain, jumlahDadu)
}

func mainDadu(jumlahPemain int, jumlahDadu int) {
	Pemain := make(map[int][]int)
	for i := 0; i < jumlahPemain; i++ {
		Hasil := LemparDadu(jumlahDadu)
		Nama := 1 + i
		Pemain[Nama] = Hasil
	}

	fmt.Printf("--------------Mulai-Permainan \n")
	for key, value := range Pemain {
		fmt.Printf("Pemain %d, : %d \n", key, value)
	}
	for {
		result := CheckDadu(Pemain)
		Pemain = result

		fmt.Printf("--------------Check-Dadu--------------------- \n")
		for key, value := range result {
			fmt.Printf("Pemain %d, : %d \n", key, value)
		}

		daduNol := 0
		for _, Dadu := range result {
			if len(Dadu) == 0 {
				daduNol = daduNol + 1
			}
		}

		if jumlahPemain-daduNol == 1 {
			break
		}

		for i, Dadu := range result {
			Hasil := LemparDadu(len(Dadu))
			Pemain[i] = Hasil
		}

		fmt.Printf("--------------Lempar-Dadu--------------------- \n")
		for key, value := range Pemain {
			fmt.Printf("Pemain %d, : %d \n", key, value)
		}

	}

}

func LemparDadu(dadu int) []int {
	var TempatDadu []int
	for j := 0; j < dadu; j++ {
		hasil := KocokDadu()
		TempatDadu = append(TempatDadu, hasil)
		time.Sleep(1 * time.Microsecond)
	}

	return TempatDadu
}

func KocokDadu() int {
	max := 7
	min := 1
	rand.Seed(time.Now().UTC().UnixNano())
	randomNum := min + rand.Intn(max-min)
	return randomNum
}

func CheckDadu(data map[int][]int) map[int][]int {
	newData := data
	hasil := make(map[int][]int)
	var tempDadu1 []int
	for j, tempatDadu := range newData {
		TempatDaduNew := tempatDadu
		if len(tempDadu1) > 0 {
			for _, item := range tempDadu1 {
				TempatDaduNew = append(TempatDaduNew, item)
			}
			tempDadu1 = []int{}
		}
		for i, dadu := range tempatDadu {
			if dadu == 6 {
				TempatDaduNew = hapusDadu(TempatDaduNew, i)

			}
			if dadu == 1 {
				tempDadu1 = append(tempDadu1, 1)
			}
		}
		hasil[j] = TempatDaduNew

	}

	return hasil
}

func hapusDadu(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return nil
	}
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
