import { type FormEvent, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { api } from "../services/api";
import "./criandoconta.css";

export function CriarConta() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [senha, setSenha] = useState("");
  const [confirmarSenha, setConfirmarSenha] = useState("");

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();

    if (!email || !senha || !confirmarSenha) {
      alert("Preencha todos os campos.");
      return;
    }

    if (senha !== confirmarSenha) {
      alert("As senhas não são iguais.");
      return;
    }

    try {
      await api.post("/criar-conta", {
        email,
        senha,
      });

      alert("Conta criada com sucesso!");

      navigate("/login");
    } catch (error) {
      console.error("Erro ao criar conta:", error);
      alert("Não foi possível criar a conta.");
    }
  }

  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          🔑
        </div>

        <h1>Criar conta</h1>

        <p>Crie seu acesso ao KeyStock</p>

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

          <div className="form-group">
            <label htmlFor="confirmarSenha">
              Confirmar senha
            </label>

            <input
              id="confirmarSenha"
              type="password"
              placeholder="Digite a senha novamente"
              value={confirmarSenha}
              onChange={(event) =>
                setConfirmarSenha(event.target.value)
              }
            />
          </div>

          <button
            className="button-primary login-button"
            type="submit"
          >
            Criar conta
          </button>

        </form>

        <div className="login-links">
          <Link to="/login">
            Já tenho uma conta
          </Link>
        </div>

      </div>
    </div>
  );
}
