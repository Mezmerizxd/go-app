import Content from "../../../models/Content";
import Table from "../../../models/Table";
import { Data } from "./data";

export default function Masks() {
  return (
    <Content title="Masks">
      <Table
        tableHeaders={[
          {
            name: "Male & Female",
            key: "item",
          },
          {
            name: "ID",
            key: "id",
          },
          {
            name: "Texture",
            key: "texture",
          },
        ]}
        data={Data}
      />
    </Content>
  );
}
