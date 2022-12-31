import MenuIcon from "@mui/icons-material/Menu";
import SideNav from "./components/SideNav/SideNav";
import { useAppSelector, useAppDispatch } from "../hooks/reduxHooks";
import { toggleSideNav } from "./reducer";
import { Pages, SidebarOptions } from "./config";
import Masks from "./components/GTA/Masks";
import Legs from "./components/GTA/Legs";
import "./styles.scss";

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
        <SideNav groups={SidebarOptions} />

        <div className="Dashboard-content">
          {state.page === Pages.Dashboard && <></>}

          {state.page === Pages.GTA5C_Masks && <Masks />}
          {state.page === Pages.GTA5C_Legs && <Legs />}
          {state.page === Pages.GTA5C_Shoes && <></>}
          {state.page === Pages.GTA5C_Accessories && <></>}
          {state.page === Pages.GTA5C_Tops && <></>}
          {state.page === Pages.GTA5C_Hats && <></>}
          {state.page === Pages.GTA5C_Glasses && <></>}
          {state.page === Pages.GTA5C_Earrings && <></>}
          {state.page === Pages.GTA5C_Watches && <></>}
          {state.page === Pages.GTA5C_Bracelets && <></>}

          {state.page === Pages.Profile && <></>}
          {state.page === Pages.Account && <></>}
          {state.page === Pages.Appearance && <></>}

          {state.page === Pages.Accessibility && <></>}
          {state.page === Pages.Legalities && <></>}
        </div>
      </div>
    </div>
  );
}
