import clsx from "clsx";
import { FC } from "react";

export const Loader: FC<{ bg?: string }> = ({ bg }) => {
  return (
    <div className="lds-ellipsis">
      <div className={clsx(`w-[6px] h-[6px] ${bg ?? "bg-[#7f56d9]"}`)}></div>
      <div className={clsx(`w-[6px] h-[6px] ${bg ?? "bg-[#7f56d9]"}`)}></div>
      <div className={clsx(`w-[6px] h-[6px] ${bg ?? "bg-[#7f56d9]"}`)}></div>
      <div className={clsx(`w-[6px] h-[6px] ${bg ?? "bg-[#7f56d9]"}`)}></div>
    </div>
  );
};
