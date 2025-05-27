# Usa una imagen base de Node.js
FROM node:20-alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos package.json y package-lock.json (o yarn.lock)
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia el resto de los archivos de la aplicación
COPY . .

# Construye la aplicación
RUN npm run build

# Usa una imagen base más ligera para el servidor de producción
FROM nginx:alpine

# Copia los archivos de construcción desde la etapa anterior
COPY --from=builder /app/dist /usr/share/nginx/html

# Expón el puerto 80 (puerto por defecto de Nginx)
EXPOSE 80

# Comando para iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]
