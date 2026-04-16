# 🔧 Go Scripts

Colección de herramientas útiles implementadas en Go. Cada herramienta está diseñada para ser rápida, confiable y fácil de usar.

## 📦 Herramientas disponibles

### [Port Scanner](./portscanner/README.md) 🔓

Escanea puertos abiertos en hosts remotos e identifica servicios comunes. 

**Características:**
- ⚡ Escaneo rápido con goroutines paralelas
- 🎨 Interfaz colorida
- 🪟 Multiplataforma (Windows, macOS, Linux)
- 🏷️ Identificación automática de servicios

---

### [EncryptGo](./encryptgo/README.md) 🔐

Herramienta robusta para encriptar y desencriptar archivos con AES-256.

**Características:**
- 🔒 Encriptación AES-256 segura
- ⚡ Procesamiento paralelo (4 goroutines simultáneas)
- 🛡️ Claves fuertes con PBKDF2 (100k iteraciones)
- 📁 Selección flexible de tipos de archivo
- 🚫 Protección automática de archivos del programa
- 🎨 Interfaz terminal colorida

---

**📥 Descargar:** [Ir a Releases](https://github.com/ValdiviaDeveloper/go_scripts/releases)

## 🚀 Empezando rápido

### Descargar ejecutables (recomendado)

```bash
# Descargar desde releases
# https://github.com/ValdiviaDeveloper/go_scripts/releases

# Ejecutar directamente (sin necesidad de Go)
./portscanner
./encryptgo
```

### O compilar desde código

```bash
# Port Scanner
cd portscanner
go build -o portscanner main.go
./portscanner

# EncryptGo
cd ../encryptgo
go build -o encryptgo main.go
./encryptgo
```

Para más detalles, consulta los READMEs individuales:
- [Port Scanner](./portscanner/README.md)
- [EncryptGo](./encryptgo/README.md)

## 📋 Requisitos

- Go 1.25.4 o superior (solo si compilas desde código)
- Si descargas ejecutables, no necesitas nada más

## 🔄 CI/CD - Releases Automáticos

Este proyecto usa **GitHub Actions** para compilar automáticamente los ejecutables:
- ✅ **Compila para múltiples plataformas:** Windows, macOS (Intel/Apple Silicon), Linux
- ✅ **Releases independientes** por herramienta (etiquetas: `portscanner/v*` o `encryptgo/v*`)
- ✅ **Binarios listos** disponibles para descargar sin necesidad de instalar Go
- ✅ **Automatizado:** Solo crea un tag y GitHub Actions compilará para todas las plataformas

**Para crear un release:**

```bash
# Ejemplo para encryptgo
git tag encryptgo/v1.0.0
git push origin encryptgo/v1.0.0

# Ejemplo para portscanner
git tag portscanner/v2.0.0
git push origin portscanner/v2.0.0
```

Ver [RELEASE_GUIDE.md](./RELEASE_GUIDE.md) para más detalles.

## 📝 Estructura del proyecto

```
go_scripts/
├── portscanner/          # Herramienta de escaneo de puertos
│   ├── main.go
│   ├── go.mod
│   └── README.md
├── encryptgo/            # Herramienta de encriptacion de archivos
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── .github/workflows/    # GitHub Actions pipelines
│   └── build-release.yml # Compilacion automatica para multiples plataformas
├── LICENSE               # Licencia MIT
├── CODE_OF_CONDUCT.md    # Codigo de conducta
├── CONTRIBUTING.md       # Guia de contribucion
├── RELEASE_GUIDE.md      # Como crear releases
├── ADD_TOOL.md          # Como agregar nuevas herramientas
└── README.md            # Este archivo
```

## 📚 Guías

- **[Crear Releases](./RELEASE_GUIDE.md)** - Cómo publicar nuevas versiones
- **[Agregar Herramientas](./ADD_TOOL.md)** - Cómo añadir nuevas herramientas
- **[Contribuir](./CONTRIBUTING.md)** - Guía para colaboradores
- **[Código de Conducta](./CODE_OF_CONDUCT.md)** - Normas de uso responsable

## 📄 Licencia

Este proyecto está licenciado bajo la **MIT License** - ver [LICENSE](./LICENSE) para detalles completos.

## 🤝 Contribuir

¿Quieres mejorar este proyecto? Lee [CONTRIBUTING.md](./CONTRIBUTING.md) para aprender cómo.

## 📋 Código de Conducta

Este proyecto sigue un [Código de Conducta](./CODE_OF_CONDUCT.md) para mantener un espacio seguro y respetuoso.

---

**¿Necesitas ayuda?** Consulta el README de cada herramienta para instrucciones detalladas.