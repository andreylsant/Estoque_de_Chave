import { type FormEvent, useState } from "react";
import { Link } from "react-router-dom";
import { api } from "../services/api";
import "./recuperarsenha.css";

export function RecuperarSenha() {
  const [email, setEmail] = useState("");
  const [enviado, setEnviado] = useState(false);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();

    if (!email) {
      alert("Informe seu email.");
      return;
    }

    try {
      await api.post("/recuperar-senha", {
        email,
      });

      setEnviado(true);
    } catch (error) {
      console.error("Erro ao recuperar senha:", error);

      alert(
        "Não foi possível solicitar a recuperação da senha."
      );
    }
  }

  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          🔑
        </div>

        <h1>Recuperar senha</h1>

        {!enviado ? (
          <>
            <p>
              Informe seu email para recuperar sua senha.
            </p>

            <form onSubmit={handleSubmit}>

              <div className="form-group">
                <label htmlFor="email">Email</label>

                <input
                  id="email"
                  type="email"
                  placeholder="seu@chipcode.com"
                  value={email}
                  onChange={(event) =>
                    setEmail(event.target.value)
                  }
                />
              </div>

              <button
                className="button-primary login-button"
                type="submit"
              >
                Recuperar senha
              </button>

            </form>
          </>
        ) : (
          <p>
            Se o email estiver cadastrado, você receberá
            as instruções para recuperar sua senha.
          </p>
        )}

        <div className="login-links">
          <Link to="/login">
            Voltar para o login
          </Link>
        </div>

      </div>
    </div>
  );
}
