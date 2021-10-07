package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
How To Run :
1. Masuk ke package dan ketik di terminal "go test"
2. Untuk melihat nama fungsi gunakan -v "go test -v"
3. Untuk menjalankan test fungsi tertentu tambahkan -run "go test -v -run TestNamaFungsi"
4. Untuk running semua unit test di semua package sekaligus "go test ./..."
*/
func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Lans")
	if result != "Hello Lans" {
		/**
		Test Failed
		Fail() jika ada gagal akan menjalankan baris kode selanjutnya
		*/
		t.Fail()
	}
	//lanjut kesini
	fmt.Println("TestHelloWorld Done")
}
func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Lans")
	if result != "Hello Lans" {
		/**
		Test Failed
		FailNow() akan mengagalkan fungsi test dan melanjutkan ke test selanjutnya
		*/
		t.FailNow()
	}
	//tidak akan lanjut kesini
	fmt.Println("TestHelloWorldLans Done")
}
func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Lans")
	if result != "Hello Lans" {
		/**
		Test Failed
		Error() akan memberikan log error dan memanggil fungsi Fail
		*/
		t.Error("Result Must Be Lans")
	}
	//lanjut kesini
	fmt.Println("TestHelloWorldLander Done")
}
func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Lans")
	if result != "Hello Lans" {
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

/**
Assert
Seperti menggunakan Fail(), jika ada error maka akan menjalankan
baris kode selanjutnya
*/
func TestHelloWorldAssert(t *testing.T) {
	res := HelloWorld("Lander")
	assert.Equal(t, "Hello Lander", res, "Result must be Hello Lander")
	fmt.Println("HelloWorldAssert Test is done")
}

/**
Require
Seperti menggunakan FailNow(), jika ada error maka fungsi
berhenti dijalankan
*/
func TestHelloWorldRequire(t *testing.T) {
	res := HelloWorld("Lans")
	require.Equal(t, "Hello Lans", res, "Result must be Hello Lans")
	fmt.Println("HelloWorldRequire Test is done")
}

// Skip Test

/**
Skip Test
Merupakan fitur untuk meng-skip test agar tidak dijalankan
*/
func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Test cannot be run on Mac OS")
	}
	res := HelloWorld("Lans")
	require.Equal(t, "Hello Lans", res, "Result must be Hello Lans")
	fmt.Println("HelloWorldRequire Test is done")
}

// Test Main (Before and After Test) hanya berjalan pada package yg sama

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")
	m.Run()
	//after
	fmt.Println("After Unit Test")
}

/**
How To Run :
1. go test -v -run namafile (Run semua sub test)
2. go test -v -run namafile/namatest (Run hanya sub test nya saja)
*/
func TestHelloWorldSubTest(t *testing.T) {
	t.Run("Alvin", func(t *testing.T) {
		res := HelloWorld("Lans")
		require.Equal(t, "Hello Lans", res, "Result must be Hello Lans")
		fmt.Println("HelloWorldRequire Test is done")
	})
	t.Run("Lander", func(t *testing.T) {
		res := HelloWorld("Lans")
		require.Equal(t, "Hello Lans", res, "Result must be Hello Lans")
		fmt.Println("HelloWorldRequire Test is done")
	})
}

// Table Test
/**
Digunakan untuk menghindari duplicate fungsi yang isinya sama
*/
func TestHelloWorldTable(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Alvin",
			request:  "Alvin",
			expected: "Hello Alvin",
		},
		{
			name:     "Lander",
			request:  "Lander",
			expected: "Hello Lander",
		},
		{
			name:     "Lans",
			request:  "Lans",
			expected: "Hello Lans",
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.request)
			require.Equal(t, test.expected, res)
		})
	}
}

/**

**Benchmark**
How To Run:
1. go test -v -bench=. (Running unit test dan benchmark)
2. go test -v -run=TestTidakAda -bench=. (hanya running benchmark)
3. go test -v -run=TestTidakAda -bench=BenchmarkHello (hanya running fungi benchmark hello)
4. go test -v -run=TestTidakAda -bench=. ./... (running semua benchmark yang ada di root)
*/
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lans")
	}
}
func BenchmarkHelloWorldLander(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lander")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Lander", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lans")
		}
	})
	b.Run("Alvin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Alvin")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Lans)",
			request: "Lans",
		},
		{
			name:    "HelloWorld(Lander)",
			request: "Lander",
		},
		{
			name:    "HelloWorld(Alvin)",
			request: "Alvin",
		},
		{
			name:    "HelloWorld(Alvin Lander)",
			request: "Alvin Lander",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
