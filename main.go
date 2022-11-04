package main

import (
	"fmt"
)

func main() {
	carBrands := []string{"Toyota", "Subaru", "Mercedes", "Lexus", "Mazda", "Lammbo"} // slice w postaci STRINGa (taki "zbiorek" :)

	fmt.Println("Moja lista aut: ")
	for idx, car := range carBrands { // idx -> to dodatkowa zmienna, oznaczająca "iterację" "pętli"
		fmt.Printf(" %d = > %s\n", idx, car)
	}

	// i kod na samą zawartośc, bez iteracji elementów :
	/*
		fmt.Println("Moja lista aut: ")
		for _, car := range carBrands { // idx -> to dodatkowa zmienna, oznaczająca "iterację" "pętli". tu z niej nie korzystamy.
			fmt.Printf(" %s\n", car)
		} //
	*/

}
