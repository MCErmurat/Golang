package main

import "fmt"

func main() {
	/*var names [3]string //Dizilerin boyutu önceden verilmek zorundadır.
	names[0] = "Mehmet"
	names[1] = "Can"
	names[2] = "Ermurat"
	fmt.Println(names)*/

	/*var names = [3]string{"Mehmet", "Can", "Ermurat"}//Diziyi bu şekilde de tanımlanabilir.
	names[1] = "Enes"//Sonradan diziye farklı bir atama yapılabilir.
	fmt.Println(names)
	//Arraya fazladan eleman verilemez. Eğer verilirse hata çıkar.*/

	/*var names = [4]string{"Mehmet", "Can", "Ermurat", "Serhat"}

	fmt.Println(names[0:2]) //İlk indeksten ikinci indekse kadar ikinci index dahil olmamak üzere alır.*/

	//Slice Kullanımı
	//Dizi ile aynı kullanılır ama boyut belirtmek zorunda değiliz. Bir nevi dinamik dizi denebilir.

	var names = []string{"Mehmet", "Can", "Ermurat"}

	names = append(names, "Serhat") //Append ile diziyi genişletebiliriz.
	fmt.Println(names)
}
