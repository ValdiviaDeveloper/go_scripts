package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/crypto/pbkdf2"
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

type Encryptor struct {
	key []byte
}

// Resultado de procesamiento para goroutines
type ProcesarResult struct {
	Archivo string
	Exitoso bool
	Error   error
}

var (
	imageExts = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp"}
	pdfExts   = []string{".pdf"}
	txtExts   = []string{".txt", ".md", ".log", ".csv", ".json", ".xml", ".yaml"}
	audioExts = []string{".mp3", ".wav", ".flac", ".aac", ".ogg"}
	videoExts = []string{".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv"}
	allExts   = []string{} // Se llenará con todas las extensiones

	// Archivos y extensiones del programa a excluir
	excludedFiles = map[string]bool{
		"main.go": true,
		"main":    true,
		"go.mod":  true,
		"go.sum":  true,
	}
	excludedExts = map[string]bool{
		".exe":   true,
		".o":     true,
		".a":     true,
		".so":    true,
		".dll":   true,
		".dylib": true,
	}
	excludedDirs = map[string]bool{
		".git":    true,
		".github": true,
	}
)

func init() {
	// Combinar todas las extensiones
	allExts = append(allExts, imageExts...)
	allExts = append(allExts, pdfExts...)
	allExts = append(allExts, txtExts...)
	allExts = append(allExts, audioExts...)
	allExts = append(allExts, videoExts...)
}

func main() {
	imprimirBanner()

	scanner := bufio.NewScanner(os.Stdin)

	// 1. Preguntar acción
	fmt.Printf("%s[?] ¿Qué deseas hacer?%s\n", colorCyan, colorReset)
	fmt.Printf("   %s1%s) Encriptar archivos\n", colorGreen, colorReset)
	fmt.Printf("   %s2%s) Desencriptar archivos\n", colorGreen, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)

	scanner.Scan()
	accion := strings.TrimSpace(scanner.Text())

	switch accion {
	case "1", "encriptar", "e", "E":
		manejarEncriptacion(scanner)
	case "2", "desencriptar", "d", "D":
		manejarDesencriptacion(scanner)
	default:
		fmt.Printf("%s[!] Opción no válida%s\n", colorRed, colorReset)
	}
}

func manejarDesencriptacion(scanner *bufio.Scanner) {
	fmt.Printf("\n%s[🔓] MODO DESENCRIPTACIÓN ACTIVADO%s\n", colorGreen, colorReset)

	fmt.Printf("\n%s[?] Ruta de la carpeta a procesar (ej: ./documentos o . para actual):%s\n", colorCyan, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)
	scanner.Scan()
	ruta := strings.TrimSpace(scanner.Text())
	if ruta == "" {
		ruta = "."
	}

	fmt.Printf("\n%s[?] Ingresa la clave de desencriptación:%s\n", colorCyan, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)
	scanner.Scan()
	clave := strings.TrimSpace(scanner.Text())

	if clave == "" {
		fmt.Printf("%s[✗] Error: Necesitas proporcionar una clave%s\n", colorRed, colorReset)
		return
	}

	encryptor, err := NewEncryptor(clave)
	if err != nil {
		fmt.Printf("%s[✗] Error: Clave inválida - %v%s\n", colorRed, err, colorReset)
		return
	}

	fmt.Printf("\n%s[🔍] Buscando archivos encriptados...%s\n", colorCyan, colorReset)
	archivosEncriptados := buscarArchivosEncriptados(ruta)

	if len(archivosEncriptados) == 0 {
		fmt.Printf("%s[!] No se encontraron archivos .encrypt en la ruta especificada%s\n", colorYellow, colorReset)
		return
	}

	fmt.Printf("%s[✓] Encontrados %d archivos encriptados%s\n", colorGreen, len(archivosEncriptados), colorReset)

	// Procesar archivos con goroutines
	exitos, errores := procesarArchivosDesencriptacion(archivosEncriptados, encryptor)

	fmt.Printf("\n%s════════════════════════════════════════%s\n", colorCyan, colorReset)
	fmt.Printf("%s[📊] RESUMEN DE DESENCRIPTACIÓN%s\n", colorBold, colorReset)
	fmt.Printf("  %s✓ Exitosos: %d%s\n", colorGreen, exitos, colorReset)
	fmt.Printf("  %s✗ Errores: %d%s\n", colorRed, errores, colorReset)
	fmt.Printf("%s════════════════════════════════════════%s\n", colorCyan, colorReset)

}

func (e *Encryptor) DesencriptarArchivo(ruta string) error {
	data, err := os.ReadFile(ruta)
	if err != nil {
		return err
	}

	if len(data) < 16 {
		return fmt.Errorf("archivo demasiado pequeño para ser válido")
	}

	iv := data[:16]
	ciphertext := data[16:]

	block, err := aes.NewCipher(e.key[:16])
	if err != nil {
		return err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return fmt.Errorf("Archivo corrupto")
	}

	plaintextPadded := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintextPadded, ciphertext)

	plaintext, err := pkcs7Unpad(plaintextPadded, aes.BlockSize)
	if err != nil {
		return err
	}

	outputPath := strings.TrimSuffix(ruta, ".encrypt")
	if err := os.WriteFile(outputPath, plaintext, 0644); err != nil {
		return err
	}

	if err := os.Remove(ruta); err != nil {
		return err
	}

	return nil
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("datos vacíos")
	}
	padding := int(data[len(data)-1])
	if padding > len(data) || padding > aes.BlockSize {
		return nil, fmt.Errorf("padding inválido")
	}
	return data[:len(data)-padding], nil
}

func buscarArchivosEncriptados(ruta string) []string {
	var archivos []string

	err := filepath.Walk(ruta, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".encrypt") {
			archivos = append(archivos, path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("%s[!] Error buscando archivos: %v%s\n", colorYellow, err, colorReset)
	}

	return archivos
}

func manejarEncriptacion(scanner *bufio.Scanner) {
	fmt.Printf("\n%s[🔐] MODO ENCRIPTACIÓN ACTIVADO%s\n", colorGreen, colorReset)

	// 2. Preguntar qué encriptar
	fmt.Printf("\n%s[?] ¿Qué tipo de archivos quieres encriptar?%s\n", colorCyan, colorReset)
	fmt.Printf("   %s1%s) Imágenes (.jpg, .png, .gif, etc)\n", colorGreen, colorReset)
	fmt.Printf("   %s2%s) PDFs\n", colorGreen, colorReset)
	fmt.Printf("   %s3%s) TXTs y documentos\n", colorGreen, colorReset)
	fmt.Printf("   %s4%s) Audios (.mp3, .wav, etc)\n", colorGreen, colorReset)
	fmt.Printf("   %s5%s) Videos (.mp4, .avi, etc)\n", colorGreen, colorReset)
	fmt.Printf("   %s6%s) TODO (todos los archivos)\n", colorGreen, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)

	scanner.Scan()
	tipo := strings.TrimSpace(scanner.Text())

	// seleccionar extensiones según la opción
	var extensiones []string
	switch tipo {
	case "1":
		// Encriptar imágenes
		extensiones = imageExts
		fmt.Printf("%s[✓] Seleccionado: IMÁGENES%s\n", colorGreen, colorReset)
	case "2":
		// Encriptar PDFs
		extensiones = pdfExts
		fmt.Printf("%s[✓] Seleccionado: PDFs%s\n", colorGreen, colorReset)
	case "3":
		// Encriptar TXTs y documentos
		extensiones = txtExts
		fmt.Printf("%s[✓] Seleccionado: TXTs y documentos%s\n", colorGreen, colorReset)
	case "4":
		// Encriptar audios
		extensiones = audioExts
		fmt.Printf("%s[✓] Seleccionado: Audios%s\n", colorGreen, colorReset)
	case "5":
		// Encriptar videos
		extensiones = videoExts
		fmt.Printf("%s[✓] Seleccionado: Videos%s\n", colorGreen, colorReset)
	case "6":
		// Encriptar todos los archivos
		extensiones = allExts
		fmt.Printf("%s[✓] Seleccionado: TODOS LOS ARCHIVOS%s\n", colorGreen, colorReset)
	default:
		extensiones = allExts
		fmt.Printf("%s[!] Opción no válida%s\n", colorRed, colorReset)
	}

	// Preguntar ruta
	fmt.Printf("\n%s[?] Ruta de la carpeta a procesar (ej: ./documentos o . para actual):%s\n", colorCyan, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)
	scanner.Scan()
	ruta := strings.TrimSpace(scanner.Text())
	if ruta == "" {
		ruta = "."
	}

	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		fmt.Printf("%s[!] La ruta '%s' no existe%s\n", colorRed, ruta, colorReset)
		return
	}

	fmt.Printf("\n%s[?] Ingresa la clave de encriptación (deja en blanco para usar una por defecto):%s\n", colorCyan, colorReset)
	fmt.Printf("%s➜ %s", colorYellow, colorReset)
	scanner.Scan()
	clave := strings.TrimSpace(scanner.Text())

	var claveBase64 string
	var err error
	if len(clave) == 0 {
		claveBase64, err = generarClaveAutomatica()
		if err != nil {
			fmt.Printf("%s[!] Error al generar clave automática: %s%s\n", colorRed, err, colorReset)
			return
		}
	} else {
		claveBase64, err = derivarClave(clave)
		if err != nil {
			fmt.Printf("%s[!] Error al procesar clave: %s%s\n", colorRed, err, colorReset)
			return
		}
	}

	fmt.Printf("\n%s[✓] Clave de encriptación: %s%s\n", colorGreen, claveBase64, colorReset)

	fmt.Printf("\n%s[🚀] Buscando archivos para encriptar...%s\n", colorCyan, colorReset)

	archivos := buscarArchivos(ruta, extensiones)

	if len(archivos) == 0 {
		fmt.Printf("%s[!] No se encontraron archivos para encriptar%s\n", colorYellow, colorReset)
		return
	}

	fmt.Printf("%s[✓] Encontrados %d archivos%s\n", colorGreen, len(archivos), colorReset)
	fmt.Printf("%s[⚠️] Comenzando encriptación (esto puede tomar un momento)...%s\n", colorYellow, colorReset)

	// Crear encryptor
	encryptor, err := NewEncryptor(claveBase64)
	if err != nil {
		fmt.Printf("%s[!] Error al crear encryptor: %s%s\n", colorRed, err, colorReset)
		return
	}

	// Procesar archivos con goroutines
	exitos, errores := procesarArchivosEncriptacion(archivos, encryptor)

	// Mostrar resumen
	fmt.Printf("\n%s════════════════════════════════════════%s\n", colorCyan, colorReset)
	fmt.Printf("%s[📊] RESUMEN DE ENCRIPTACIÓN%s\n", colorBold, colorReset)
	fmt.Printf("  %s✓ Exitosos: %d%s\n", colorGreen, exitos, colorReset)
	fmt.Printf("  %s✗ Errores: %d%s\n", colorRed, errores, colorReset)
	fmt.Printf("  %s🔑 Clave: %s%s\n", colorYellow, claveBase64, colorReset)
	fmt.Printf("%s════════════════════════════════════════%s\n", colorCyan, colorReset)

	if exitos > 0 {
		fmt.Printf("%s[⚠️] IMPORTANTE: Guarda la clave en un lugar seguro%s\n", colorYellow, colorReset)
		fmt.Printf("%s[!] Sin la clave NO podrás recuperar tus archivos%s\n", colorRed, colorReset)
	}
}

func (e *Encryptor) EncriptarArchivo(ruta string) error {
	data, err := os.ReadFile(ruta)
	if err != nil {
		return err
	}

	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	block, err := aes.NewCipher(e.key[:16])
	if err != nil {
		return err
	}

	paddedData := pkcs7Pad(data, aes.BlockSize)

	ciphertext := make([]byte, len(paddedData))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedData)

	encryptedData := append(iv, ciphertext...)
	outputPath := ruta + ".encrypt"

	if err := os.WriteFile(outputPath, encryptedData, 0644); err != nil {
		return err
	}

	if err := os.Remove(ruta); err != nil {
		return err
	}

	return nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

func NewEncryptor(claveBase64 string) (*Encryptor, error) {
	key, err := base64.URLEncoding.DecodeString(claveBase64)
	if err != nil {
		return nil, fmt.Errorf("clave inválida: %v", err)
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("la clave debe ser de 32 bytes (base64 de 44 caracteres)")
	}
	return &Encryptor{key: key}, nil
}

func buscarArchivos(ruta string, extensiones []string) []string {
	var archivos []string

	extMap := make(map[string]bool)
	for _, ext := range extensiones {
		extMap[strings.ToLower(ext)] = true
	}

	err := filepath.Walk(ruta, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Saltar directorios excluidos
		if info.IsDir() {
			if excludedDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		// Verificar si el archivo debe ser excluido
		if isExcludedFile(path) {
			return nil
		}

		// Verificar extensión
		ext := strings.ToLower(filepath.Ext(path))
		if extMap[ext] {
			archivos = append(archivos, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("%s[!] Error al buscar archivos: %s%s\n", colorRed, err, colorReset)
	}

	return archivos
}

func generarClaveAutomatica() (string, error) {
	key := make([]byte, 32) // AES-256
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func imprimirBanner() {
	fmt.Println()

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
	fmt.Printf("%s  [%s by argent7 %s]  file encrypter · go tool%s\n",
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

// Verifica si un archivo debe ser excluido del procesamiento
func isExcludedFile(path string) bool {
	basename := filepath.Base(path)

	// Chequear archivos específicos
	if lowerName := strings.ToLower(basename); excludedFiles[lowerName] {
		fmt.Printf("    %s[⊘] Saltando archivo del programa: %s%s\n", colorDim, basename, colorReset)
		return true
	}

	// Chequear extensiones excluidas
	ext := strings.ToLower(filepath.Ext(basename))
	if excludedExts[ext] {
		fmt.Printf("    %s[⊘] Saltando archivo binario: %s%s\n", colorDim, basename, colorReset)
		return true
	}

	return false
}

// Derivar clave fuerte usando PBKDF2
func derivarClave(password string) (string, error) {
	// Usar un salt consistente basado en el password (en producción usar random salt)
	salt := []byte("encryptgo-2024-salt-v1")

	// PBKDF2 con 100,000 iteraciones (estándar moderno)
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	return base64.URLEncoding.EncodeToString(key), nil
}

// Procesar archivos de encriptación con goroutines
func procesarArchivosEncriptacion(archivos []string, encryptor *Encryptor) (int, int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	exitos := 0
	errores := 0
	resultChan := make(chan ProcesarResult, len(archivos))

	// Limitar concurrencia a 4 goroutines
	semaphore := make(chan struct{}, 4)

	for i, archivo := range archivos {
		wg.Add(1)
		go func(idx int, path string) {
			defer wg.Done()
			semaphore <- struct{}{}        // Adquirir
			defer func() { <-semaphore }() // Liberar

			fmt.Printf("  [%d/%d] Encriptando: %s\n", idx+1, len(archivos), filepath.Base(path))
			err := encryptor.EncriptarArchivo(path)

			if err != nil {
				fmt.Printf("    %s[✗] Error: %v%s\n", colorRed, err, colorReset)
				resultChan <- ProcesarResult{Archivo: path, Exitoso: false, Error: err}
			} else {
				fmt.Printf("    %s[✓] Encriptado exitosamente%s\n", colorGreen, colorReset)
				resultChan <- ProcesarResult{Archivo: path, Exitoso: true}
			}
		}(i+1, archivo)
	}

	// Goroutine para recolectar resultados
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Procesar resultados
	for result := range resultChan {
		mu.Lock()
		if result.Exitoso {
			exitos++
		} else {
			errores++
		}
		mu.Unlock()
	}

	return exitos, errores
}

// Procesar archivos de desencriptación con goroutines
func procesarArchivosDesencriptacion(archivos []string, encryptor *Encryptor) (int, int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	exitos := 0
	errores := 0
	resultChan := make(chan ProcesarResult, len(archivos))

	// Limitar concurrencia a 4 goroutines
	semaphore := make(chan struct{}, 4)

	for i, archivo := range archivos {
		wg.Add(1)
		go func(idx int, path string) {
			defer wg.Done()
			semaphore <- struct{}{}        // Adquirir
			defer func() { <-semaphore }() // Liberar

			fmt.Printf("  [%d/%d] Desencriptando: %s\n", idx+1, len(archivos), filepath.Base(path))
			err := encryptor.DesencriptarArchivo(path)

			if err != nil {
				fmt.Printf("    %s[✗] Error: %v%s\n", colorRed, err, colorReset)
				resultChan <- ProcesarResult{Archivo: path, Exitoso: false, Error: err}
			} else {
				fmt.Printf("    %s[✓] Desencriptado exitosamente%s\n", colorGreen, colorReset)
				resultChan <- ProcesarResult{Archivo: path, Exitoso: true}
			}
		}(i+1, archivo)
	}

	// Goroutine para recolectar resultados
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Procesar resultados
	for result := range resultChan {
		mu.Lock()
		if result.Exitoso {
			exitos++
		} else {
			errores++
		}
		mu.Unlock()
	}

	return exitos, errores
}
