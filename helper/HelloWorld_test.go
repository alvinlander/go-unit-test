package helper

import (
	"fmt"
	"testing"
)
/**
How To Run :
1. Masuk ke package dan ketik di terminal "go test"
2. Untuk melihat nama fungsi gunakan -v "go test -v"
3. Untuk menjalankan test fungsi tertentu tambahkan -run "go test -v -run TestNamaFungsi"
4. Untuk running semua unit test di semua package sekaligus "go test ./..."
*/
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Lans")
	if(result != "Hello Lanss") {
		/**
		Test Failed
		Fail() jika ada gagal akan menjalankan baris kode selanjutnya
		*/
		t.Fail()
	}
	//lanjut kesini
	fmt.Println("TestHelloWorld Done")
}
func TestHelloWorldLans(t *testing.T) {
	result := HelloWorld("Lans")
	if(result != "Hello Lanss") {
		/**
		Test Failed
		FailNow() akan mengagalkan fungsi test dan melanjutkan ke test selanjutnya
		*/
		t.FailNow()
	}
	//tidak akan lanjut kesini
	fmt.Println("TestHelloWorldLans Done")
}
func TestHelloWorldLander(t *testing.T) {
	result := HelloWorld("Lans")
	if(result != "Hello Lanss") {
		/**
		Test Failed
		Error() akan memberikan log error dan memanggil fungsi Fail
		*/
		t.Error("Result Must Be Lans")
	}
	//lanjut kesini
	fmt.Println("TestHelloWorldLander Done")
}
func TestHelloWorldAlvin(t *testing.T) {
	result := HelloWorld("Lans")
	if(result != "Hello Lanss") {
		/**
		Test Failed
		Fatal() akan memberikan log error dan memanggil fungsi FailNow
		*/
		t.Fatal("Result Must Be Lans")
	}
	//tidak akan lanjut kesini
	fmt.Println("TestHelloWorldAlvin Done")
}
/**
Note :
Lebih baik gunakan Error atau Fatal dibanding Fail dan FailNow
*/