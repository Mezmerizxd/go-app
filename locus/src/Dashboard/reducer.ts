import { createSlice } from "@reduxjs/toolkit";
import { DashboardSlice } from "../types/dashboard";
import { Pages } from "./config";

export const dashboardInitialData: DashboardSlice = {
  isSideNavOpen: true,
  page: Pages.Dashboard,
};

export const dashboardSlice = createSlice({
  name: "dashboard",
  initialState: dashboardInitialData,
  reducers: {
    toggleSideNav: (state) => {
      state.isSideNavOpen = !state.isSideNavOpen;
    },
    setPage: (state, action) => {
      state.page = action.payload;
    },
  },
});

export const { toggleSideNav, setPage } = dashboardSlice.actions;

export default dashboardSlice.reducer;
