package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	colorReset  = "\033[0m"  // Vuelve al color normal
	colorRed    = "\033[31m" // Rojo  → errores
	colorGreen  = "\033[32m" // Verde → puerto abierto
	colorYellow = "\033[33m" // Amarillo → advertencias / info
	colorCyan   = "\033[36m" // Cyan  → títulos y encabezados
	colorBold   = "\033[1m"  // Negrita
	colorDim    = "\033[2m"  // Texto tenue (gris)
)

var servicios = map[int]string{
	21:    "FTP      (transferencia de archivos)",
	22:    "SSH      (conexión remota segura)",
	23:    "Telnet   (conexión remota antigua)",
	25:    "SMTP     (envío de correos)",
	53:    "DNS      (resolución de dominios)",
	80:    "HTTP     (páginas web)",
	110:   "POP3     (recibir correos)",
	143:   "IMAP     (recibir correos)",
	443:   "HTTPS    (páginas web seguras)",
	445:   "SMB      (archivos compartidos Windows)",
	3306:  "MySQL    (base de datos)",
	3389:  "RDP      (escritorio remoto Windows)",
	5432:  "PostgreSQL (base de datos)",
	6379:  "Redis    (caché / base de datos)",
	8080:  "HTTP-Alt (servidor web alternativo)",
	8443:  "HTTPS-Alt(servidor web seguro alt.)",
	9200:  "Elasticsearch (motor de búsqueda)",
	27017: "MongoDB (base de datos)",
}

type ResultadoPuerto struct {
	Puerto   int
	Abierto  bool
	Servicio string
}

func obtenerServicio(puerto int) string {
	if nombre, existe := servicios[puerto]; existe {
		return nombre
	}
	return "Desconocido"
}

func trabajador(host string, puertos <-chan int, results chan<- ResultadoPuerto, wg *sync.WaitGroup) {
	defer wg.Done()

	for puerto := range puertos {
		direccion := fmt.Sprintf("%s:%d", host, puerto)

		conn, err := net.DialTimeout("tcp", direccion, 1*time.Second)

		resultado := ResultadoPuerto{
			Puerto:   puerto,
			Abierto:  false,
			Servicio: obtenerServicio(puerto),
		}

		if err == nil {
			resultado.Abierto = true
			conn.Close()
		}

		results <- resultado
	}
}

func imprimirBanner() {
	fmt.Println()

	// Logo principal — argent7 en bloques
	logo := []string{
		`  ██████╗ ██╗   ██╗     █████╗ ██████╗  ██████╗ ███████╗███╗  ██╗████████╗███████╗`,
		`  ██╔══██╗╚██╗ ██╔╝    ██╔══██╗██╔══██╗██╔════╝ ██╔════╝████╗ ██║╚══██╔══╝╚════██║`,
		`  ██████╔╝ ╚████╔╝     ███████║██████╔╝██║  ███╗█████╗  ██╔██╗██║   ██║       ██╔╝`,
		`  ██╔══██╗  ╚██╔╝      ██╔══██║██╔══██╗██║   ██║██╔══╝  ██║╚████║   ██║      ██╔╝ `,
		`  ██████╔╝   ██║       ██║  ██║██║  ██║╚██████╔╝███████╗██║ ╚███║   ██║      ██║  `,
		`  ╚═════╝    ╚═╝       ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚══╝  ╚═╝      ╚═╝  `,
	}

	for _, linea := range logo {
		fmt.Printf("%s%s%s\n", colorCyan, linea, colorReset)
	}

	// Línea separadora
	fmt.Printf("%s  %s%s\n", colorDim, repeatChar("─", 84), colorReset)

	// Pie del banner
	fmt.Printf("%s  [%s by argent7 %s]  port scanner · go tool%s\n",
		colorDim, colorYellow, colorDim, colorReset)

	fmt.Println()
}

func repeatChar(char string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}

func pedirInput(mensaje string) string {
	fmt.Printf("%s%s-> %s%s", colorCyan, colorBold, colorReset, mensaje)
	var entrada string
	fmt.Scanln(&entrada)
	return strings.TrimSpace(entrada)
}

func main() {
	imprimirBanner()

	var host string
	if len(os.Args) > 1 {
		host = os.Args[1]
		fmt.Printf("%s Host tomado de argumentos: %s%s\n", colorGreen, host, colorReset)
	} else {
		host = pedirInput("Ingresa IP o dominio a escanera (ej: 192.168.1.1 o google.com): ")
	}

	if host == "" {
		fmt.Printf("%s Error: debes ingresar un host válido%s\n", colorRed, colorReset)
		os.Exit(1)
	}

	inicioStr := pedirInput("Puerto inicial (Enter para usar 1): ")
	finStr := pedirInput("Puerto final (Enter para usar 1024): ")

	inicio := 1
	fin := 1024

	if inicioStr != "" {
		if n, err := strconv.Atoi(inicioStr); err == nil && n > 0 {
			inicio = n
		}
	}

	if finStr != "" {
		if n, err := strconv.Atoi(finStr); err == nil && n <= 65535 {
			fin = n
		}
	}

	if inicio > fin {
		inicio, fin = fin, inicio
	}

	totalPuertos := fin - inicio + 1

	fmt.Println()
	fmt.Printf("%s%s  ┌─ Configuración del escaneo ──────────────┐%s\n", colorCyan, colorBold, colorReset)
	fmt.Printf("%s  │%s  Host       : %s%-35s%s%s  │%s\n", colorCyan, colorReset, colorYellow, host, colorReset, colorCyan, colorReset)
	rangoStr := fmt.Sprintf("%d - %d", inicio, fin)
	padding := 27 - len(rangoStr)
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("%s  │%s  Rango      : %s%s%s%-*s%s  │%s\n", colorCyan, colorReset, colorYellow, rangoStr, colorReset, padding, "", colorCyan, colorReset)
	fmt.Printf("%s  │%s  Total      : %s%d puertos%-27s%s  │%s\n", colorCyan, colorReset, colorYellow, totalPuertos, "", colorCyan, colorReset)
	fmt.Printf("%s  └──────────────────────────────────────────┘%s\n\n", colorCyan, colorReset)

	const numTrabajadores = 100

	canalPuertos := make(chan int, 100)
	canalResultados := make(chan ResultadoPuerto, totalPuertos)

	var wg sync.WaitGroup

	fmt.Printf("%s Iniciando %d trabajadores...%s\n\n", colorYellow, numTrabajadores, colorReset)

	for i := 0; i < numTrabajadores; i++ {
		wg.Add(1)
		go trabajador(host, canalPuertos, canalResultados, &wg)
	}

	inicio_tiempo := time.Now()

	go func() {
		for p := inicio; p <= fin; p++ {
			canalPuertos <- p
		}
		close(canalPuertos)
	}()

	go func() {
		wg.Wait()
		close(canalResultados)
	}()

	var puertosAbiertos []ResultadoPuerto
	escaneados := 0

	fmt.Printf("%s Escaneado", colorDim)

	for resultado := range canalResultados {
		escaneados++

		if escaneados%50 == 0 {
			fmt.Printf(".")
		}

		if resultado.Abierto {
			puertosAbiertos = append(puertosAbiertos, resultado)
		}
	}

	duracion := time.Since(inicio_tiempo)

	fmt.Printf("Listo!%s\n\n", colorReset)

	sort.Slice(puertosAbiertos, func(i, j int) bool {
		return puertosAbiertos[i].Puerto < puertosAbiertos[j].Puerto
	})

	fmt.Printf("%s%s  ┌─ Resultados ─────────────────────────────┐%s\n", colorCyan, colorBold, colorReset)
	fmt.Printf("%s  │%s  Host escaneado : %-26s%s%s  │%s\n", colorCyan, colorReset, host, colorReset, colorCyan, colorReset)
	fmt.Printf("%s  │%s  Puertos revisados: %-24d%s%s  │%s\n", colorCyan, colorReset, totalPuertos, colorReset, colorCyan, colorReset)
	fmt.Printf("%s  │%s  Tiempo total   : %-26s%s%s  │%s\n", colorCyan, colorReset, duracion.Round(time.Millisecond), colorReset, colorCyan, colorReset)
	fmt.Printf("%s  └──────────────────────────────────────────┘%s\n\n", colorCyan, colorReset)

	if len(puertosAbiertos) == 0 {
		fmt.Printf("%s  ✘ No se encontraron puertos abiertos en el rango %d-%d%s\n", colorRed, inicio, fin, colorReset)
	} else {
		fmt.Printf("%s%s  PUERTOS ABIERTOS (%d encontrados):%s\n\n", colorGreen, colorBold, len(puertosAbiertos), colorReset)

		// Encabezado de la tabla
		fmt.Printf("%s  %-8s  %-45s%s\n", colorBold, "PUERTO", "SERVICIO", colorReset)
		fmt.Printf("%s  %s%s\n", colorDim, strings.Repeat("─", 55), colorReset)

		// Una fila por cada puerto abierto
		for _, r := range puertosAbiertos {
			fmt.Printf("%s  %-8d%s  %s\n",
				colorGreen,
				r.Puerto,
				colorReset,
				r.Servicio,
			)
		}
	}

	fmt.Println()
	if runtime.GOOS == "windows" {
		fmt.Printf("%sPresiona Enter para cerrar...%s", colorDim, colorReset)
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
}
