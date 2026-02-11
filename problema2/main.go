package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular tareas que toman tiempo con time.Sleep y comparar
// ejecución secuencial vs concurrente midiendo la duración total.


func tarea(id int, dur time.Duration) {
	fmt.Printf("[tarea %d] iniciando, dur=%v\n", id, dur)
	time.Sleep(dur)
	fmt.Printf("[tarea %d] finalizada\n", id)
}

func secuencial(durs []time.Duration) time.Duration {
	inicio := time.Now()

	for i, d := range durs {
		tarea(i, d)

	}
	return time.Since(inicio)
}

func concurrente(durs []time.Duration) time.Duration {
	inicio := time.Now()
	var wg sync.WaitGroup
	
	for i, d := range durs {
		wg.Add(1)
		go func(id int, dur time.Duration) {
			defer wg.Done()
			tarea(id, dur)
		}(i, d)

	}
	wg.Wait()
	return time.Since(inicio)
}

func main() {
	

	d1 := secuencial([]time.Duration{700 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second})
	fmt.Println("Duración SEC:", d1)	

	d2 := concurrente([]time.Duration{700 * time.Millisecond, 500 * time.Millisecond, 1 * time.Second})
	fmt.Println("Duración CONC:", d2)

	fmt.Println("Nota: la ejecución concurrente debería ser ~max(durs). Cambia valores y observa.")
}
