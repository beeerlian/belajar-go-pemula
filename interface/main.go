package main

import "fmt"

type Product interface {
	PotongHarga(besarPotongan int)
}

type Sepatu struct {
	Nama  string
	Harga int
}

func (sepatu *Sepatu) PotongHarga(besarPotongan int) {
	sepatu.Harga -= sepatu.Harga * besarPotongan / 100
}

type Baju struct {
	Nama  string
	Harga int
}

func (baju *Baju) PotongHarga(besarPotongan int) {
	baju.Harga -= baju.Harga * besarPotongan / 100
}
func main() {
	baju := Baju{Nama: "Screamous", Harga: 450000}
	Discount(&baju, 20)
	fmt.Println(baju.Harga)

}

func Discount(product Product, besarPotongan int) {
	product.PotongHarga(besarPotongan)
}
