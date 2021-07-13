package go_say_hello

/*
* Membuat Module
* Untuk membuat module kita bisa gunakan perintah "go mod ini <nama_module>
* Biasanya programmer Go memakai project git sebagai nama modulenya (misal github.com/gerakanamalindonesia/go-module_example)
* Kemudian kita bis buat file yang nanti akan dirilis di githubnya, kemudain push ke github
* Untuk release kita bisa gunakan v (umumnya untuk rilis kigunakan v diawal) --> git tag v1.0.0,
* Tinggal push deh --> git push origin v1.0.0
*/

/*
* Upgrade Module
* Jika kita ingin mengupdate modul kita maka kita hanya perlu membuat tag baru lalu kita push
*/

/*
* Major Upgrade
* Sebenarnya hal ini sebisa mungkin dihindari, karena dapat merusak kode
* Namun jika memang terpaksa dan harus melakukan perusakan code (artinya adalah major upgrade) maka
  strategi yang dapat kita gunakan adalah mengubah nama module
* Misal seperti ini "github.com/gerakanamalindonesia/go-module_example" kita ubah menjadi "github.com/gerakanamalindonesia/go-module_example/v2",
  harapannya dalah ketika ada yang mengakses "github.com/gerakanamalindonesia/go-module_example" maka mereka akan mendapakan yagn v1
*/

func SayHello(address string) string {
	// Sebelum diupgrade module-nya
	//return "Hello gan"

	// Setelah upgrade module
	return (address)
}