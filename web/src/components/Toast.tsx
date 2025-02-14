"use client";

import { useAppStore, useAppStoreDispatch } from "@/hooks";
import { IoIosClose } from "react-icons/io";
import { AppStore } from "@/types";
import clsx from "clsx";
import { FC, useEffect } from "react";

const Toast: FC<{
  notification: AppStore["notifications"][number];
  index: number;
}> = ({ notification, index }) => {
  const dispatch = useAppStoreDispatch();

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      dispatch({
        type: "REMOVE_NOTIFICATION",
        payload: { index },
      });
    }, 3000);

    return () => clearTimeout(timeoutId);
  }, [index]);

  const handleClick = () => {
    dispatch({
      type: "REMOVE_NOTIFICATION",
      payload: { index },
    });
  };

  return (
    <div
      className={clsx(
        `flex justify-between items-start w-[300px] p-2 rounded-md text-white ${
          notification.isSuccess ? "bg-green-400" : "bg-red-400"
        }`
      )}
    >
      <span>{notification.text}</span>
      <IoIosClose
        className="font-bold text-x cursor-pointer"
        onClick={handleClick}
      />
    </div>
  );
};

export const ToastContainer = () => {
  const store = useAppStore();

  if (!store.notifications.length) {
    return;
  }

  return (
    <div className="absolute z-[1000] top-0 right-0 w-max flex flex-col gap-1 p-4">
      {store.notifications.map((n, i) => (
        <Toast key={i} notification={n} index={i} />
      ))}
    </div>
  );
};
