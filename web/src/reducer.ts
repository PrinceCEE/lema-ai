import { AppStore, PayloadAction } from "./types";

export const initialState: AppStore = {
  notifications: [],
};

export const appStoreReducer = (state: AppStore, action: PayloadAction) => {
  switch (action.type) {
    case "ADD_NOTIFICATION": {
      state.notifications.push({
        text: action.payload.text,
        isSuccess: action.payload.isSuccess,
      });
      break;
    }
    case "REMOVE_NOTIFICATION": {
      state.notifications.splice(action.payload.index, 1);
      break;
    }

    default:
      throw new Error("Unknown action");
  }
};
