import { useEffect, useState } from "react";
import { useFetch } from "../../hooks/useHookFetchData";
import type { Imagen } from "../../images/types/imageComponent/ImagenComponentType";
import type { Producto } from "../../products/types/productComponent/ProductComponentType";
import { API_PRODUCTOS, API_IMAGES } from "../../utils/API/key_API";

export const ProductGalleryCarousel = () => {
  const { data: productos, loading: loadingP } =
    useFetch<Producto[]>(API_PRODUCTOS);
  const [imagenesPorProducto, setImagenesPorProducto] = useState<
    Record<string, Imagen[]>
  >({});
  const [loadingI, setLoadingI] = useState(true);

  
  if (loadingP || loadingI) return <p>Cargando galería...</p>;



  return (
    <div>
      {productos?.map((producto) => (
        <div key={producto.ID}>
          <h3>{producto.Nombre}</h3>
          <p>
            {producto.Categoria} - ${producto.Precio}
          </p>
          {imagenesPorProducto[producto.ID]?.length > 0 ? (
            <div>
              {imagenesPorProducto[producto.ID].map((img) => (
                <img key={img._id} src={img.url} alt={img.alt} />
              ))}
            </div>
          ) : (
            <p>Sin imágenes</p>
          )}
        </div>
      ))}
    </div>
  );
};
