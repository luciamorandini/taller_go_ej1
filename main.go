package main

import "fmt"

func main() {

	votos_recibidos := []int{2, 1, 5, 3, 4, 3, 4, 1, 1, 4}
	fmt.Println(votos_recibidos)

	map_votos := map[int]int{}
	votos_positivos := 0
	votos_negativos := 0

	for _, v := range votos_recibidos {
		map_votos[v]++
		if v >= 4 {
			votos_positivos++
		} else if v <= 2 {
			votos_negativos++
		}
	}

	// Mostrar resultados
	fmt.Println("Cantidad de votos por puntaje:")
	//for puntaje, cantidad := range map_votos {
	//fmt.Printf("Puntaje %d: %d votos\n", puntaje, cantidad)
	//}

	for i := 1; i <= 5; i++ {
		fmt.Printf("Puntaje %d: %d votos\n", i, map_votos[i])
	}

	fmt.Printf("Cantidad de votos positivos: %d\n", votos_positivos)
	fmt.Printf("Cantidad de votos negativos: %d\n", votos_negativos)

	if votos_positivos > votos_negativos {
		fmt.Println("¡Buen resultado!")
	} else {
		fmt.Println("Resultado mejorable :(")
	}

}
