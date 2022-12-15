import DashboardIcon from "@mui/icons-material/Dashboard";
import AccessibilityIcon from "@mui/icons-material/Accessibility";
import Person2Icon from "@mui/icons-material/Person2";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import FormatPaintIcon from "@mui/icons-material/FormatPaint";
import PeopleIcon from "@mui/icons-material/People";
import LocalPoliceIcon from "@mui/icons-material/LocalPolice";
import { SideNavGroupProps } from "../types/dashboard";
import { useAppSelector, useAppDispatch } from "../hooks/reduxHooks";

export enum Pages {
  Dashboard = "dashboard",
  Friends = "friends",
  Profile = "profile",
  Account = "account",
  Appearance = "appearance",
  Accessibility = "accessibility",
  Legalities = "legalities",
}

export const SidebarOptions: SideNavGroupProps[] = [
  {
    groupName: "Main",
    options: [
      {
        name: "Dashboard",
        icon: <DashboardIcon />,
        page: Pages.Dashboard,
      },
      {
        name: "Friends",
        icon: <PeopleIcon />,
        page: Pages.Friends,
      },
    ],
  },
  {
    groupName: "Settings",
    options: [
      {
        name: "Profile",
        icon: <Person2Icon />,
        page: Pages.Profile,
      },
      {
        name: "Account",
        icon: <AccountBoxIcon />,
        page: Pages.Account,
      },
      {
        name: "Appearance",
        icon: <FormatPaintIcon />,
        page: Pages.Appearance,
      },
    ],
  },
  {
    groupName: "Other",
    options: [
      {
        name: "Accessibility",
        icon: <AccessibilityIcon />,
        page: Pages.Accessibility,
      },
      {
        name: "Legalities",
        icon: <LocalPoliceIcon />,
        page: Pages.Legalities,
      },
    ],
  },
];
