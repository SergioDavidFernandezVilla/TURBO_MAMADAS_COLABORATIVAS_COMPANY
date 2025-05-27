/**
 * Archivo de configuración para las URLs base de las APIs usadas en la aplicación.
 * Aquí se definen las constantes globales para interactuar con los distintos
 * endpoints de productos e imágenes.
 */

// Detecta entorno local
const isLocalhost =
  typeof window !== "undefined" && window.location.hostname === "localhost";

// URL base del backend de productos
const BASE_PRODUCTOS = isLocalhost
  ? "http://localhost:8081"
  : "https://8081-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev";

// URL base del backend de imágenes
const BASE_IMAGES = isLocalhost
  ? "http://localhost:8082"
  : "https://8082-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev";

// Endpoints exportados
export const API_PRODUCTOS = `${BASE_PRODUCTOS}/api/productos`;

export const API_IMAGES = `${BASE_IMAGES}/api/v1/upload/productos`;
