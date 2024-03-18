package library

type Mahasiswa struct{
	Nama string
	Umur int
}

func (m Mahasiswa) GetMahasiswa(){
	println(m.Nama, m.Umur)
}