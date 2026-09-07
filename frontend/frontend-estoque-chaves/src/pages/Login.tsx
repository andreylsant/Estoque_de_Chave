import {type FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { api } from "../services/api";

export function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");

  async function handleSubmit(event: FormEvent) {
  event.preventDefault();

  console.log("verificando email e senha")
  if (!email || !senha) {
    alert("Informe email e senha.");
    return;
  }

  try {
    const response = await api.post("/login", {
      email: email,
      senha: senha,
    });

    const token = response.data.token;

    localStorage.setItem("token", token);

    navigate("/dashboard");

  } catch (error) {
  console.error("ERRO:", error);
  alert("Erro ao conectar com o servidor.");

  }
}
  return (
    <div className="login-page">
      <div className="login-card">
        <div className="login-logo">
          🔑
        </div>

        <h1>KeyStock</h1>

        <p>
          Controle de estoque de chaves
        </p>

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Email</label>

            <input
              type="email"
              placeholder="seu@chipcode.com"
              value={email}
              onChange={(event) =>
                setEmail(event.target.value)
              }
            />
          </div>

          <div className="form-group">
            <label>Senha</label>

            <input
              type="password"
              placeholder="Digite sua senha"
              value={senha}
              onChange={(event) =>
                setSenha(event.target.value)
              }
            />
          </div>

          <button
            className="button-primary login-button"
            type="submit"
          >
            Entrar
          </button>
        </form>
      </div>
    </div>
  );
}