import { useEffect, useState } from "react";
import API from "../../../../classes/API";
import Content from "../../../models/Content";
import DataTable from "../../../models/DataTable";
import Table from "../../../models/Table";

export default function Legs() {
  const [data, setData] = useState<any>(null);

  useEffect(() => {
    setTimeout(async () => {
      const r = await API.Get("http://localhost:3001/api/v1/locus/gta/masks");
      setData(r.data);
    });
  }, []);

  const rows = [
    { id: 1, lastName: "Snow", firstName: "Jon", age: 35 },
    { id: 2, lastName: "Lannister", firstName: "Cersei", age: 42 },
    { id: 3, lastName: "Lannister", firstName: "Jaime", age: 45 },
    { id: 4, lastName: "Stark", firstName: "Arya", age: 16 },
    { id: 5, lastName: "Targaryen", firstName: "Daenerys", age: null },
  ];

  return (
    <Content title="Legs">
      {data?.map((v) => (
        <DataTable
          tableHeaders={[
            {
              field: "item",
              name: v?.class_name,
              width: 400,
            },
            {
              field: "item_id",
              name: "ID",
              width: 100,
            },
            {
              field: "texture_id",
              name: "Texture",
              width: 100,
            },
          ]}
          data={v?.class_items}
        />
      ))}
    </Content>
  );
}
