import "./App.css";
import { NavbarComponent } from "./components/Navbar/NavbarComponent";
import { ListProductsComponent } from "./products/components/lIstProductsComponent/ListProductsComponent";

function App() {
  return (
    <>
      <NavbarComponent />
      <ListProductsComponent />

      <div className={"container__app__"}>
        <main>
          <ListProductsComponent />
        </main>
      </div>
    </>
  );
}

export default App;
