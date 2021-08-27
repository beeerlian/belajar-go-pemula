package main

import "fmt"

type Mahasiswa struct {
	ID        int
	NIM       string
	FirstName string
	LastName  string
	IPK       float32
}

//method
func (mahasiswa Mahasiswa) display() (data string) {
	return fmt.Sprintf("NIM : %s \nNama : %s %s", mahasiswa.NIM, mahasiswa.FirstName, mahasiswa.LastName)
}

//function
func display(mahasiswa Mahasiswa) (data string) {
	return fmt.Sprintf("NIM : %s \nNama : %s %s", mahasiswa.NIM, mahasiswa.FirstName, mahasiswa.LastName)
}

type Kelas struct {
	ID           int
	NamaKelas    string
	Pengajar     Dosen
	PesertaDidik []Mahasiswa
}

//method
func (kelas Kelas) display() (data string) {
	return fmt.Sprint("Kelas : ", kelas.NamaKelas, "\nPengajar : ", kelas.Pengajar, "\nPeserta Didik : ", kelas.PesertaDidik)
}

//function
func displayDataKelas(kelas Kelas) (data string) {
	return fmt.Sprint("Kelas : ", kelas.NamaKelas, "\nPengajar : ", kelas.Pengajar, "\nPeserta Didik : ", kelas.PesertaDidik)
}

type Dosen struct {
	NIP       string
	FirstName string
	LastName  string
}

func main() {
	mhs1 := Mahasiswa{
		ID:        1,
		NIM:       "1197050094",
		FirstName: "Naufal",
		LastName:  "Berlian",
		IPK:       2.50,
	}
	mhs2 := Mahasiswa{
		ID:        2,
		NIM:       "1197050094",
		FirstName: "Ujang",
		LastName:  "Imam",
		IPK:       2.50,
	}
	mhs3 := Mahasiswa{
		ID:        3,
		NIM:       "1197050094",
		FirstName: "Soniawan",
		LastName:  "Abbs",
		IPK:       2.50,
	}

	dosen1 := Dosen{
		NIP:       "11109283",
		FirstName: "Farhan",
		LastName:  "Akhdan",
	}

	kelas1 := Kelas{
		ID:           2,
		NamaKelas:    "B",
		Pengajar:     dosen1,
		PesertaDidik: []Mahasiswa{mhs1, mhs2, mhs3},
	}

	kelas2 := Kelas{
		ID:        1,
		NamaKelas: "A",
		Pengajar: Dosen{
			NIP:       "11109283",
			FirstName: "Farhan",
			LastName:  "Akhdan",
		},
		PesertaDidik: []Mahasiswa{
			{
				ID:        1,
				NIM:       "1197050094",
				FirstName: "Naufal",
				LastName:  "Berlian",
				IPK:       2.50,
			},
			{
				ID:        2,
				NIM:       "1197050094",
				FirstName: "Ujang",
				LastName:  "Imam",
				IPK:       2.50,
			},
			{
				ID:        3,
				NIM:       "1197050094",
				FirstName: "Soniawan",
				LastName:  "Abbs",
				IPK:       2.50,
			},
		},
	}

	fmt.Println("----------------------------------")

	displayMethod := kelas1.display()
	fmt.Println(displayMethod)

	displayFunc := displayDataKelas(kelas2)
	fmt.Print(displayFunc)
}
