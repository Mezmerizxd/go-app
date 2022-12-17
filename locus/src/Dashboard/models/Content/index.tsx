import "./styles.scss";

interface ContentProps {
  title: string;
  children?: any;
}

export default function Content({ title, children }: ContentProps): any {
  return (
    <div className="Content-container">
      <div className="Content-header">
        <h1>{title}</h1>
      </div>
      <div className="Content-body">{children}</div>
    </div>
  );
}
