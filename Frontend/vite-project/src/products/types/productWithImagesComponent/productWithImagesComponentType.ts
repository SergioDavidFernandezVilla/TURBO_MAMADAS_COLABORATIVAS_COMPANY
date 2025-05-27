import type { Imagen } from "../../../images/types/imageComponent/ImagenComponentType";
import type { Producto } from "../productComponent/ProductComponentType";

export type ProductsWithImages = Producto & {
    imagenes: Imagen[];
  };