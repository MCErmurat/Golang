package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*var productName string // Değişkenler bu şekilde tanımlanır, türünü belirtmek zorundayız.
	var quantity int
	var discount float32
	var inInStock bool

	productName = "Golang Dersleri" //Değişkeni tanımladıktan sonra kullanmazsak sorun çıkartır.
	quantity = 10
	discount = 0.37
	inInStock = true

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity)) // reflect paketinin TypeOf metodu içerisine verdiğimiz değişkenin türünü yazdırır.
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(inInStock, reflect.TypeOf(inInStock)) */

	/*var productName string = "Golang Dersleri" //Eğer istersek aynı satır içerisinde de değişkeni tanımlayıp kullanabiliriz.
	fmt.Println(productName)

	var productName2 = "Golang Dersleri" //Değişkenin türünü belirmesekde sağ tarafta bir string gördüğü için go kendisi değişkenin türünü string olarak ayarlar.
	fmt.Println(productName2, reflect.TypeOf(productName2))*/

	/*productName:="Golang Dersleri"//Bu şekilde de var'ı kullanmadan da bir değişken tanımlayıp kullanabiliriz.
	fmt.Println(productName)*/

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = 0.37
	isInStock = true

	//fmt.Println("Product Name: ", productName, "Quantity: ", quantity, "Discount: ", discount, "isInStock: ", isInStock)
	fmt.Printf("Product Name: %s, Quantity: %d , Discount: %f, isInStock: %t \n", productName, quantity, discount, isInStock) //Yukarıdaki kullanımdan tek farkı formatlı bir şekilde yazılmış olmasıdır.
	// %s %b %d %f %t vs. gibi kullanımlar veri türlerinin formatlanmasında kullanılıyor. Hepsinin kullanımı ve açıklaması go'nun fmt paketinin dökümanında var.
	// %v ile de default bir şekilde kullanabiliriz. %v diğer bütün hepsinin yerine yazılabilir.
	fmt.Println(productName, reflect.TypeOf(productName))
}
