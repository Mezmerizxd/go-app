import ReactDOM from "react-dom/client";
import { Provider } from "react-redux";
import Dashboard from "./Dashboard";
import { store } from "./store";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <Provider store={store}>
    <Dashboard />
  </Provider>
);
