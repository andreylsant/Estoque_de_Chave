import type { Produto } from "../types/produto";

interface ProductTableProps {
  produtos: Produto[];
  onEdit: (produto: Produto) => void;
}

export function ProductTable({
  produtos,
  onEdit,
}: ProductTableProps) {
  function getStockStatus(quantidade: number) {
    if (quantidade <= 100) {
      return "critical";
    }

    if (quantidade <= 200) {
      return "low";
    }

    return "normal";
  }

  function getStatusLabel(quantidade: number) {
    if (quantidade <= 100) {
      return "Crítico";
    }

    if (quantidade <= 200) {
      return "Baixo";
    }

    return "Normal";
  }

  return (
    <div className="table-container">
      <table>
        <thead>
          <tr>
            <th>Marca</th>
            <th>Modelo</th>
            <th>Quantidade</th>
            <th>Status</th>
            <th>Ação</th>
          </tr>
        </thead>

        <tbody>
          {produtos.map((produto) => {
            const status = getStockStatus(produto.quantidade);

            return (
              <tr key={produto.id}>
                <td>{produto.marca}</td>

                <td>
                  <strong>{produto.modelo}</strong>
                </td>

                <td>
                  <strong>{produto.quantidade}</strong>
                </td>

                <td>
                  <span className={`status ${status}`}>
                    {getStatusLabel(produto.quantidade)}
                  </span>
                </td>

                <td>
                  <button
                    className="button-secondary"
                    onClick={() => onEdit(produto)}
                  >
                    Editar
                  </button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}