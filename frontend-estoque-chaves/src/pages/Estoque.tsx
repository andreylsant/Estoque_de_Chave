import { useEffect, useMemo, useState } from "react";
import { ProductTable } from "../componets/ProductTable";
import { StockModal } from "../componets/StockModal";
import { buscarProdutos, movimentarEstoque } from "../services/produtoService";
import type { Produto } from "../types/produto";

export function Estoque() {
  const [produtos, setProdutos] = useState<Produto[]>([]);
  const [filtro, setFiltro] = useState("");
  const [ordem, setOrdem] = useState<"maior" | "menor">("menor");
  const [somenteBaixo, setSomenteBaixo] = useState(false);

  const [modalAberto, setModalAberto] = useState(false);
  const [produtoSelecionado, setProdutoSelecionado] =
    useState<Produto | null>(null);

  const [tipoMovimento, setTipoMovimento] =
    useState<"entrada" | "saida">("entrada");

  useEffect(() => {
    buscarProdutos()
      .then(setProdutos)
      .catch((error) => {
        console.error("Erro ao buscar produtos:", error);
      });
  }, []);

  const produtosFiltrados = useMemo(() => {
    let resultado = [...produtos];

    if (filtro) {
      resultado = resultado.filter((produto) =>
        `${produto.marca} ${produto.modelo}`
          .toLowerCase()
          .includes(filtro.toLowerCase())
      );
    }

    if (somenteBaixo) {
      resultado = resultado.filter(
        (produto) => produto.quantidade <= 200
      );
    }

    resultado.sort((a, b) => {
      if (ordem === "maior") {
        return b.quantidade - a.quantidade;
      }

      return a.quantidade - b.quantidade;
    });

    return resultado;
  }, [produtos, filtro, ordem, somenteBaixo]);

  function abrirModal(
    produto: Produto,
    tipo: "entrada" | "saida"
  ) {
    setProdutoSelecionado(produto);
    setTipoMovimento(tipo);
    setModalAberto(true);
  }

  async function confirmarMovimento(quantidade: number) {
    if (!produtoSelecionado) {
      return;
    }

    try {
      await movimentarEstoque(produtoSelecionado.id, quantidade, tipoMovimento);
    } catch (error) {
      console.error("Erro ao movimentar estoque:", error);
      alert("Não foi possível atualizar o estoque no servidor.");
      return;
    }

    setProdutos((produtosAtuais) =>
      produtosAtuais.map((produto) => {
        if (produto.id !== produtoSelecionado.id) {
          return produto;
        }

        return {
          ...produto,
          quantidade:
            tipoMovimento === "entrada"
              ? produto.quantidade + quantidade
              : produto.quantidade - quantidade,
        };
      })
    );

    setModalAberto(false);
    setProdutoSelecionado(null);
  }

  function fecharModal() {
    setModalAberto(false);
    setProdutoSelecionado(null);
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Estoque</h2>
          <p>Consulte e gerencie suas chaves.</p>
        </div>
      </div>

      <div className="filters">
        <input
          type="text"
          placeholder="Buscar marca ou modelo..."
          value={filtro}
          onChange={(event) =>
            setFiltro(event.target.value)
          }
        />

        <select
          value={ordem}
          onChange={(event) =>
            setOrdem(
              event.target.value as "maior" | "menor"
            )
          }
        >
          <option value="menor">
            Menor estoque primeiro
          </option>

          <option value="maior">
            Maior estoque primeiro
          </option>
        </select>

        <label className="checkbox">
          <input
            type="checkbox"
            checked={somenteBaixo}
            onChange={(event) =>
              setSomenteBaixo(event.target.checked)
            }
          />

          Estoque abaixo de 200
        </label>
      </div>

      <ProductTable
        produtos={produtosFiltrados}
        onAdd={(produto) =>
          abrirModal(produto, "entrada")
        }
        onRemove={(produto) =>
          abrirModal(produto, "saida")
        }
      />

      {modalAberto && produtoSelecionado && (
        <StockModal
          produto={produtoSelecionado}
          tipo={tipoMovimento}
          onConfirm={confirmarMovimento}
          onClose={fecharModal}
        />
      )}
    </div>
  );
}