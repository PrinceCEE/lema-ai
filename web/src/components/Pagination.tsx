import { FC, Fragment } from "react";
import { NextLabel, PreviousLabel } from "./Button";
import { VoidCallback } from "@/types";
import clsx from "clsx";

interface PaginatorProps {
  currentIndex: number;
  lastIndex: number;
  isLeftButtonEnabled: boolean;
  isRightButtonEnabled: boolean;
  onPaginatorItemClick(val: number): void;
  onLeftButtonClick: VoidCallback;
  onRightButtonClick: VoidCallback;
}

export const Paginator: FC<PaginatorProps> = (props) => {
  return (
    <div className="flex items-center self-end md:gap-x-[42px]">
      <PreviousLabel
        isDisable={!props.isLeftButtonEnabled}
        onClick={props.onLeftButtonClick}
      />

      <div className="flex items-center">
        {paginator(props.currentIndex, props.lastIndex).map((item, index) => (
          <Fragment key={index}>
            {item == -1 ? (
              <Ellipse />
            ) : (
              <PaginatorItem
                props={{
                  value: item,
                  isActive: item == props.currentIndex,
                  onPress: props.onPaginatorItemClick,
                }}
              />
            )}
          </Fragment>
        ))}
      </div>

      <NextLabel
        isDisble={!props.isRightButtonEnabled}
        onClick={props.onRightButtonClick}
      />
    </div>
  );
};

function PaginatorItem({
  props,
}: {
  props: { isActive: boolean; value: number; onPress: (val: number) => void };
}) {
  return (
    <button
      onClick={() => {
        window.scrollTo(0, 0);
        document.getElementById("wrapper")?.scrollTo(0, 0);
        props.onPress(props.value);
      }}
      className={clsx(`w-10 rounded-lg h-10 flex items-center justify-center text-sm font-medium text-lightblack hover:text-paginationBtnHoverText hover:bg-paginationBtnBg cursor-pointer 
        ${
          props.isActive ? "text-paginationBtnHoverText bg-paginationBtnBg" : ""
        }
        `)}
    >
      {props.value}
    </button>
  );
}

function Ellipse() {
  return <div>...</div>;
}

function paginator(currentIndex: number, length: number): number[] {
  if (currentIndex < 1 || currentIndex > length) return [];
  const pick = 5;
  if (length - 1 < pick + 2)
    return new Array(length).fill("").map((_, index) => index + 1);
  if (currentIndex < pick)
    return [
      ...Array(pick)
        .fill("")
        .map((_, index) => index + 1),
      -1,
      length,
    ];
  if (currentIndex + pick > length + 1)
    return [
      1,
      -1,
      ...Array(pick)
        .fill("")
        .map((_, index) => length + 1 - (pick - index)),
    ];
  return [
    1,
    -1,
    ...new Array(pick - 2).fill("").map((_, index) => currentIndex - 1 + index),
    -1,
    length,
  ];
}
