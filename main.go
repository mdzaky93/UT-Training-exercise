package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// func main() {
// fmt.Println("Test")
// fmt.Println("Hello World")

// json.Unmarshal(JsonData)
// fmt.Println(JsonData)

// }
// ------------------------------------------------------------------------
// JSON

// type User struct {
// 	Name string `json:"name"`
// 	Age  string `json:"age"`
// }
// func main()  {
// var JsonString = `{"name": "Dzakiy", "age": "26"}`
// var JsonData = []byte(JsonString)

// var data User
// var err = json.Unmarshal(JsonData, &data)

// if err != nil {
//	fmt.Println("Error Unmarhasilling Json " + err.Error())
//  return
// }

// fmt.Println(data)
// fmt.Println("My name is " + data.Name)
// fmt.Println("My age is " + data.Age)
// }

// -------------------------------------------------------------------------
// IF ELSE
// func main()  {
// number := 8.0
// result, err := calculate(number)

// if err != nil {
// 	fmt.Println(err)
// } else {
// 	fmt.Println(result)
// }
// }

// func calculate(d float64) (float64, error) {
// 	if d == 8 {
// 		return 0, errors.New("dont't 8")
// 	}

// 	cal := d + 10
// 	return cal, nil

// }
// --------------------------------------------------------------------------
// Exercise 1 (+,-,*,/)
// func main() {
// 	x := float64(10)
// 	y := float64(3)
// 	add := add(x, y)
// 	min := min(x, y)
// 	mul := mul(x, y)
// 	div := div(x, y)

// 	fmt.Println("Nilai x =", x)
// 	fmt.Println("Nilai y =", y)
// 	fmt.Println("Penjumlahan ", x, " + ", y, "=", add)
// 	fmt.Println("Pengurangan ", x, " - ", y, "=", min)
// 	fmt.Println("Perkalian ", x, " * ", y, "=", mul)
// 	fmt.Println("Pembagian ", x, " / ", y, "=", div)

// }

// func add(x float64, y float64) float64 {
// 	add := x + y
// 	return add
// }

// func min(x float64, y float64) float64 {
// 	min := x - y
// 	return min
// }

// func mul(x float64, y float64) float64 {
// 	mul := x * y
// 	return mul
// }

// func div(x float64, y float64) float64 {
// 	div := x / y
// 	return div
// }
// ---------------------------------------------------------------------------------
// Exercise 2
// func main() {
// 	var hasil, x, y float32
// 	var z string
// 	fmt.Print("Masukkan bilangan X : ")
// 	fmt.Scanln(&x)
// 	fmt.Print("Masukkan Operator : ")
// 	fmt.Scanln(&z)
// 	fmt.Print("Masukkan bilangan Y : ")
// 	fmt.Scanln(&y)
// 	// hasil = x, z, y
// 	hasil = perhitungan(x, y, z)
// 	fmt.Println("Hasil nya adalah", hasil)
// }

// func perhitungan(x float32, y float32, z string) float32 {

// 	switch z {
// 	case "+":
// 		return x + y
// 		break
// 	case "-":
// 		return x - y
// 		break
// 	case "*":
// 		return x * y
// 		break
// 	case "/":
// 		return x / y
// 	}
// 	return 0
// }
// Exercise 3
// func main() {
// 	a := []string{"foo1", "Bar1", "foo2", "Bar2"}

// 	for i, s := range a {
// 		fmt.Println(i, s)
// 	}
// }
// -----------------------------------------------------------------------------------------------
// Exercise 3
// type Profile struct {
// 	Nama   string `json:"nama"`
// 	Tempat string `json:"alamat"`
// 	Umur   string `json:"umur"`
// }

// func main() {
// 	slice := []Profile{{"Dzakiy", "Jakarta", "26"}}
// 	profile, err := json.Marshal(slice)

// 	if err != nil {
// 		fmt.Println("Error Marshalling Json " + err.Error())
// 		return
// 	}

// 	fmt.Println(string(profile))
// }

// --------------------------------------------------------------------------------------------------

func main() {

	jsonStr := `[{
	"id": "1",
	"name": "user1"
}, {
	"id": "2",
	"name": "user2"
}, {
	"id": "3",
	"name": "user3"
}]`

	type Person struct {
		id   string
		name int
	}
	var people []Person

	var personMap []map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &personMap)

	if err != nil {
		panic(err)
	}

	for _, personData := range personMap {

		// convert map to array of Person struct
		var p Person
		p.id = fmt.Sprintf("", personData["id"])
		p.name, _ = strconv.Atoi(fmt.Sprintf("", personData["name"]))
		people = append(people, p)
	}

	fmt.Println("isSuccess:true")
	for _, value := range personMap {
		fmt.Println("id_", value["id"], "#", "name_", value["name"])
	}

}
