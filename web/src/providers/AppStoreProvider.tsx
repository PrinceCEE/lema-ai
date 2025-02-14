"use client";
import { appStoreReducer, initialState } from "@/reducer";
import { AppStore, PayloadAction } from "@/types";
import { createContext, Dispatch, FC, ReactNode, useCallback } from "react";
import { useImmerReducer } from "use-immer";

export const AppStoreContext = createContext(null as unknown as AppStore);
export const AppStoreDispatchContext = createContext(
  null as unknown as Dispatch<PayloadAction>
);

export const AppStoreProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [store, dispatch] = useImmerReducer(appStoreReducer, initialState);

  const fn = useCallback(dispatch, []);

  return (
    <AppStoreContext.Provider value={store}>
      <AppStoreDispatchContext.Provider value={fn}>
        {children}
      </AppStoreDispatchContext.Provider>
    </AppStoreContext.Provider>
  );
};
