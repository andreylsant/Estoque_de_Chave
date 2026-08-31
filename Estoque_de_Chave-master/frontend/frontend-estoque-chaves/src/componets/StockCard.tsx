interface StockCardProps {
  title: string;
  value: number;
  description: string;
  icon: string;
}

export function StockCard({
  title,
  value,
  description,
  icon,
}: StockCardProps) {
  return (
    <div className="stock-card">
      <div className="stock-card-top">
        <span>{title}</span>

        <div className="stock-icon">
          {icon}
        </div>
      </div>

      <strong>{value}</strong>

      <p>{description}</p>
    </div>
  );
}