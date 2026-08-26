import { useState } from "react";
import { StockCard } from "../componets/StockCard";
import { ProductTable } from "../componets/ProductTable";
import { produtosIniciais } from "../data/produtos";
import { Produto } from "../types/produto";

export function Dashboard() {
  const [produtos, setProdutos] = useState<Produto[]>(
    produtosIniciais
  );

  const [produtoSelecionado, setProdutoSelecionado] =
    useState<Produto | null>(null);

  const totalProdutos = produtos.length;

  const totalEstoque = produtos.reduce(
    (total, produto) => total + produto.quantidade,
    0
  );

  const estoqueBaixo = produtos.filter(
    (produto) => produto.quantidade <= 200
  ).length;

  function editarProduto(produto: Produto) {
    setProdutoSelecionado(produto);
  }

  function salvarQuantidade() {
    if (!produtoSelecionado) {
      return;
    }

    setProdutos((produtosAtuais) =>
      produtosAtuais.map((produto) =>
        produto.id === produtoSelecionado.id
          ? produtoSelecionado
          : produto
      )
    );

    setProdutoSelecionado(null);
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Dashboard</h2>
          <p>Visão geral do seu estoque.</p>
        </div>
      </div>

      <div className="cards-grid">
        <StockCard
          title="Modelos cadastrados"
          value={totalProdutos}
          description="Modelos de chaves"
          icon="📦"
        />

        <StockCard
          title="Total em estoque"
          value={totalEstoque}
          description="Chaves disponíveis"
          icon="🔑"
        />

        <StockCard
          title="Estoque baixo"
          value={estoqueBaixo}
          description="Precisam de atenção"
          icon="⚠️"
        />
      </div>

      <div className="section-header">
        <div>
          <h3>Estoque atual</h3>
          <p>Visão rápida dos produtos cadastrados.</p>
        </div>
      </div>

      <ProductTable
        produtos={produtos}
        onEdit={editarProduto}
      />

      {produtoSelecionado && (
        <div className="modal-overlay">
          <div className="modal">
            <h3>Editar quantidade</h3>

            <p>
              {produtoSelecionado.marca}{" "}
              {produtoSelecionado.modelo}
            </p>

            <label>
              Quantidade

              <input
                type="number"
                min="0"
                value={produtoSelecionado.quantidade}
                onChange={(event) =>
                  setProdutoSelecionado({
                    ...produtoSelecionado,
                    quantidade: Number(event.target.value),
                  })
                }
              />
            </label>

            <div className="modal-actions">
              <button
                className="button-secondary"
                onClick={() => setProdutoSelecionado(null)}
              >
                Cancelar
              </button>

              <button
                className="button-primary"
                onClick={salvarQuantidade}
              >
                Salvar
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}