import { useFetch } from "../../../hooks/useHookFetchData";
import { ProductComponent } from "../productComponent/ProductComponent";
import { ListImagesComponent } from "../../../images/components/listImagesComponent/ListImagesComponent";
import type { Producto } from "../../types/productComponent/ProductComponentType";
import { API_PRODUCTOS } from "../../../utils/API/key_API";

export const ListProductsComponent = () => {
  const {
    data: productos,
    loading: loadingProductos,
    error: errorProductos,
  } = useFetch<Producto[]>(API_PRODUCTOS);

  if (loadingProductos) {
    return <p>Cargando productos...</p>;
  }

  if (errorProductos) {
    return <p>Error al cargar productos.</p>;
  }

  const DATO = productos?.map((product) => {
    return product.ID;
  });

  console.log(DATO);

  return (
    <div style={{ padding: "2rem" }}>
      {productos?.map((product) => (
        <div key={product.ID} style={{ marginBottom: "2rem" }}>
          <article>
            <header>
              <ProductComponent product={product} key={product.ID} />
            </header>

            <div>
              <ListImagesComponent tipo="productos" productId={product.ID} />
            </div>
          </article>
        </div>
      ))}
    </div>
  );
};
