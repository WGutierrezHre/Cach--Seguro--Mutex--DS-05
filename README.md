## Pruebas

### 1. Test de Condiciones de Carrera

```bash
go test -race -v ./registry -run TestRegistry
```
### 2. Benchmark 100% Lecturas
```bash
go test -run=^$ -bench=BenchmarkRegister100 -benchmem ./registry
```
### 3. Benchmark 50/50 Lecturas/Escrituras
```bash
go test -run=^$ -bench=BenchmarkRegister50vs50 -benchmem ./registry
```
### 4. Comparación Completa
```bash
go test -run=^$ -bench=. -benchmem -cpu=1,2,4 ./registry
```
