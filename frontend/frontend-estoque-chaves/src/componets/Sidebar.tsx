import { NavLink } from "react-router-dom";

export function Sidebar() {
    return(
         <aside className="sidebar">
      <div className="logo">
        <div className="logo-icon">🔑</div>

        <div>
          <strong>KeyStock</strong>
          <span>Estoque de Chaves</span>
        </div>
      </div>

      <nav className="menu">
        <NavLink to="/dashboard">
          📊 Dashboard
        </NavLink>

        <NavLink to="/estoque">
          📦 Estoque
        </NavLink>

        <NavLink to="/produtos/novo">
          ➕ Novo produto
        </NavLink>

        <NavLink to="/pedidos">
          🛒 Pedidos
        </NavLink>

        <NavLink to="/relatorios">
          📄 Relatórios
        </NavLink>
      </nav>

      <div className="sidebar-footer">
        <span>Sistema de estoque</span>
        <small>v1.0.0</small>
      </div>
    </aside>
    );
}