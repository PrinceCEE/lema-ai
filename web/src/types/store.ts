export type AppStore = {
  notifications: {
    text: string;
    isSuccess: boolean;
  }[];
};

export type PayloadAction =
  | { type: "ADD_NOTIFICATION"; payload: { text: string; isSuccess: boolean } }
  | { type: "REMOVE_NOTIFICATION"; payload: { index: number } };
