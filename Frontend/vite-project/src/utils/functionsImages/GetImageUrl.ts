// KEY_API
import { API_IMAGES_BASE } from "../API/key_API";

/**
 * Construye la URL completa de una imagen desde el nombre o path.
 * @param path Ruta o nombre de la imagen (ej: "123.png")
 * @returns URL completa de la imagen.
 */
export const getImageUrl = (path: string): string => {
  const cleanPath = path.startsWith("/") ? path.slice(1) : path;
  return `${API_IMAGES_BASE}${cleanPath}`;
};
