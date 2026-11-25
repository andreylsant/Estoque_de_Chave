export function mostrarMensagem(
  setMensagem: (mensagem: string) => void,
  texto: string
) {
  setMensagem(texto);

  setTimeout(() => {
    setMensagem("");
  }, 3000);
}