import { Outlet } from "react-router-dom";
import { Sidebar } from "./Sidebar";
import { Header } from "./Header";
import "./layout.css";

export function Layout() {
  return (
    <div className="app-layout">
      <Sidebar />

      <main className="main-content">
        <Header />

        <section className="page-content">
          <Outlet />
        </section>
      </main>
    </div>
  );
}