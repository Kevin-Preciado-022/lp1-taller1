package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// : completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)

		// : simular trabajo
		time.Sleep(100 * time.Millisecond) // simula trabajo
		ch <- x * x
	}()
	return ch
}

func main() {
	// : crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(3)
	f2 := asyncCuadrado(4)
	f3 := asyncCuadrado(5)

	// : Opción 1: esperar cada futuro secuencialmente
	fmt.Println("Resultado f1:", <-f1)
	fmt.Println("Resultado f2:", <-f2)
	fmt.Println("Resultado f3:", <-f3)

	// : Opción 2: fan-in (combinar múltiples canales)
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	fanIn := func(channels ...<-chan int) <-chan int {
		out := make(chan int)
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for _, ch := range channels {
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}(ch)
		}

		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	fanInCh := fanIn(asyncCuadrado(6), asyncCuadrado(7), asyncCuadrado(8))
	for result := range fanInCh {
		fmt.Println("Resultado fan-in:", result)
	}
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.

}
