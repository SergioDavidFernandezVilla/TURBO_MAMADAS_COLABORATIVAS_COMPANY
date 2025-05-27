import "./App.css";
import { NavbarComponent } from "./components/Navbar/NavbarComponent";
import { ProductGalleryCarousel } from "./productGalleryCarousel/components/productGalleryCarousel";
import { ListProductsComponent } from "./products/components/lIstProductsComponent/ListProductsComponent";

function App() {
  return (
    <>
      <NavbarComponent />
     <ProductGalleryCarousel />

      <div className={"container__app__"}>
        <main>
          <ListProductsComponent />
        </main>
      </div>
    </>
  );
}

export default App;
