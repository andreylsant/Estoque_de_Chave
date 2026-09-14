import { useEffect, useMemo, useState } from "react";
import { StockCard } from "../componets/StockCard";
import { buscarProdutos } from "../services/produtoService";
import type { Produto } from "../types/produto";
import "./dashboardestoque.css";

export function Dashboard() {
  const [produtos, setProdutos] = useState<Produto[]>([]);

  useEffect(() => {
    buscarProdutos()
      .then(setProdutos)
      .catch((error) => {
        console.error("Erro ao buscar produtos:", error);
      });
  }, []);

  const totalProdutos = produtos.length;

  const totalEstoque = useMemo(() => {
    return produtos.reduce(
      (total, produto) => total + produto.quantidade,
      0
    );
  }, [produtos]);

  const estoqueBaixo = useMemo(() => {
    return produtos.filter(
      (produto) => produto.quantidade <= 200
    ).length;
  }, [produtos]);

  const estoqueNormal = useMemo(() => {
    return produtos.filter(
      (produto) => produto.quantidade > 200
    ).length;
  }, [produtos]);

  const estoqueCritico = useMemo(() => {
    return produtos.filter(
      (produto) => produto.quantidade <= 100
    ).length;
  }, [produtos]);

  const estoqueAbaixoDoLimite = useMemo(() => {
    return produtos.filter(
      (produto) => produto.quantidade <= 200
    );
  }, [produtos]);

  return (
    <div className="dashboard">
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

      <section className="dashboard-section">
        <div className="section-header">
          <div>
            <h3>Situação do estoque</h3>
            <p>Resumo da situação atual das suas chaves.</p>
          </div>
        </div>

        <div className="stock-status-grid">
          <div className="status-card normal">
            <div className="status-icon">🟢</div>

            <div>
              <span>Estoque normal</span>
              <strong>{estoqueNormal}</strong>
              <small>modelos</small>
            </div>
          </div>

          <div className="status-card low">
            <div className="status-icon">🟡</div>

            <div>
              <span>Estoque baixo</span>
              <strong>{estoqueBaixo}</strong>
              <small>modelos</small>
            </div>
          </div>

          <div className="status-card critical">
            <div className="status-icon">🔴</div>

            <div>
              <span>Estoque crítico</span>
              <strong>{estoqueCritico}</strong>
              <small>modelos</small>
            </div>
          </div>
        </div>
      </section>

      <section className="dashboard-section">
        <div className="section-header">
          <div>
            <h3>⚠️ Chaves que precisam de atenção</h3>
            <p>
              Modelos com 200 chaves ou menos em estoque.
            </p>
          </div>
        </div>

        <div className="attention-list">
          {estoqueAbaixoDoLimite.length === 0 ? (
            <div className="empty-state">
              <span>✅</span>
              <p>
                Nenhuma chave precisa de reposição.
              </p>
            </div>
          ) : (
            estoqueAbaixoDoLimite.map((produto) => {
              const critico = produto.quantidade <= 100;

              return (
                <div
                  className="attention-item"
                  key={produto.id}
                >
                  <div className="attention-product">
                    <div className="key-icon">🔑</div>

                    <div>
                      <strong>
                        {produto.marca} {produto.modelo}
                      </strong>

                      <span>
                        {critico
                          ? "Estoque crítico"
                          : "Estoque baixo"}
                      </span>
                    </div>
                  </div>

                  <div
                    className={
                      critico
                        ? "attention-quantity critical"
                        : "attention-quantity low"
                    }
                  >
                    <strong>{produto.quantidade}</strong>
                    <span>chaves</span>
                  </div>
                </div>
              );
            })
          )}
        </div>
      </section>

      <section className="dashboard-section">
        <div className="section-header">
          <div>
            <h3>🔄 Últimas movimentações</h3>
            <p>Resumo das últimas entradas e saídas.</p>
          </div>
        </div>

        <div className="movements-empty">
          <span>📋</span>

          <div>
            <strong>
              Nenhuma movimentação registrada
            </strong>

            <p>
              As entradas e saídas realizadas no estoque
              aparecerão aqui.
            </p>
          </div>
        </div>
      </section>
    </div>
  );
}