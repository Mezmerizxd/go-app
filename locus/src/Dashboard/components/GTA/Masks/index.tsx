import { useEffect, useState } from "react";
import API from "../../../../classes/API";
import Content from "../../../models/Content";
import Table from "../../../models/Table";

export default function Masks() {
  const [data, setData] = useState<any>(null);

  useEffect(() => {
    setTimeout(async () => {
      const r = await API.Get("http://localhost:3001/api/v1/locus/gta/masks");
      setData(r.data);
    });
  }, []);

  return (
    <Content title="Masks">
      {data?.map((v) => (
        <Table
          key={v?.class_name}
          tableHeaders={[
            {
              name: v?.class_name,
              key: "item",
            },
            {
              name: "ID",
              key: "item_id",
            },
            {
              name: "Texture",
              key: "texture_id",
            },
          ]}
          data={v?.class_items}
        />
      ))}
    </Content>
  );
}
