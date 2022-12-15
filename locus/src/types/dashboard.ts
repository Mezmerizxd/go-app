import { Pages } from "../Dashboard/config";

export interface SideNavButtonProps {
  name: string;
  icon: React.ReactNode;
  page: Pages;
}

export interface SideNavGroupProps {
  groupName: string;
  options: SideNavButtonProps[];
}

export interface SideNavProps {
  groups: SideNavGroupProps[];
}

export interface DashboardSlice {
  isSideNavOpen: boolean;
  page: Pages;
}
