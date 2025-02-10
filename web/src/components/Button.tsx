import clsx from "clsx";
import { FC } from "react";

export const Button: FC<{
  text: string;
  className: string;
  isSubmit?: boolean;
  isDisabled?: boolean;
  onClick?: () => void;
}> = ({ text, onClick, className, isSubmit, isDisabled }) => {
  return (
    <button
      className={clsx(
        `${className} flex justify-center items-center px-4 py-[11.5px] border rounded-[4px]`
      )}
      onClick={onClick}
      type={isSubmit ? "submit" : "button"}
      disabled={isDisabled}
    >
      {text}
    </button>
  );
};
