import { AppStoreContext, AppStoreDispatchContext } from "@/providers";
import { useContext } from "react";

export const useAppStore = () => {
  return useContext(AppStoreContext);
};

export const useAppStoreDispatch = () => {
  return useContext(AppStoreDispatchContext);
};
