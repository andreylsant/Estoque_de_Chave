import { NavLink } from "react-router-dom";
const links = [
  { section: "VISÃO GERAL", items: [{ to: "/dashboard", icon: "▦", label: "Dashboard" }] },
  { section: "ESTOQUE & CATÁLOGO", items: [{ to: "/estoque", icon: "◇", label: "Produtos" }, { to: "/produtos/novo", icon: "＋", label: "Cadastrar chave" }] },
  { section: "OPERAÇÕES", items: [{ to: "/pedidos", icon: "🛒", label: "Pedidos" }, { to: "/relatorios", icon: "▤", label: "Relatórios" }] },
];
export function Sidebar() { return <aside className="sidebar"><div className="logo"><div className="logo-icon">🔑</div><div><strong>KeyStock</strong><span>GESTÃO INTELIGENTE</span></div></div><nav className="menu">{links.map(group => <div key={group.section}><div className="menu-label">{group.section}</div>{group.items.map(item => <NavLink key={item.to} to={item.to}><span>{item.icon}</span>{item.label}</NavLink>)}</div>)}</nav><div className="sidebar-footer">Sistema de estoque<small>KeyStock · v1.0.0</small></div></aside>; }
