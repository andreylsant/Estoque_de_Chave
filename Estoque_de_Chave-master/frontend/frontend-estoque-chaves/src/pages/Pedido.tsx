import { useMemo } from "react";
import { produtosIniciais } from "../data/produtos";

export function Pedidos() {
  const produtosParaPedido = useMemo(() => {
    return produtosIniciais
      .filter((produto) => produto.quantidade <= 200)
      .sort((a, b) => a.quantidade - b.quantidade);
  }, []);

  function gerarMensagem() {
    if (produtosParaPedido.length === 0) {
      return "Não existem produtos com estoque baixo.";
    }

    let mensagem =
      "Olá, gostaria de fazer um pedido de chaves:%0A%0A";

    produtosParaPedido.forEach((produto) => {
      const quantidadePedido =
        Math.max(0, 300 - produto.quantidade);

      mensagem +=
        `${produto.marca} ${produto.modelo}: ` +
        `${quantidadePedido} unidades%0A`;
    });

    mensagem +=
      "%0AGostaria de confirmar a disponibilidade.";

    return mensagem;
  }

  function enviarWhatsApp() {
    const mensagem = gerarMensagem();

    const url = `https://wa.me/?text=${mensagem}`;

    window.open(url, "_blank");
  }

  return (
    <div>
      <div className="page-title">
        <div>
          <h2>Pedidos</h2>
          <p>
            Produtos que precisam de reposição.
          </p>
        </div>

        <button
          className="button-primary"
          onClick={enviarWhatsApp}
        >
          Enviar pedido pelo WhatsApp
        </button>
      </div>

      <div className="order-info">
        <strong>
          {produtosParaPedido.length}
        </strong>

        <span>
          produtos precisam de reposição
        </span>
      </div>

      <div className="table-container">
        <table>
          <thead>
            <tr>
              <th>Marca</th>
              <th>Modelo</th>
              <th>Estoque atual</th>
              <th>Quantidade sugerida</th>
            </tr>
          </thead>

          <tbody>
            {produtosParaPedido.map((produto) => {
              const quantidadePedido =
                Math.max(
                  0,
                  300 - produto.quantidade
                );

              return (
                <tr key={produto.id}>
                  <td>{produto.marca}</td>

                  <td>
                    <strong>
                      {produto.modelo}
                    </strong>
                  </td>

                  <td>
                    <span className="status critical">
                      {produto.quantidade}
                    </span>
                  </td>

                  <td>
                    {quantidadePedido}
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </div>
  );
}