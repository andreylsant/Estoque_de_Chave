import { api } from "./api";
import type { Produto } from "../types/produto";

export async function buscarProdutos(): Promise<Produto[]> {
  const response = await api.get("/produtos");

  return response.data;
}