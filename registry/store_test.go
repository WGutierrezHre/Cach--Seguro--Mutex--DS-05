package registry

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func GenerarServicios() map[string]string {
	services := make(map[string]string, 500)

	r := rand.New(rand.NewSource(42))
	for i := 0; i < 500; i++ {
		service := fmt.Sprintf("Servicio -- %d", i)
		ip := fmt.Sprintf("192.168.%d.%d", r.Intn(256), r.Intn(256))
		services[service] = ip
	}
	return services
}

// TEST DE CONDICIONES DE CARRERA
func TestRegistry(t *testing.T) {
	t.Parallel()

	reg := NewRegistry()
	services := GenerarServicios()
	for service, ip := range services {
		reg.Register(service, ip)
	}

	var wg sync.WaitGroup

	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			r := rand.New(rand.NewSource(int64(1000 + id)))
			for j := 0; j < 2000; j++ {
				service := fmt.Sprintf("Servicio -- %d", r.Intn(500))
				reg.Lookup(service)
			}

		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			r := rand.New(rand.NewSource(int64(2000 + id)))
			for j := 0; j < 2000; j++ {
				service := fmt.Sprintf("Servicio -- %d", r.Intn(500))
				ip := fmt.Sprintf("10.0.%d.%d", r.Intn(256), r.Intn(256))
				reg.Register(service, ip)
			}

		}(i)
	}
	wg.Wait()
}

// BENCHMARK 100% DE LECTURAS
func BenchmarkRegister100(b *testing.B) {
	reg := NewRegistry()
	services := GenerarServicios()
	for service, ip := range services {
		reg.Register(service, ip)
	}

	b.ResetTimer()
	b.RunParallel(
		func(pb *testing.PB) {
			r := rand.New(rand.NewSource(1))
			for pb.Next() {
				service := fmt.Sprintf("Servicio -- %d", r.Intn(500))
				reg.Lookup(service)
			}
		})
}

// BENCHMARK 50% DE LECTURAS, 50% DE ESCRITURAS
func BenchmarkRegister50vs50(b *testing.B) {
	reg := NewRegistry()
	services := GenerarServicios()
	for service, ip := range services {
		reg.Register(service, ip)
	}

	b.ResetTimer()
	b.RunParallel(
		func(pb *testing.PB) {
			r := rand.New(rand.NewSource(2))
			i := 0
			for pb.Next() {
				service := fmt.Sprintf("Servicio -- %d", r.Intn(500))
				if i%2 == 0 {
					reg.Lookup(service)
				} else {
					ip := fmt.Sprintf("10.0.%d.%d:9000", r.Intn(256), r.Intn(256))
					reg.Register(service, ip)
				}
				i++
			}
		})
}
