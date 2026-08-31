export function Header(){
    return(
        <header className="header">
      <div>
        <h1>Controle de Estoque</h1>
        <p>Gerencie o estoque de chaves da loja</p>
      </div>

      <div className="user-info">
        <div className="avatar">
          A
        </div>

        <div>
          <strong>Administrador</strong>
          <span>Usuário</span>
        </div>
      </div>
    </header>
    );
}