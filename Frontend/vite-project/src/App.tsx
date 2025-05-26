import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

type Producto = {
  ID: string;
  Nombre: string;
  Precio: number;
  Stock: number;
  Categoria: string;
  FechaCreacion: string;
  FechaActualizacion: string;
};

type ResponseAPI = {
  status: number;
  message: string;
  data: Producto[];
};

function App() {
  const [count, setCount] = useState(0);
  const [data, setData] = useState<Producto[]>([]);

  const API =
    "https://8081-firebase-turbomamadascolaborativascompany-1748038475149.cluster-aj77uug3sjd4iut4ev6a4jbtf2.cloudworkstations.dev/api/productos";

  useEffect(() => {
    const fetchData = async (URL: string) => {
      try {
        const response = await fetch(URL);
        const json: ResponseAPI = await response.json();
        console.log("Respuesta de la API:", json);
        if (Array.isArray(json.data)) {
          setData(json.data);
        }
      } catch (error) {
        console.error("Error al obtener productos:", error);
        alert("ERROR al cargar productos");
      }
    };

    fetchData(API);
  }, []);

  return (
    <>
      {data.map((item) => (
        <div key={item.ID}>
          <p>{item.Nombre}</p>
          <p>${item.Precio}</p>
          <p>Stock: {item.Stock}</p>
        </div>
      ))}

      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
