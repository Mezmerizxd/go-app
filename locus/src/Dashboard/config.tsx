import DashboardIcon from "@mui/icons-material/Dashboard";
import AccessibilityIcon from "@mui/icons-material/Accessibility";
import Person2Icon from "@mui/icons-material/Person2";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import FormatPaintIcon from "@mui/icons-material/FormatPaint";
import LocalPoliceIcon from "@mui/icons-material/LocalPolice";
import { SideNavGroupProps } from "../types/dashboard";
import { FaMask, FaGlasses } from "react-icons/fa";
import {
  GiHieroglyphLegs,
  GiConverseShoe,
  GiNecklace,
  GiTShirt,
  GiDropEarrings,
  GiWatch,
  GiGemNecklace,
} from "react-icons/gi";
import { SiRedhat } from "react-icons/si";

export enum Pages {
  Dashboard = "dashboard",
  GTA5C_Masks = "gta5c_masks",
  GTA5C_Legs = "gta5c_legs",
  GTA5C_Shoes = "gta5c_shoes",
  GTA5C_Accessories = "gta5c_accessories",
  GTA5C_Tops = "gta5c_tops",
  GTA5C_Hats = "gta5c_hats",
  GTA5C_Glasses = "gta5c_glasses",
  GTA5C_Earrings = "gta5c_earrings",
  GTA5C_Watches = "gta5c_watches",
  GTA5C_Bracelets = "gta5c_bracelets",
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
    ],
  },
  {
    groupName: "GTA 5 Clothing",
    options: [
      {
        name: "Masks",
        icon: <FaMask />,
        page: Pages.GTA5C_Masks,
      },
      {
        name: "Legs",
        icon: <GiHieroglyphLegs />,
        page: Pages.GTA5C_Legs,
      },
      {
        name: "Shoes",
        icon: <GiConverseShoe />,
        page: Pages.GTA5C_Shoes,
      },
      {
        name: "Accessories",
        icon: <GiNecklace />,
        page: Pages.GTA5C_Accessories,
      },
      {
        name: "Tops",
        icon: <GiTShirt />,
        page: Pages.GTA5C_Tops,
      },
      {
        name: "Hats",
        icon: <SiRedhat />,
        page: Pages.GTA5C_Hats,
      },
      {
        name: "Glasses",
        icon: <FaGlasses />,
        page: Pages.GTA5C_Glasses,
      },
      {
        name: "Earrings",
        icon: <GiDropEarrings />,
        page: Pages.GTA5C_Earrings,
      },
      {
        name: "Watches",
        icon: <GiWatch />,
        page: Pages.GTA5C_Watches,
      },
      {
        name: "Bracelets",
        icon: <GiGemNecklace />,
        page: Pages.GTA5C_Bracelets,
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
