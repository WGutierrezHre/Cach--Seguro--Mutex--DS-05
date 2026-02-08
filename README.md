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

## Starvation
Ocurre cuando una goroutine no puede ejecutarse durante un periodo prolongado porque otras tienen prioridad. 

En el caso de ```sync.RWMutex```, cuando una operación de escritura solicita el bloqueo mediante ```Lock()```, se impide la entrada de nuevos lectores que intenten adquirir ```RLock()```, esta restricción se mantiene hasta que la operación de escritura finaliza.

De esta forma, los lectores se bloquean temporalmente para evitar que el escritor quede esperando indefinidamente (writer starvation), garantizando así tanto el progreso de la escritura como la consistencia de los datos compartidos.
