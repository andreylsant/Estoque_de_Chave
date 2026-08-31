import { FormEvent, useState } from "react";

export function NovoProduto() {
  const [marca, setMarca] = useState("");
  const [modelo, setModelo] = useState("");
  const [quantidade, setQuantidade] = useState("");

  function handleSubmit(event: FormEvent) {
    event.preventDefault();

    if (!marca || !modelo || !quantidade) {
      alert("Preencha todos os campos.");
      return;
    }

    console.log({
      marca,
      modelo,
      quantidade: Number(quantidade),
    });

    alert("Produto cadastrado!");

    setMarca("");
    setModelo("");
    setQuantidade("");
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Novo produto</h2>
          <p>Cadastre uma nova chave no estoque.</p>
        </div>
      </div>

      <form
        className="form-card"
        onSubmit={handleSubmit}
      >
        <div className="form-group">
          <label>Marca</label>

          <input
            type="text"
            placeholder="Ex: Fiat"
            value={marca}
            onChange={(event) =>
              setMarca(event.target.value)
            }
          />
        </div>

        <div className="form-group">
          <label>Modelo</label>

          <input
            type="text"
            placeholder="Ex: Uno"
            value={modelo}
            onChange={(event) =>
              setModelo(event.target.value)
            }
          />
        </div>

        <div className="form-group">
          <label>Quantidade</label>

          <input
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
        >
          Cadastrar produto
        </button>
      </form>
    </div>
  );
}