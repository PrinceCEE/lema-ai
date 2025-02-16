import clsx from "clsx";
import { FC, ReactNode } from "react";
import { ArrowLeft, ArrowRight } from "./Arrows";
import { VoidCallback } from "@/types";

export const Button: FC<{
  children: ReactNode;
  className: string;
  isSubmit?: boolean;
  isDisabled?: boolean;
  onClick?: VoidCallback;
}> = ({ children, onClick, className, isSubmit, isDisabled }) => {
  return (
    <button
      className={clsx(`${className}`)}
      onClick={onClick}
      type={isSubmit ? "submit" : "button"}
      disabled={isDisabled}
    >
      {children}
    </button>
  );
};

export const BackButton: FC<{ text: string; handleClick: VoidCallback }> = ({
  text,
  handleClick,
}) => {
  return (
    <div
      className="w-max flex items-center justify-center gap-x-2 cursor-pointer"
      onClick={() => handleClick()}
    >
      <ArrowLeft />
      <span className="text-sm font-semibold text-lightblack">{text}</span>
    </div>
  );
};

export const PreviousLabel: FC<{
  isDisable: boolean;
  onClick: VoidCallback;
}> = ({ isDisable, onClick }) => {
  return (
    <Button
      onClick={onClick}
      className={clsx(
        `w-max flex items-center justify-center gap-x-2 ${
          isDisable ? "cursor-not-allowed" : "cursor-pointer"
        }`
      )}
      isDisabled={isDisable}
    >
      <ArrowLeft />
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Previous
      </span>
    </Button>
  );
};

export const NextLabel: FC<{ isDisble: boolean; onClick: VoidCallback }> = ({
  isDisble,
  onClick,
}) => {
  return (
    <Button
      onClick={onClick}
      className={clsx(
        `w-max flex items-center justify-center gap-x-2 ${
          isDisble ? "cursor-not-allowed" : "cursor-pointer"
        }`
      )}
    >
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Next
      </span>
      <ArrowRight />
    </Button>
  );
};
