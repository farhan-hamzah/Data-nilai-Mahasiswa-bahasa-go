package main
import "fmt"

func main(){
	var jumMahasiswa int
	fmt.Scan(&jumMahasiswa)
	dataMahasiswa(&jumMahasiswa)

}
func dataMahasiswa(jumMahasiswa *int){
	var mahasiswa string
	var i int
	var nilaiAkhir float64
	var uts, uas, tugas float64
	for i = 0; i < *jumMahasiswa; i++{
		fmt.Print("Nama: ")
		fmt.Scan(&mahasiswa)
		fmt.Print("Nilai UTS, UAS, Tugas: ")
		fmt.Scan(&uts, &uas, &tugas)
		nilaiAkhir = (uts * 0.3) + (uas * 0.4) + (tugas * 0.3)
		daftarNilai(mahasiswa, nilaiAkhir, i)
		nilaiAkhir = 0
	}
}
func daftarNilai(mahasiswa string, nilaiAkhir float64, jumMahasiswa int){
	var predikat string

	if nilaiAkhir >= 80{
		predikat = "A"
	}else if nilaiAkhir <= 79 && nilaiAkhir >= 70{
		predikat ="B"
	}else if nilaiAkhir <= 69 && nilaiAkhir >= 60{
		predikat = "C"
	}else if nilaiAkhir <= 59 && nilaiAkhir >= 50{
		predikat = "D"
	}else{
		predikat = "E"
	}
	fmt.Println("Daftar Nilai Mahasiswa : ", jumMahasiswa+1)
	fmt.Println("--------------------------------------")
	fmt.Println("Nama :",mahasiswa)
	fmt.Println("Nilai Akhir :", nilaiAkhir)
	fmt.Println("Predikat :", predikat)
	fmt.Println("--------------------------------------")

}