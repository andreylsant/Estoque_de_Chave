import { useState } from "react";
import type { Produto } from "../types/produto";
import "./StockModal.css";

interface StockModalProps {
  produto: Produto;
  tipo: "entrada" | "saida";
  onConfirm: (quantidade: number) => void;
  onClose: () => void;
}

export function StockModal({
  produto,
  tipo,
  onConfirm,
  onClose,
}: StockModalProps) {
  const [quantidade, setQuantidade] = useState("");

  function handleConfirm() {
    const valor = Number(quantidade);

    if (!Number.isInteger(valor) || valor <= 0) {
      alert("Digite uma quantidade válida.");
      return;
    }

    if (tipo === "saida" && valor > produto.quantidade) {
      alert("Não há quantidade suficiente no estoque.");
      return;
    }

    onConfirm(valor);
  }

  const entrada = tipo === "entrada";

  return (
    <div className="modal-overlay">
      <div className="stock-modal">
        <h3>
          {entrada
            ? "Adicionar ao estoque"
            : "Retirar do estoque"}
        </h3>

        <p>
          {produto.marca} {produto.modelo}
        </p>

        <div className="current-stock">
          Estoque atual: <strong>{produto.quantidade}</strong>
        </div>

        <label htmlFor="quantidade">
          Quantidade
        </label>

        <input
          id="quantidade"
          type="number"
          min="1"
          value={quantidade}
          onChange={(event) =>
            setQuantidade(event.target.value)
          }
          placeholder="Digite a quantidade"
          autoFocus
        />

        <div className="modal-actions">
          <button
            type="button"
            className="button-cancel"
            onClick={onClose}
          >
            Cancelar
          </button>

          <button
            type="button"
            className={
              entrada
                ? "button-confirm-add"
                : "button-confirm-remove"
            }
            onClick={handleConfirm}
          >
            {entrada ? "Adicionar" : "Retirar"}
          </button>
        </div>
      </div>
    </div>
  );
}