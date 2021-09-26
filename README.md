# IntermediateGo
## TESTING
You can create test of ***xyz.go*** files just creating another file called ***xyz_test.go*** as a convention 

There you can create tables with test cases:

	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

To test Run:

	go test

### Code Coverage
Run the test with the flag ***-coverprofile=coverage.out*** lets you know which parts of the code have been tested and which have not:

	$ go test -coverprofile=coverage.out

To read and visualize the metrics use: 
##### Metrics in terminal

	$ go tool cover -func=coverage.out
##### o para ver resumen en el navegador

	$ go tool cover -html=coverage.out  

### Profiling
Of you want to test the performance of your code you should use profiling.

To see the cpu usage of tested code use:

	$ go test -cpuprofile=cpu.out

To see the cpu usage sumarized use:
	
	$ go tool pprof cpu.out

!!![] To use pprof you may install it using: $ sudo apt install graphviz 

Dentro de pprof escribimos top para ver como se han comportado los programas en nuestro test

	(pprof) top

Además, dentro de pprof podemos inspeccionar el tiempo promedio de ejecución de cada línea de una función, usando el comando list <nombre_funcion>

	(pprof) list Fibonacci

Tambien podemos ver el reporte del promedio de ejecución:
##### en el navegador usando web

	(pprof) web
##### o exportarlo en pdf usando pdf
	(pprof) pdf

Para salir de (pprof) puedes usar **quit** o **Ctrl + D**


