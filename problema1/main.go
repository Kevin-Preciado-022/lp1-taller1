package main

import (
	"fmt"
	"sync"
)

// Objetivo: Lanzar varias goroutines que imprimen mensajes y esperar a que todas terminen.
// : Completa los pasos marcados con  para entender goroutines y WaitGroup.

func worker(id int, veces int, wg *sync.WaitGroup) {
	// : asegurar que al finalizar la función se haga wg.Done()
	defer wg.Done()

	for i := 1; i <= veces; i++ {
		fmt.Printf("[worker %d] hola %d\n", id, i)
		// : dormir un poco para simular trabajo (p. ej. 100–300 ms)

	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 3

	// : cambiar estos parámetros y observar el intercalado de salidas
	// numGoroutines
	// veces

	// : lanzar varias goroutines, sumar al WG y esperar con wg.Wait()
	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1)
		go worker(id, 3, &wg)

	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	fmt.Println("Listo: todas las goroutines terminaron.")
}
