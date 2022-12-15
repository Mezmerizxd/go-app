import { configureStore } from "@reduxjs/toolkit";
import { dashboardSlice } from "./Dashboard/reducer";

export const store = configureStore({
  reducer: {
    dashboard: dashboardSlice.reducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
