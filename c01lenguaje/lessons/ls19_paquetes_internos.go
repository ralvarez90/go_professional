package lessons

import (
	"c01lenguaje/tools"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// PAQUETES INTERNOS
//
// 1. time
// Contiene funcionalidades que nos permiten trabajar con fechas, horas y duraciones. Es esencial
// para cualquier aplicación que necesite medir el tiempo, programar tareas, trabajar con zonas
// horarias o formatear fechas.
//
// Los objetos Time tienen métodos que nos permiten acceder a los campos de dicho string. También
// podemos comparar fechas obteniendo un booleano con los métodos Before, After e Equal
//
// Podemos obtener diferencias entre fechas usando el método Sub.
//
// 2. os
// El paquete os en Go proporciona una interfaz independiente de la plataforma para interactuar
// con el sistema operativo. Ofrece funciones para realizar operaciones comunes como:
// - Manipulación de archivos
// - Control de procesos
// - Manejo de Variables de Entorno
//
// 3. log
// El paquete log en Go proporciona una forma sencilla de agregar registro (logging) a
// tus aplicaciones.  Te permite escribir mensajes de registro en la salida estándar
// (consola), en la salida de error estándar o en un archivo.
//
// log.Println
// Imprime mensaje en consola agregando más información
//
// log.Fatal
// Muestra mensaje de error y detiene la ejecución del programa usando os.Exit(1)
//
// log.Panic
// Muestra mensaje de error y ejecuta un panic
func LS19PaquetesInternos() {
	paqueteTime()
	paqueteOs()
	paqueteLog()
}

func paqueteTime() {
	fmt.Println(tools.MakeTitle("PACKAGE: time"))

	// objeto Time
	now := time.Now()
	fmt.Printf("\tnow: %v\n", now)

	// objeto Time
	fecha := time.Date(1990, time.March, 10, 0, 0, 0, 0, time.Local)
	fmt.Printf("\tDate: %v\n", fecha)
	fecha = fecha.Add(time.Hour * 128000)
	fmt.Printf("\tAgregando 128,000 horas: %v\n", fecha)
	showTimeProperties(fecha)
	showTimeComparing(now, fecha)
	fmt.Println()
}

func paqueteOs() {
	fmt.Println(tools.MakeTitle("PACKAGE: os"))

	// obtenemos directorio actual
	filePath, _ := os.Getwd()
	filePath = filepath.Join(filePath, "lessons", "someFile.txt")

	// abrimos archivo
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("\t%v\n", err)
	}

	defer func() {
		file.Close()
	}()

	// leemos contenido
	data := make([]byte, 1024)
	bytesReaded, err := file.Read(data)
	fmt.Println("\tArchivo:", file.Name())
	if err != nil {
		fmt.Println("\tError:", err)
		os.Exit(1)
	}
	fmt.Printf("\tread %d bytes: %q\n", bytesReaded, data[:bytesReaded])

	// leemos archivo
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("\tError:", err)
		os.Exit(1)
	}

	fmt.Println("\n\tFile content: ")
	fmt.Println(string(contentBytes))

	// mostramos variables de entorno
	showEnviromentVariables()

	// mostramos nombre de usuario
	username, ok := os.LookupEnv("USERNAME")
	if ok {
		fmt.Printf("\tUsername: %s\n", username)
	}

	// mostramos home path
	homePath, ok := os.LookupEnv("HOMEPATH")
	if ok {
		fmt.Printf("\tHomepath: %s\n", homePath)
	}

	// uso de Setenv y Expandenv
	os.Setenv("DB_USERNAME", "asdasdasd")
	os.Setenv("DB_PASSWORD", "asdasdasd")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "ASDASDASDA")

	// usamos ${var} o $var
	strConnection := os.ExpandEnv("mongodb://${DB_USERNAME}@${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Printf("\tstrConnection: %q\n", strConnection)

	// end lesson
	fmt.Println()
}

func paqueteLog() {
	fmt.Println(tools.MakeTitle("PACKAAGE: log"))

	// imprime mensaje
	log.Println("Hello World in Go, log version")

	// genera aleatorio y en base a este ejecutamos fatal
	randInteger := tools.RandInt(1, 2)
	fmt.Println("randInteger:", randInteger)
	// if randInteger%2 == 0 {
	// 	log.Fatalln("Error: se generó entero incorrecto")
	// }

	// ejecutamos panic
	// log.Panic("error: se genera error")

	// creamos archivo donde se almacenarán los logs
	file, err := os.Create(fmt.Sprintf("%d.log", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	for i := 1; i <= 10; i++ {
		logger.Printf("%q\n", fmt.Sprintf("test %d", i))
	}
}

// Muestra algunas propiedades del objeto Time por medio de invocaciónes
// de sus respectivos métodos.
func showTimeProperties(t time.Time) {
	fmt.Printf("\t%v\n", t)
	fmt.Printf("\tHour    : %v\n", t.Hour())
	fmt.Printf("\tMinute  : %v\n", t.Minute())
	fmt.Printf("\tSecond  : %v\n", t.Second())
	fmt.Printf("\tDay     : %v\n", t.Day())
	fmt.Printf("\tYear    : %v\n", t.Year())
	fmt.Printf("\tMonth   : %v\n", t.Month())
	fmt.Printf("\tWeekday : %v\n", t.Weekday())
	fmt.Printf("\tLocation: %v\n", t.Location())
	fmt.Printf("\tWeekday : %v\n", t.Weekday())
}

// Muetra los resultados de comparar dos fechas Time.
func showTimeComparing(t1, t2 time.Time) {
	fmt.Printf("\tt1.Before(t2) : %v\n", t1.Before(t2))
	fmt.Printf("\tt1.After(t2)  : %v\n", t1.After(t2))
	fmt.Printf("\tt1.Equal(t2)  : %v\n", t1.Equal(t2))
}

// Muestra variables de entorno.
func showEnviromentVariables() {
	fmt.Println("\tMostramos variables de entorno:")
	envVars := os.Environ()
	for _, v := range envVars {
		fmt.Printf("\t- %v\n", v)
	}
}
