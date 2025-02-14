import clsx from "clsx";
import { FC, ReactNode } from "react";
import { FaArrowLeft, FaArrowRight } from "react-icons/fa";

export const Button: FC<{
  children: ReactNode;
  className: string;
  isSubmit?: boolean;
  isDisabled?: boolean;
  onClick?: () => void;
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

export const BackButton: FC<{ text: string; handleClick: () => void }> = ({
  text,
  handleClick,
}) => {
  return (
    <div
      className="w-max flex items-center justify-center gap-x-2 cursor-pointer"
      onClick={() => handleClick()}
    >
      <FaArrowLeft className="h-[11.67px] w-[11.67px] text-tableArrow" />
      <span className="text-sm font-semibold text-lightblack">{text}</span>
    </div>
  );
};

export const PreviousLabel = () => {
  return (
    <Button className="w-max flex items-center justify-center gap-x-2">
      <FaArrowLeft className="h-[11.67px] w-[11.67px] text-tableArrow" />
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Previous
      </span>
    </Button>
  );
};

export const NextLabel = () => {
  return (
    <Button className="w-max flex items-center justify-center gap-x-2">
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Next
      </span>
      <FaArrowRight className="h-[11.67px] w-[11.67px] text-tableArrow" />
    </Button>
  );
};
