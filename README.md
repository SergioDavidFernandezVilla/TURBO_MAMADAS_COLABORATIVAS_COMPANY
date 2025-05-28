# 🛒 PROYECTO TIENDA ONLINE COLABORATIVA

Bienvenido al proyecto **Tienda Online Colaborativa**, una plataforma moderna construida con un stack completo que integra un frontend en React + Vite, múltiples microservicios en Go (API REST), manejo de imágenes, todo orquestado con Docker y Docker Compose.

---

## 📦 Estructura del Proyecto

TURBO_MAMADAS_COLABORATIVAS_COMPANY/
├── Backend/
│ ├── server-go-welcome/ # Servicio de bienvenida (punto de entrada)
│ ├── server-go-products/ # Microservicio de productos (CRUD)
│ └── server-go-images/ # Microservicio de imágenes
├── Frontend/
│ └── vite-project/ # Aplicación frontend construida con Vite + React
├── images_products/ # Carpeta donde se almacenan las imágenes subidas
├── docker-compose.yml # Orquestador de servicios
└── README.md # Este archivo 📄

---

## 🚀 Tecnologías Usadas

- **Frontend**: React + TypeScript + Vite
- **Backend**: Go (con framework Gin)
- **API de Imágenes**: Subida, consulta y asociación por UUID
- **Almacenamiento**: Archivos JSON y sistema de carpetas para imágenes
- **Docker**: Contenedores individuales para cada servicio
- **Docker Compose**: Orquestación y redes internas

---

## 🧱 Estructura de Imágenes

images_products/
└── uploads/
└── tipo/ # Ej: productos, perfil, etc.
└── <uuid>/ # UUID por entidad (producto, usuario, etc.)
└── <uuid>.ext # Archivo de imagen con su extensión original


---

## ⚙️ Cómo levantar el proyecto

### 1. Clonar el repositorio

```bash
git clone https://github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY.git
cd TURBO_MAMADAS_COLABORATIVAS_COMPANY
```


### 2. Levantar con Docker
```
docker-compose up --build
```

#### Esto levanta los siguientes servicios:

Servicio	URL de acceso
Welcome	http://localhost:8080
Products	http://localhost:8081/api/productos
Images	http://localhost:8082/api/v1/upload/productos
Frontend	http://localhost:3000


📡 Endpoints Principales
Productos (/api/productos)
GET → Lista todos los productos

POST → Crea un producto

PUT → Actualiza un producto

DELETE → Elimina un producto

Imágenes (/api/v1/upload/productos)
GET → Lista de imágenes subidas

POST → Sube una imagen asociada a un producto


🧪 Estado del Proyecto
 Servicios backend en Go (welcome, productos, imágenes)

 Frontend con consumo dinámico de APIs

 Subida de imágenes por tipo y UUID

 Vistas y componentes React reutilizables

👥 Colaboradores
