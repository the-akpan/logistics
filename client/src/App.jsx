import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import { Nav, Footer } from "./components/partials";
import { Dashboard, TrackingItem } from "./pages/cms";

function App() {
  return (
    <>
      <Nav />

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />

        <Route path="dashboard" element={<Dashboard />}>
          <Route path="track-item" element={<TrackingItem />} />
        </Route>
      </Routes>

      <Footer />
    </>
  );
}

export default App;
