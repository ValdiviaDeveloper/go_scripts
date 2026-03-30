# 🔓 Port Scanner - Herramienta de Escaneo de Puertos

Una herramienta rápida y eficiente escrita en Go para escanear puertos abiertos en hosts remotos. Identifica servicios comunes y proporciona resultados detallados con soporte para múltiples plataformas.

## ✨ Características

- 🚀 **Escaneo rápido** - Utiliza goroutines para escanear hasta 100 puertos simultáneamente
- 🎨 **Interfaz colorida** - Resultados claros y fáciles de leer
- 🔧 **Portátil** - Descarga el ejecutable, sin necesidad de instalar Go
- 🪟 **Multiplataforma** - Disponible para Windows, macOS y Linux
- 🏷️ **Identificación de servicios** - Detecta automáticamente servicios comunes (HTTP, SSH, MySQL, etc.)
- ⚙️ **Flexible** - Escanea rangos personalizados de puertos

## 📥 Instalación

### Opción 1: Descargar ejecutable (Recomendado)

Ve a [Releases](https://github.com/ValdiviaDeveloper/go_scripts/releases) y descarga el ejecutable para tu sistema operativo:

- **Windows**: `portscanner-windows-amd64.exe`
- **macOS** (Intel): `portscanner-darwin-amd64`
- **macOS** (Apple Silicon): `portscanner-darwin-arm64`
- **Linux** (x86_64): `portscanner-linux-amd64`
- **Linux** (ARM): `portscanner-linux-arm64`

#### En Windows:
1. Descarga el archivo `.exe`
2. Haz doble clic para ejecutar
3. ¡Listo!

#### En macOS/Linux:
```bash
# Descarga el archivo
# Dale permisos de ejecución
chmod +x portscanner-linux-amd64

# Ejecuta
./portscanner-linux-amd64
```

### Opción 2: Compilar desde código

Si tienes Go instalado:

```bash
cd portscanner
go build -o portscanner main.go
./portscanner
```

## 🚀 Uso

### Modo interactivo (Recomendado)

Simplemente ejecuta el programa sin argumentos:

```bash
./portscanner
```

Luego ingresa:
- **Host**: IP o dominio a escanear (ej: `192.168.1.1` o `google.com`)
- **Puerto inicial**: Número inicial del rango (por defecto: 1)
- **Puerto final**: Número final del rango (por defecto: 1024)

### Modo línea de comandos

Pasa el host como argumento:

```bash
./portscanner 192.168.1.1
```

Luego ingresa los puertos como se pide.

## 📊 Ejemplos

### Escanear primeros 1000 puertos de localhost

```bash
./portscanner 127.0.0.1
# Puerto inicial: 1
# Puerto final: 1024
```

### Escanear puertos SSH, HTTP, HTTPS

```bash
./portscanner google.com
# Puerto inicial: 22
# Puerto final: 443
```

### Escanear máquina en red local

```bash
./portscanner 192.168.0.100
# Puerto inicial: 1
# Puerto final: 65535  (todos los puertos - puede tardar)
```

## 🎯 Puertos monitoreados

El scanner identifica automáticamente estos servicios comunes:

| Puerto | Servicio | Descripción |
|--------|----------|-------------|
| 21 | FTP | Transferencia de archivos |
| 22 | SSH | Conexión remota segura |
| 25 | SMTP | Envío de correos |
| 53 | DNS | Resolución de dominios |
| 80 | HTTP | Páginas web |
| 110 | POP3 | Recibir correos |
| 143 | IMAP | Recibir correos |
| 443 | HTTPS | Páginas web seguras |
| 445 | SMB | Archivos compartidos Windows |
| 3306 | MySQL | Base de datos |
| 3389 | RDP | Escritorio remoto Windows |
| 5432 | PostgreSQL | Base de datos |
| 6379 | Redis | Caché/base de datos |
| 8080 | HTTP-Alt | Servidor web alternativo |
| 9200 | Elasticsearch | Motor de búsqueda |
| 27017 | MongoDB | Base de datos NoSQL |

## ⚙️ Configuración del escaneo

- **Trabajadores simultáneos**: 100 goroutines para escaneo paralelo
- **Timeout por puerto**: 1 segundo
- **Rango de puertos**: 1 - 65535

## 💡 Consejos

- Para escaneos rápidos, deja los valores por defecto (1-1024)
- Para escaneos completos, usa 1-65535 (puede tardar varios minutos)
- Escanea solo hosts que tengas permiso de probar
- Ciertos hosts pueden bloquear escaneos frecuentes

## ⚠️ Disclaimer Legal

Esta herramienta está diseñada **exclusivamente para propósitos educativos, diagnósticos y de administración autorizada de sistemas**. Úsala responsablemente:

**Obligaciones legales:**
- ✅ **SOLO** escanea sistemas que administres, poseas o tengas autorización expresa por escrito del propietario
- ✅ Cumple con todas las leyes aplicables en tu país (incluyendo CFAA, GDPR, y leyes equivalentes)
- ✅ No intentes acceder a sistemas sin autorización
- ✅ No uses para propósitos maliciosos, ilícitos o no autorizados
- ✅ Respeta la privacidad y seguridad ajena

**Responsabilidad:**
- El autor no se responsabiliza por daños, pérdidas o consecuencias derivadas del uso de esta herramienta
- Usas esta herramienta bajo tu propio riesgo y responsabilidad completa
- Podrías enfrentar consecuencias legales graves por uso no autorizado

**Casos de uso válidos:**
- Administración de tu propia infraestructura
- Testing en ambientes de desarrollo/staging autorizados
- Auditorías de seguridad con consentimiento escrito
- Fines educativos y de estudio

## 🛠️ Requisitos técnicos

- **Memoria**: ~10MB
- **Espacio**: ~5MB
- **Sin dependencias externas**: Incluye todo lo necesario

## 📝 Licencia

Código de argent7 - Puerto Scanner Tool

## 🤝 Soporte

Si encuentras problemas:
1. Verifica que el host sea accesible
2. Intenta con un rango de puertos más pequeño
3. Comprueba tu conexión a internet
4. Crea un issue en GitHub

---

**¡Gracias por usar Port Scanner!** 🙌
