import type { Producto } from "../../types/productComponent/ProductComponentType";

type ProductComponentProps = {
  product: Producto;
};

export const ProductComponent: React.FC<ProductComponentProps> = ({
  product,
}) => {
  return (
    <div key={product.ID}>
      <h3>{product.Nombre}</h3>
      <p>Precio: {product.Precio}</p>
      <p>Stock: {product.Stock}</p>
      <p>Categoría: {product.Categoria}</p>
    </div>
  );
};
