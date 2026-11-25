import {type FormEvent, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { api } from "../services/api";
export function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");
  const [carregando, setCarregando] = useState(false);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();

    if (!email || !senha) {
      alert("Informe email e senha.");
      return;
    }

    try {
      setCarregando(true);

      const response = await api.post("/login", {
        email,
        senha,
      });

      const token = response.data.token;

      localStorage.setItem("token", token);

      navigate("/dashboard");
    } catch (error) {
      console.error("Erro ao fazer login:", error);
      alert("Email ou senha inválidos.");
    } finally {
      setCarregando(false);
    }
  }

  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          🔑
        </div>

        <h1>KeyStock</h1>

        <p>Controle de estoque de chaves</p>

        <form onSubmit={handleSubmit}>

          <div className="form-group">
            <label htmlFor="email">Email</label>

            <input
              id="email"
              type="email"
              placeholder="seu@chipcode.com"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
            />
          </div>

          <div className="form-group">
            <label htmlFor="senha">Senha</label>

            <input
              id="senha"
              type="password"
              placeholder="Digite sua senha"
              value={senha}
              onChange={(event) => setSenha(event.target.value)}
            />
          </div>

          <button
            className="button-primary login-button"
            type="submit"
            disabled={carregando}
          >
            {carregando ? "Entrando..." : "Entrar"}
          </button>

        </form>

        <div className="login-links">

          <Link to="/recuperar-senha">
            Esqueci minha senha
          </Link>

          <span>•</span>

          <Link to="/criar-conta">
            Criar uma conta
          </Link>

        </div>

      </div>
    </div>
  );
}