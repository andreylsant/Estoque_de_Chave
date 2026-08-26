import { produtosIniciais } from "../data/produtos";

export function Relatorios() {
  const total = produtosIniciais.reduce(
    (soma, produto) =>
      soma + produto.quantidade,
    0
  );

  const baixo = produtosIniciais.filter(
    (produto) => produto.quantidade <= 200
  );

  function exportarCSV() {
    const header =
      "Marca,Modelo,Quantidade\n";

    const rows = produtosIniciais
      .map(
        (produto) =>
          `${produto.marca},${produto.modelo},${produto.quantidade}`
      )
      .join("\n");

    const blob = new Blob(
      [header + rows],
      {
        type: "text/csv;charset=utf-8;",
      }
    );

    const url =
      URL.createObjectURL(blob);

    const link =
      document.createElement("a");

    link.href = url;
    link.download = "relatorio-estoque.csv";

    link.click();

    URL.revokeObjectURL(url);
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Relatórios</h2>
          <p>
            Informações gerais do estoque.
          </p>
        </div>

        <button
          className="button-primary"
          onClick={exportarCSV}
        >
          Exportar CSV
        </button>
      </div>

      <div className="cards-grid">
        <div className="report-card">
          <span>Total de chaves</span>
          <strong>{total}</strong>
        </div>

        <div className="report-card">
          <span>Produtos cadastrados</span>
          <strong>
            {produtosIniciais.length}
          </strong>
        </div>

        <div className="report-card">
          <span>Estoque baixo</span>
          <strong>{baixo.length}</strong>
        </div>
      </div>

      <div className="section-header">
        <h3>Produtos com estoque baixo</h3>
      </div>

      <div className="table-container">
        <table>
          <thead>
            <tr>
              <th>Marca</th>
              <th>Modelo</th>
              <th>Quantidade</th>
            </tr>
          </thead>

          <tbody>
            {baixo.map((produto) => (
              <tr key={produto.id}>
                <td>{produto.marca}</td>
                <td>{produto.modelo}</td>
                <td>
                  {produto.quantidade}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}