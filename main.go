package main

import (
	"ejercicio2/registry"
	"fmt"
	"sync"
	"time"
)

func main() {
	reg := registry.NewRegistry()
	reg.Register("gateway", "192.168.1.1")
	reg.Register("auth-service", "192.168.1.20")
	reg.Register("db-service", "192.168.1.30")

	var wg sync.WaitGroup

	// Simular 100 servicios leyendo simultáneamente
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			addrIP, ok := reg.Lookup("gateway")
			if ok {
				fmt.Printf("Lector %d: Gateway -- %s\n", id, addrIP)
			}

			addrIP, ok = reg.Lookup("auth-service")
			if ok {
				fmt.Printf("Lector %d: Auth-Service -- %s\n", id, addrIP)
			}

		}(i)
	}
	// Simular cambios de IP
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		fmt.Println("\n--- Actualizando IP del GATEWAY ---")
		reg.Register("gateway", "192.168.1.11")

		time.Sleep(10 * time.Millisecond)
		fmt.Println("--- Actualizando IP del Auth-Service ---")
		reg.Register("auth-service", "192.168.1.50")
	}()
	time.Sleep(20 * time.Millisecond)
	for i := 100; i < 150; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			addrIP, ok := reg.Lookup("gateway")
			if ok {
				fmt.Printf("Lector %d: Gateway(ACT) -- %s\n", id, addrIP)
			}

			addrIP, ok = reg.Lookup("auth-service")
			if ok {
				fmt.Printf("Lector %d: Auth-Service(ACT) -- %s\n", id, addrIP)
			}

		}(i)
	}
	wg.Wait()

}
