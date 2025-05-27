// KEY_APIS
import { API_PRODUCTOS } from "../../../utils/API/key_API";

// TYPES
import { ProductComponent } from "../productComponent/ProductComponent";
import { useFetch } from "../../../hooks/useHookFetchData";
import { ListImagesComponent } from "../../../images/components/listImagesComponent/ListImagesComponent";
import type { Producto } from "../../types/productComponent/ProductComponentType";

export const ListProductsComponent = () => {
  const { data, loading, error } = useFetch<Producto[]>(API_PRODUCTOS);

  if (loading) return <p>Cargando productos...</p>;
  if (error) return <p>Error al cargar productos: {error}</p>;

  return (
    <div style={{ padding: "2rem" }}>
      {data?.map((product) => (
        <div key={product.ID} style={{ marginBottom: "2rem" }}>
          <article>
            <header>
              <ProductComponent product={product} />
            </header>

            <div>
              <ListImagesComponent productId={product.ID} />
            </div>
          </article>
        </div>
      ))}
    </div>
  );
};
