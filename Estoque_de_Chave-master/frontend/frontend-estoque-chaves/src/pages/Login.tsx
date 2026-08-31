import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

export function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");

  function handleSubmit(event: FormEvent) {
    event.preventDefault();

    if (!email || !senha) {
      alert("Informe email e senha.");
      return;
    }

    /*
     * Temporariamente simulando login.
     *
     * Depois:
     *
     * POST /api/auth/login
     *
     * O backend Go retornará um JWT.
     */

    localStorage.setItem(
      "token",
      "token-temporario"
    );

    navigate("/dashboard");
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
              placeholder="seu@email.com"
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