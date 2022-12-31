import { DataGrid } from "@mui/x-data-grid";
import "./styles.scss";

interface HeaderProps {
  field: string;
  name: string;
  width: number;
}

interface TableProps {
  tableHeaders: HeaderProps[];
  data?: any;
}

export default function DataTable({ tableHeaders, data }: TableProps) {
  return (
    <div className="DataTable-container">
      <div className="DataTable">
        <DataGrid
          density="standard"
          rows={data ? data : []}
          columns={
            tableHeaders
              ? tableHeaders.map((header) => {
                  return {
                    field: header.field,
                    headerName: header.name,
                    width: header.width,
                  };
                })
              : []
          }
          headerHeight={60}
          rowHeight={57}
          rowsPerPageOptions={[10, 25, 50, 100]}
          pageSize={10}
          checkboxSelection
          sx={{
            border: "0",
            "& .MuiDataGrid-root": {
              backgroundColor: "rgb(40, 40, 40)",
            },
            // create light theme version
            "& .MuiDataGrid-columnsContainer": {
              backgroundColor: "rgb(40, 40, 40)",
            },
            // create light theme header
            "& .MuiDataGrid-columnHeaders": {
              backgroundColor: "rgb(40, 40, 40)",
              color: "rgba(69, 255, 178, 0.8)",
            },
            "& .MuiDataGrid-columnHeaderTitle": {
              fontWeight: "bold",
            },
            // create light theme rows
            "& .MuiDataGrid-cell": {
              backgroundColor: "rgb(30, 30, 30)",
              color: "rgb(174, 199, 188)",
            },
            // footer
            "& .MuiDataGrid-footerContainer": {
              backgroundColor: "rgb(40, 40, 40)",
              color: "rgb(174, 199, 188)",
              border: "0",
            },
            "& .MuiToolbar-root": {
              color: "rgba(69, 255, 178, 0.8)",
            },
            "& .MuiSvgIcon-root": {
              color: "rgba(69, 255, 178, 0.8)",
            },
          }}
        />
      </div>
    </div>
  );
}
