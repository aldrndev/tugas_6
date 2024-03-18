package main

import "tugas_6/library"

func main() {
    var mahasiswa = library.Mahasiswa{
        Nama: "Daffa",
        Umur: 20,
    }
    mahasiswa.GetMahasiswa() 
}