import "./styles.scss";

interface HeaderProps {
  name: string;
  key: string;
}

interface TableProps {
  tableHeaders: HeaderProps[];
  data?: any;
}

export default function Table({ tableHeaders, data }: TableProps) {
  return (
    <div className="Table-container">
      <div className="Table">
        <table cellSpacing={0}>
          <thead>
            <tr>
              {tableHeaders.map((k, i) => (
                <th scope="col" key={i}>
                  {k.name}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {data.map((v, i) => (
              <tr key={i}>
                {tableHeaders.map((k, i2) => (
                  <td>{v[k.key]}</td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
