package main

import ("fmt"
		"os"
		"sort")


func main(){
	var (
		opc int
		abc string
		slide [] string
		
	)
	
	fmt.Println("Cuantos quieres registrar? ")
	fmt.Scan(&opc)

	for i := 0 ; i < opc ; i++ {
		fmt.Println("registro numero ", i + 1)
		fmt.Println("Escribe el string a registrar: ")
		fmt.Scan(&abc)
		slide = append(slide, abc + "\n")
		
	}
	fmt.Println("\n\n DESORDENADO")
	fmt.Println(slide)
	
	sort.Strings(slide)
	

	//crea el archivo ascendente
	
	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("\n\n ORDENADO ASCENDENTE")

	for _, v :=  range slide{
		file.WriteString(v)
		fmt.Println(v)
	}
	
	sort.Sort(sort.Reverse(sort.StringSlice(slide)))
	

	//crea el archivo descendente
	file, err2 := os.Create("descendente.txt")
	if err2 != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	fmt.Println("\n\n ORDENADO DESCENDENTE")

	for _, v :=  range slide{
		file.WriteString(v)
		fmt.Println(v)
	}
	
	

}