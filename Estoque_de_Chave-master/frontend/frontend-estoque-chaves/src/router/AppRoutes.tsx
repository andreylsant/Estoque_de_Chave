import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
} from "react-router-dom";

import { Layout } from "../componets/layout";
import { Login } from "../pages/Login";
import { Dashboard } from "../pages/Dashboard";
import { Estoque } from "../pages/Estoque";
import { NovoProduto } from "../pages/NovoProduto";
import { Pedidos } from "../pages/Pedido";
import { Relatorios } from "../pages/Relatorios";

function ProtectedRoute() {
  const token = localStorage.getItem("token");

  if (!token) {
    return <Navigate to="/login" replace />;
  }

  return <Layout />;
}

export function AppRoutes() {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/login"
          element={<Login />}
        />

        <Route element={<ProtectedRoute />}>
          <Route
            path="/dashboard"
            element={<Dashboard />}
          />

          <Route
            path="/estoque"
            element={<Estoque />}
          />

          <Route
            path="/produtos/novo"
            element={<NovoProduto />}
          />

          <Route
            path="/pedidos"
            element={<Pedidos />}
          />

          <Route
            path="/relatorios"
            element={<Relatorios />}
          />
        </Route>

        <Route
          path="*"
          element={
            <Navigate
              to="/dashboard"
              replace
            />
          }
        />
      </Routes>
    </BrowserRouter>
  );
}