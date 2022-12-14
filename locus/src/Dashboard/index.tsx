import MenuIcon from "@mui/icons-material/Menu";
import DashboardIcon from "@mui/icons-material/Dashboard";
import AccessibilityIcon from "@mui/icons-material/Accessibility";
import Person2Icon from "@mui/icons-material/Person2";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import FormatPaintIcon from "@mui/icons-material/FormatPaint";
import PeopleIcon from "@mui/icons-material/People";
import LocalPoliceIcon from "@mui/icons-material/LocalPolice";
import SideNav from "../models/SideNav";
import { useState } from "react";
import "./styles.scss";

export default function Dashboard() {
  const [isSideNavOpen, setIsSideNavOpen] = useState(true);

  return (
    <div className="Dashboard-container">
      <div className="Dashboard-top-nav-container">
        <div className="top-nav-left">
          <i>
            <MenuIcon onClick={() => setIsSideNavOpen(!isSideNavOpen)} />
          </i>
        </div>
      </div>

      <div
        className="Dashboard"
        id={isSideNavOpen ? "ShowSideNav" : "HideSideNav"}
      >
        <div className="Dashboard-side-nav-container">
          <SideNav
            isSideNavOpen={isSideNavOpen}
            groups={[
              {
                groupName: "Main",
                options: [
                  {
                    name: "Dashboard",
                    icon: <DashboardIcon />,
                  },
                  {
                    name: "Friends",
                    icon: <PeopleIcon />,
                  },
                ],
              },
              {
                groupName: "Settings",
                options: [
                  {
                    name: "Profile",
                    icon: <Person2Icon />,
                  },
                  {
                    name: "Account",
                    icon: <AccountBoxIcon />,
                  },
                  {
                    name: "Appearance",
                    icon: <FormatPaintIcon />,
                  },
                ],
              },
              {
                groupName: "Other",
                options: [
                  {
                    name: "Accessibility",
                    icon: <AccessibilityIcon />,
                  },
                  {
                    name: "Legalities",
                    icon: <LocalPoliceIcon />,
                  },
                ],
              },
            ]}
          />
        </div>

        <div className="Dashboard-content">
          <p>Dashboard</p>
        </div>
      </div>
    </div>
  );
}
