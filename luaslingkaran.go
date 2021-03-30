package main
import "fmt"

func luaslingkaran(jarijari float64) float64 {
	return 3.14159*jarijari*jarijari
}

func main() {
	var x float64
	for true {
		fmt.Printf("Jari-jari lingkaran:")
		fmt.Scan(&x)
		fmt.Println(luaslingkaran(x))
	}
}