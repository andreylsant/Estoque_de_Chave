import { useEffect, useMemo, useState } from "react";
import { StockCard } from "../componets/StockCard";
import { buscarProdutos } from "../services/produtoService";
import type { Produto } from "../types/produto";

export function Dashboard(){
 const [produtos,setProdutos]=useState<Produto[]>([]);
 useEffect(()=>{buscarProdutos().then(setProdutos).catch(err=>console.error("Erro ao buscar produtos:",err));},[]);
 const total=useMemo(()=>produtos.reduce((s,p)=>s+p.quantidade,0),[produtos]);
 const criticos=produtos.filter(p=>p.quantidade<=100);
 const baixos=produtos.filter(p=>p.quantidade>100&&p.quantidade<=200);
 const normal=produtos.filter(p=>p.quantidade>200);
 const format=(n:number)=>new Intl.NumberFormat("pt-BR").format(n);
 const bars=[{m:"Jul",v:Math.round(total*.54)},{m:"Ago",v:Math.round(total*.62)},{m:"Set",v:Math.round(total*.69)},{m:"Out",v:Math.round(total*.81)},{m:"Nov",v:Math.round(total*.9)},{m:"Dez",v:total}];
 const max=Math.max(...bars.map(b=>b.v),1);
 return <div className="dashboard">
  <div className="page-title"><div><div className="eyebrow">MASTER DASHBOARD&nbsp; • &nbsp;VISÃO GERAL</div><h2>Painel de Controle</h2><p>Bem-vindo de volta. Aqui está o resumo da operação do seu estoque.</p></div><div className="page-actions"><button className="button-secondary" onClick={()=>window.print()}>▣ &nbsp; Exportar</button></div></div>
  <div className="cards-grid">
   <StockCard title="Modelos cadastrados" value={produtos.length} description="Itens no catálogo" icon="▣"/>
   <StockCard title="Total em estoque" value={format(total)} description="Chaves disponíveis" icon="⌑"/>
   <StockCard title="Estoque crítico" value={criticos.length} description="Até 100 unidades" icon="⚠" accent="amber"/>
   <StockCard title="Estoque normal" value={normal.length} description="Acima de 200 unidades" icon="↗" accent="green"/>
  </div>
  <div className="dashboard-tabs"><span>Visão Geral</span><span>Estoque</span><span>Indicadores</span></div>
  <div className="dashboard-grid">
   <section className="panel"><div className="panel-heading"><div><h3>▥ &nbsp;Volume de estoque</h3><p>Visão ilustrativa da evolução do estoque</p></div><span style={{color:"#85859e",fontSize:10}}>ÚLTIMOS 6 MESES</span></div>
    <div className="chart-wrap"><svg viewBox="0 0 680 280" role="img" aria-label="Gráfico ilustrativo do volume de estoque"><defs><linearGradient id="barFill" x1="0" x2="0" y1="0" y2="1"><stop offset="0%" stopColor="#7772e5"/><stop offset="100%" stopColor="#4039a5"/></linearGradient></defs>{[35,85,135,185,235].map(y=><g key={y}><line x1="48" y1={y} x2="665" y2={y} stroke="#29293b" strokeDasharray="3 5"/><text x="38" y={y+4} fill="#77778e" fontSize="10" textAnchor="end">{format(Math.round(max*(280-y)/245))}</text></g>)}{bars.map((b,i)=>{const x=85+i*98;const h=b.v/max*190;return <g key={b.m}><rect x={x} y={235-h} width="34" height={h} rx="5" fill="url(#barFill)"/><circle cx={x+17} cy={235-h} r="4" fill="#24c99a" stroke="#133f37" strokeWidth="3"/><text x={x+17} y="258" fill="#85859e" fontSize="10" textAnchor="middle">{b.m}</text></g>})}<polyline points={bars.map((b,i)=>`${102+i*98},${235-(b.v/max*190)}`).join(" ")} fill="none" stroke="#16b88b" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round"/></svg></div>
   </section>
   <section className="panel"><div className="panel-heading"><div><h3>⌁ &nbsp;Atenção no estoque</h3><p>Itens que precisam de reposição</p></div></div>
    {criticos.slice(0,3).map(p=><div className="activity-item" key={p.id}><div className="activity-badge" style={{background:"#3b1723",color:"#ff8295"}}>!</div><div><strong>{p.marca} {p.modelo}</strong><small>Estoque crítico</small></div><div className="activity-meta">{p.quantidade} un.<small>Repor agora</small></div></div>)}
    {baixos.slice(0,2).map(p=><div className="activity-item" key={p.id}><div className="activity-badge" style={{background:"#3a2c12",color:"#f5bd56"}}>!</div><div><strong>{p.marca} {p.modelo}</strong><small>Estoque baixo</small></div><div className="activity-meta">{p.quantidade} un.<small>Acompanhar</small></div></div>)}
    {criticos.length+baixos.length===0&&<div className="empty-state">Tudo certo! Nenhum item precisa de reposição.</div>}
    <div className="panel-footer">{criticos.length+baixos.length} item(ns) precisam de atenção</div>
   </section>
  </div>
  <section className="dashboard-section"><div className="section-header"><h3>Distribuição por situação</h3><p>Classificação automática conforme a quantidade disponível</p></div><div className="stock-status-grid"><div className="status-card normal"><div className="status-icon">●</div><div><span>Estoque normal</span><strong>{normal.length}</strong><small>modelos acima de 200</small></div></div><div className="status-card low"><div className="status-icon">●</div><div><span>Estoque baixo</span><strong>{baixos.length}</strong><small>de 101 a 200 unidades</small></div></div><div className="status-card critical"><div className="status-icon">●</div><div><span>Estoque crítico</span><strong>{criticos.length}</strong><small>até 100 unidades</small></div></div></div></section>
 </div>;
}
