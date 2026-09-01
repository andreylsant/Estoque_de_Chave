import { useMemo, useState } from "react";
import { produtosIniciais } from "../data/produtos";
import { ProductTable } from "../componets/ProductTable";
import type { Produto } from "../types/produto";

export function Estoque() {
  const [produtos, setProdutos] = useState<Produto[]>(
    produtosIniciais
  );

  const [filtro, setFiltro] = useState("");

  const [ordem, setOrdem] = useState<
    "maior" | "menor"
  >("menor");

  const [somenteBaixo, setSomenteBaixo] =
    useState(false);

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

  function editarProduto(produto: Produto) {
    const novaQuantidade = window.prompt(
      `Quantidade para ${produto.marca} ${produto.modelo}:`,
      String(produto.quantidade)
    );

    if (novaQuantidade === null) {
      return;
    }

    const quantidade = Number(novaQuantidade);

    if (isNaN(quantidade) || quantidade < 0) {
      alert("Digite uma quantidade válida.");
      return;
    }

    setProdutos((produtosAtuais) =>
      produtosAtuais.map((item) =>
        item.id === produto.id
          ? { ...item, quantidade }
          : item
      )
    );
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
        onEdit={editarProduto}
      />
    </div>
  );
}