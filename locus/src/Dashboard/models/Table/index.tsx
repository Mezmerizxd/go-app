import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import KeyboardArrowRightIcon from "@mui/icons-material/KeyboardArrowRight";
import { useState } from "react";
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
  const [isOpen, setIsOpen] = useState<boolean>(true);
  return (
    <div className="Table-container">
      <div className="Table">
        <table cellSpacing={0}>
          <thead onClick={() => setIsOpen(!isOpen)}>
            <tr>
              {tableHeaders.map((k: any, i: number) => (
                <th scope="col" key={i}>
                  {k.name}
                </th>
              ))}
              <td onClick={() => setIsOpen(!isOpen)}>
                {isOpen ? (
                  <i>
                    <KeyboardArrowDownIcon />
                  </i>
                ) : (
                  <i>
                    <KeyboardArrowRightIcon />
                  </i>
                )}
              </td>
            </tr>
          </thead>
          {isOpen && (
            <tbody>
              {data.map((v: any, i: number) => (
                <tr key={i}>
                  {tableHeaders.map((k: any, i2: number) => (
                    <td key={i + i2}>{v[k.key]}</td>
                  ))}
                </tr>
              ))}
            </tbody>
          )}
        </table>
      </div>
    </div>
  );
}
