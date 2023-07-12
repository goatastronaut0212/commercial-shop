import {BrowserRouter, Link, Route , Routes} from "react-router-dom"
import Home from "../src/pages/Home";
import Product from "../src/pages/Product";
import ProductDetail from "../src/pages/ProductDetail";
import New from "../src/pages/New";
import About from "../src/pages/About";

function App() {
  return (
    <>

    <BrowserRouter>
      <Routes>
          <Route path="/" element={ <Home />} />
          
          <Route path="/product" element={ <Product /> } >
              <Route path="product/:id" element={ <ProductDetail /> } />
          </Route>

          <Route path="/new" element={ <New />} />
          <Route path="/about" element={ <About />} />

          <Route path="*" element={<h1>Not Found The Page</h1>} />
      </Routes>
    </BrowserRouter>
    </>

  )
}

export default App
