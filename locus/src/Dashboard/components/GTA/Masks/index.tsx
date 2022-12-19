import Content from "../../../models/Content";
import Table from "../../../models/Table";
import { Data } from "./data2";

export default function Masks() {
  return (
    <Content title="Masks">
      {Data.map((v) => (
        <Table
          tableHeaders={[
            {
              name: v.class_name,
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
          data={v.class_items}
        />
      ))}
    </Content>
  );
}
