# 🔐 EncryptGo - Herramienta de Encriptación de Archivos

Una herramienta robusta escrita en Go para encriptar y desencriptar archivos de forma segura con AES-256 y procesamiento paralelo.

## ✨ Características

- **🔒 Encriptación AES-256**: Cifrado de alta seguridad con CBC mode
- **⚡ Procesamiento Paralelo**: Procesa múltiples archivos simultáneamente (4 goroutines)
- **🛡️ Claves Fuertes**: Derivación de claves con PBKDF2 (100,000 iteraciones)
- **📁 Selección de Tipos**: Encripta solo los tipos de archivo que necesitas
- **🚫 Protección del Programa**: Excluye automáticamente archivos críticos
- **🎨 Terminal Colorida**: Interfaz amigable con colores ANSI
- **⚙️ Generación Automática de Claves**: O ingresa la tuya propia

## 🚀 Instalación

### Requisitos
- Go 1.21 o superior
- Linux/macOS/Windows

### Compilar

```bash
go mod tidy
go build -o encryptgo main.go
```

## 📖 Uso

### Ejecutar el programa

```bash
./encryptgo
```

Te aparecerá un menú interactivo:

```
  ██████╗ ██╗   ██╗     █████╗ ██████╗  ██████╗ ███████╗███╗  ██╗████████╗███████╗
  ██╔══██╗╚██╗ ██╔╝    ██╔══██╗██╔══██╗██╔════╝ ██╔════╝████╗ ██║╚══██╔══╝╚════██║
  ██████╔╝ ╚████╔╝     ███████║██████╔╝██║  ███╗█████╗  ██╔██╗██║   ██║       ██╔╝
  ██╔══██╗  ╚██╔╝      ██╔══██║██╔══██╗██║   ██║██╔══╝  ██║╚████║   ██║      ██╔╝ 
  ██████╔╝   ██║       ██║  ██║██║  ██║╚██████╔╝███████╗██║ ╚███║   ██║      ██║  
  ╚═════╝    ╚═╝       ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝  ╚══╝  ╚═╝      ╚═╝  
```

### Opción 1: Encriptar Archivos

```
[?] ¿Qué deseas hacer?
   1) Encriptar archivos
   2) Desencriptar archivos
➜ 1
```

Luego elige qué tipo de archivos encriptar:

```
[?] ¿Qué tipo de archivos quieres encriptar?
   1) Imágenes (.jpg, .png, .gif, etc)
   2) PDFs
   3) TXTs y documentos
   4) Audios (.mp3, .wav, etc)
   5) Videos (.mp4, .avi, etc)
   6) TODO (todos los archivos)
➜ 1
```

Especifica la carpeta:

```
[?] Ruta de la carpeta a procesar (ej: ./documentos o . para actual):
➜ ./fotos
```

Ingresa la clave (o déjala en blanco para generar una automática):

```
[?] Ingresa la clave de encriptación (deja en blanco para usar una por defecto):
➜ miClaveSegura123!
```

### Opción 2: Desencriptar Archivos

```
[?] ¿Qué deseas hacer?
   1) Encriptar archivos
   2) Desencriptar archivos
➜ 2
```

Especifica la carpeta con archivos `.encrypt`:

```
[?] Ruta de la carpeta a procesar (ej: ./documentos o . para actual):
➜ ./fotos
```

Ingresa la misma clave que usaste para encriptar:

```
[?] Ingresa la clave de desencriptación:
➜ miClaveSegura123!
```

## 📝 Tipos de Archivo Soportados

| Tipo | Extensiones |
|------|-------------|
| **Imágenes** | .jpg, .jpeg, .png, .gif, .bmp, .tiff, .webp |
| **PDFs** | .pdf |
| **Documentos** | .txt, .md, .log, .csv, .json, .xml, .yaml |
| **Audios** | .mp3, .wav, .flac, .aac, .ogg |
| **Videos** | .mp4, .avi, .mkv, .mov, .wmv, .flv |

## 🛡️ Medidas de Seguridad

### Encriptación
- **Algoritmo**: AES-256 (CBC mode)
- **Derivación de clave**: PBKDF2 con 100,000 iteraciones
- **IV Aleatorio**: Generado para cada archivo

### Protección del Programa
El programa **automáticamente excluye** estos archivos:
- Archivos del sistema: `main.go`, `main`, `go.mod`, `go.sum`
- Binarios: `.exe`, `.o`, `.a`, `.so`, `.dll`, `.dylib`
- Directorios: `.git`, `.github`

### Velocidad
Procesa hasta **4 archivos simultáneamente** usando goroutines, reduciendo significativamente el tiempo de procesamiento.

## ⚠️ Información Importante

### Cuidado con la Clave
- **🔑 Guarda tu clave en un lugar seguro**
- Sin la clave, **NO podrás recuperar tus archivos**
- Si olvidas la clave, los archivos estarán perdidos para siempre

### Archivos Encriptados
- Los archivos encriptados tienen la extensión `.encrypt`
- El archivo original se elimina automáticamente después de encriptarse
- Para desencriptar, necesitas una carpeta con archivos `.encrypt`

## 💻 Ejemplo de Uso Completo

```bash
# Encriptar todas las imágenes en la carpeta actual
$ ./encryptgo
[?] ¿Qué deseas hacer?
➜ 1                          # Encriptar
[?] ¿Qué tipo de archivos quieres encriptar?
➜ 1                          # Imágenes
[?] Ruta de la carpeta a procesar:
➜ .                          # Carpeta actual
[?] Ingresa la clave de encriptación:
➜                            # Dejar en blanco para generar automática

[✓] Encontrados 15 archivos
[🚀] Comenzando encriptación
  [1/15] Encriptando: photo1.jpg
  [2/15] Encriptando: photo2.jpg
  [3/15] Encriptando: photo3.jpg
  ...
[📊] RESUMEN DE ENCRIPTACIÓN
  ✓ Exitosos: 15
  ✗ Errores: 0
  🔑 Clave: VmF2aUlnM3dVRy94SjJRSDRjVE96bkxoZUtHTkVXUDFTOUhXVVJIVE9GeUk=
```

## 🔧 Estructura del Código

```
main.go
├── main()                          # Menú principal
├── manejarEncriptacion()           # Lógica de encriptación
├── manejarDesencriptacion()        # Lógica de desencriptación
├── procesarArchivosEncriptacion()  # Goroutines para encriptación
├── procesarArchivosDesencriptacion() # Goroutines para desencriptación
├── EncriptarArchivo()              # Cifrado de un archivo
├── DesencriptarArchivo()           # Descifrado de un archivo
├── derivarClave()                  # PBKDF2 para claves fuertes
├── buscarArchivos()                # Búsqueda con exclusiones
└── isExcludedFile()                # Validación de archivos
```

## 🐛 Solución de Problemas

**P: ¿Qué pasa si la encriptación falla a mitad de camino?**
R: El archivo se procesará con goroutines, intentando continuar con los demás. Los errores se reportan en el resumen final.

**P: ¿Puedo encriptar archivos muy grandes?**
R: Sí, el programa puede procesar archivos de cualquier tamaño.

**P: ¿Es seguro usar la clave generada automáticamente?**
R: Sí, se genera usando `crypto/rand` que es criptográficamente seguro.

**P: ¿Qué pasa si pierdo la clave?**
R: Lamentablemente, los archivos no podrán ser recuperados. La encriptación AES-256 no tiene "puerta trasera".

## 📝 Licencia

Código de argent7 - Encrypt Go Tool


**¡Gracias por usar Encrypt Go!** 🙌