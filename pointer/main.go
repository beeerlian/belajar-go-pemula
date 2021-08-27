package main

import "fmt"

func main() {
	//Pointer adalah variabel yang menyimpan adress suatu variabel di memory

	variabel := 5
	//pointer menyimpan adress variabel
	pointer := &variabel

	fmt.Println(variabel)
	fmt.Println(pointer)

	//untuk mengetahui nilai yg disimpannya, bisa diakses melalui pointer,
	//bis diakses melalui pointer dengan menggunakan refference *
	fmt.Println(*pointer) //hasil = 5

	//pointer dan variabel akan memiliki keterikatan
	//sehingga, kita bisa merubah nilai variabel, dengan me-assign nilai ke pointernnya

	*pointer = 10
	fmt.Println(*pointer)
	fmt.Println(variabel) // hasil = 10
}
