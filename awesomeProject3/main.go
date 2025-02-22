package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Lutfen secim yapiniz")
	fmt.Println("1-Toplama\n" +
		"2-Cikarma\n" +
		"3-Carpma\n" +
		"4-Bolme\n" +
		"5-ModAlma")

	scanner.Scan()

	secim := scanner.Text()

	fmt.Print("1.sayiyi giriniz:")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("2.sayiyi giriniz:")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	if secim == "1" {
		fmt.Printf("Toplama:%f", sayi1+sayi2)

	} else if secim == "2" {
		fmt.Printf("Cikarma:%f", sayi1-sayi2)

	} else if secim == "3" {
		fmt.Printf("Carpma:%f", sayi1*sayi2)

	} else if secim == "4" {
		fmt.Printf("Bolme:%f", sayi1/sayi2)

	} else if secim == "5" {
		modAlma := int(sayi1) % int(sayi2)
		fmt.Printf("Mod Alma: %d", modAlma)

	} else {
		fmt.Println("Hatali giris yaptiniz...")
	}

}
