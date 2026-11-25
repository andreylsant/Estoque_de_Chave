import { api } from "./api";
import type { Produto } from "../types/produto";

export async function buscarProdutos(): Promise<Produto[]> {
  const response = await api.get("/produtos");

  return response.data;
}
export async function movimentarEstoque(id: string, quantidade: number, tipo: "entrada" | "saida"): Promise<void> {
  await api.patch(`/produtos/${id}/estoque`, { quantidade, tipo });
}
