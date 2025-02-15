import clsx from "clsx";
import { FC } from "react";

export const Loader: FC<{ bg?: string; isSmall?: boolean }> = ({
  bg,
  isSmall,
}) => {
  return (
    <div className="lds-ellipsis">
      <div
        className={clsx(
          `${
            isSmall ? "w-[4.18px] h-[4.18px]" : "w-[11.6633px] h-[11.6633px]"
          } ${bg ?? "bg-[#7f56d9]"}`
        )}
      ></div>
      <div
        className={clsx(
          `${
            isSmall ? "w-[5.85px] h-[5.85px]" : "w-[13.3333px] h-[13.3333px]"
          } ${bg ?? "bg-[#7f56d9]"}`
        )}
      ></div>
      <div
        className={clsx(
          `${
            isSmall ? "w-[5.85px] h-[5.85px]" : "w-[13.3333px] h-[13.3333px]"
          } ${bg ?? "bg-[#7f56d9]"}`
        )}
      ></div>
    </div>
  );
};
