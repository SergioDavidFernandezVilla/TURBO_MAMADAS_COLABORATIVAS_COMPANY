// HOOKS
import { useFetch } from "../../../hooks/useHookFetchData";

// KEY_API
import {
  BASE_API_ONLY_IMAGES_BYID_PRODUCT,
} from "../../../utils/API/key_API";

// TYPES
import type { Imagen } from "../../types/imageComponent/ImagenComponentType";

// COMPONENTS
import { ImageComponent } from "../imageComponent/ImageComponent";

type Props = {
  productId: string;
  tipo: string;
};

export const ListImagesComponent = ({ productId, tipo }: Props) => {
  const {
    data: imagenes,
    loading,
    error,
  } = useFetch<Imagen[]>(
    `${BASE_API_ONLY_IMAGES_BYID_PRODUCT}/${tipo}/${productId}`
  );

  if (loading) return <p>Cargando imágenes...</p>;
  if (error) return <p>Error cargando imágenes: {error}</p>;
  if (!imagenes || imagenes.length === 0) return <p>Sin imágenes</p>;

  console.log("imagenes", imagenes.map((img) => img.url))

  return (
    <>
      {imagenes.map((img) => (
        <figure key={img._id}>
          <ImageComponent key={img._id} url={img.url} alt={img.alt} />
        </figure>
      ))}
    </>
  );
};
