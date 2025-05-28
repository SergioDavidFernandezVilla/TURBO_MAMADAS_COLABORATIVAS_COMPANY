/**
 * Configuración de URLs base para APIs de productos e imágenes.
 * Cambia automáticamente entre entorno local y producción.
 */

const isLocalhost =
  typeof window !== "undefined" && window.location.hostname === "localhost";

// ======== BASES ======== //
const BASE_PRODUCTOS = isLocalhost
  ? "http://localhost:8081"
  : "https://8081-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev";

const BASE_IMAGES = isLocalhost
  ? "http://localhost:8082"
  : "https://8082-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev";

// ======== API PRODUCTOS ======== //
export const API_PRODUCTOS = `${BASE_PRODUCTOS}/api/productos`;

// ======== API IMÁGENES ======== //
export const API_IMAGES_UPLOAD = `${BASE_IMAGES}/api/v1/upload/`;
export const API_IMAGES_BY_PRODUCT = `${BASE_IMAGES}/api/v1/upload`;
export const API_IMAGES_BASE = `${BASE_IMAGES}/api/v1/`;

// ======== SOLO PARA PRODUCCIÓN (si necesitas apuntar fijo) ======== //
export const BASE_API_ONLY_PRODUCTOS =
  "https://8081-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev";
export const BASE_API_ONLY_IMAGES_BYID_PRODUCT =
  "https://8082-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev/api/v1/upload";
export const API_ONLY_IMAGES_SHOW =
  "https://8082-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev/api/v1/";
