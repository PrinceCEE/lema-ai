import { FC } from "react";
import { FaArrowLeft, FaArrowRight } from "react-icons/fa";

export const PreviousLabel = () => {
  return (
    <div className="w-max flex items-center justify-center gap-x-2">
      <FaArrowLeft className="h-[11.67px] w-[11.67px] text-tableArrow" />
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Previous
      </span>
    </div>
  );
};

export const NextLabel = () => {
  return (
    <div className="w-max flex items-center justify-center gap-x-2">
      <span className="text-sm font-semibold text-lightblack hidden md:inline">
        Next
      </span>
      <FaArrowRight className="h-[11.67px] w-[11.67px] text-tableArrow" />
    </div>
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
