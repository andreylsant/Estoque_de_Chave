import { useMemo, useState } from "react";
import { produtosIniciais } from "../data/produtos";
import { ProductTable } from "../componets/ProductTable";
import { StockModal } from "../componets/StockModal";
import type { Produto } from "../types/produto";

export function Estoque() {
  const [produtos, setProdutos] = useState<Produto[]>(
    produtosIniciais
  );

  const [filtro, setFiltro] = useState("");

  const [ordem, setOrdem] = useState<"maior" | "menor">(
    "menor"
  );

  const [somenteBaixo, setSomenteBaixo] = useState(false);

  const [modalAberto, setModalAberto] = useState(false);

  const [produtoSelecionado, setProdutoSelecionado] =
    useState<Produto | null>(null);

  const [tipoMovimento, setTipoMovimento] =
    useState<"entrada" | "saida">("entrada");

  const produtosFiltrados = useMemo(() => {
    let resultado = [...produtos];

    // Filtro por marca ou modelo
    if (filtro) {
      resultado = resultado.filter((produto) =>
        `${produto.marca} ${produto.modelo}`
          .toLowerCase()
          .includes(filtro.toLowerCase())
      );
    }

    // Filtrar somente produtos com estoque baixo
    if (somenteBaixo) {
      resultado = resultado.filter(
        (produto) => produto.quantidade <= 200
      );
    }

    // Ordenação
    resultado.sort((a, b) => {
      if (ordem === "maior") {
        return b.quantidade - a.quantidade;
      }

      return a.quantidade - b.quantidade;
    });

    return resultado;
  }, [produtos, filtro, ordem, somenteBaixo]);

  // Abre o modal para entrada ou saída
  function abrirModal(
    produto: Produto,
    tipo: "entrada" | "saida"
  ) {
    setProdutoSelecionado(produto);
    setTipoMovimento(tipo);
    setModalAberto(true);
  }

  // Confirma a movimentação do estoque
  function confirmarMovimento(quantidade: number) {
    if (!produtoSelecionado) {
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

  // Fecha o modal
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