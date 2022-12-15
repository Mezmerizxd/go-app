import MenuIcon from "@mui/icons-material/Menu";
import SideNav from "./models/SideNav";
import "./styles.scss";
import { useAppSelector, useAppDispatch } from "../hooks/reduxHooks";
import { toggleSideNav } from "./reducer";
import { SidebarOptions } from "./config";

export default function Dashboard() {
  const state = useAppSelector((state) => state.dashboard);
  const dispatch = useAppDispatch();

  return (
    <div className="Dashboard-container">
      <div className="Dashboard-top-nav-container">
        <div className="top-nav-left">
          <i>
            <MenuIcon onClick={() => dispatch(toggleSideNav())} />
          </i>
          <div className="top-nav-left-title">
            <h1>LOCUS</h1>
          </div>
        </div>
      </div>

      <div
        className="Dashboard"
        id={state.isSideNavOpen ? "ShowSideNav" : "HideSideNav"}
      >
        <div className="Dashboard-side-nav-container">
          <SideNav groups={SidebarOptions} />
        </div>

        <div className="Dashboard-content">
          <p>Dashboard</p>
        </div>
      </div>
    </div>
  );
}
