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

**Descargar:** [Ir a Releases](https://github.com/ValdiviaDeveloper/go_scripts/releases)

## 🚀 Empezando rápido

### Port Scanner

```bash
# Opción 1: Descargar ejecutable (sin Go requerido)
# Ve a https://github.com/ValdiviaDeveloper/go_scripts/releases

# Opción 2: Compilar desde código
cd portscanner
go build -o portscanner main.go
./portscanner
```

Para más detalles, consulta el [README de Port Scanner](./portscanner/README.md)

## 📋 Requisitos

- Go 1.25.4 o superior (solo si compilas desde código)
- Si descargas ejecutables, no necesitas nada más

## 🔄 CI/CD

Este proyecto usa **GitHub Actions** para compilar automáticamente los ejecutables:
- ✅ Compila para múltiples plataformas en cada release
- ✅ Binarios disponibles para descargar sin instalar Go
- ✅ Releases automáticos e independientes por herramienta

## 📝 Estructura del proyecto

```
go_scripts/
├── portscanner/          # Herramienta de escaneo de puertos
│   ├── main.go
│   ├── go.mod
│   └── README.md
├── .github/workflows/    # GitHub Actions pipelines
│   └── build-release.yml
├── LICENSE               # Licencia MIT
├── CODE_OF_CONDUCT.md    # Código de conducta
├── CONTRIBUTING.md       # Guía de contribución
├── RELEASE_GUIDE.md      # Cómo crear releases
├── ADD_TOOL.md          # Cómo agregar nuevas herramientas
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