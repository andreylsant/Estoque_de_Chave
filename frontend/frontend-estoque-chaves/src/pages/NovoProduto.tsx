import { type FormEvent, useState } from "react";
import { api } from "../services/api";
import "./nProduto.css";

export function NovoProduto() {
  const [marca, setMarca] = useState("");
  const [modelo, setModelo] = useState("");
  const [quantidade, setQuantidade] = useState("");

  const [mensagem, setMensagem] = useState("");
  const [erro, setErro] = useState("");
  const [carregando, setCarregando] = useState(false);

    function mostrarMensagem(texto: string) {
      setMensagem(texto);

      setTimeout(() => {
        setMensagem("");
      }, 3000);
    }

    function mostrarErro(texto: string) {
      setErro(texto);

      setTimeout(() => {
        setErro("");
      }, 3000);
    }

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();

    setMensagem("");
    setErro("");

    if (!marca.trim() || !modelo.trim() || !quantidade) {
      mostrarErro("Preencha todos os campos.");
      return;
    }

    const quantidadeNumero = Number(quantidade);

    if (!Number.isInteger(quantidadeNumero) || quantidadeNumero < 0) {
      mostrarErro("Digite uma quantidade válida.");
      return;
    }

    try {
      setCarregando(true);

      await api.post("/produtos", {
        marca: marca.trim(),
        modelo: modelo.trim(),
        quantidade: quantidadeNumero,
      });

      mostrarMensagem("Chave cadastrada com sucesso!");

      setMarca("");
      setModelo("");
      setQuantidade("");
    } catch (error) {
      console.error("Erro ao cadastrar chave:", error);

      mostrarErro(
        "Não foi possível cadastrar a chave."
      );
    } finally {
      setCarregando(false);
    }
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Nova Chave</h2>
          <p>Cadastre uma nova chave no estoque.</p>
        </div>
      </div>

      {mensagem && (
        <div className="message success">
          {mensagem}
        </div>
      )}

      {erro && (
        <div className="message error">
          {erro}
        </div>
      )}

      <form
        className="form-card"
        onSubmit={handleSubmit}
      >
        <div className="form-group">
          <label htmlFor="marca">Marca</label>

          <input
            id="marca"
            type="text"
            placeholder="Ex: MGM"
            value={marca}
            onChange={(event) =>
              setMarca(event.target.value)
            }
          />
        </div>

        <div className="form-group">
          <label htmlFor="modelo">Modelo</label>

          <input
            id="modelo"
            type="text"
            placeholder="Ex: 1080"
            value={modelo}
            onChange={(event) =>
              setModelo(event.target.value)
            }
          />
        </div>

        <div className="form-group">
          <label htmlFor="quantidade">
            Quantidade
          </label>

          <input
            id="quantidade"
            type="number"
            min="0"
            placeholder="Ex: 100"
            value={quantidade}
            onChange={(event) =>
              setQuantidade(event.target.value)
            }
          />
        </div>

        <button
          type="submit"
          className="button-primary"
          disabled={carregando}
        >
          {carregando
            ? "Cadastrando..."
            : "Cadastrar chave"}
        </button>
      </form>
    </div>
  );
}