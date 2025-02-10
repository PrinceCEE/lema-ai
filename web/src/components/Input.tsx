import { FC } from "react";

interface InputProps
  extends React.DetailedHTMLProps<
    React.InputHTMLAttributes<HTMLInputElement>,
    HTMLInputElement
  > {
  label: string;
}

interface TextAreaProps
  extends React.DetailedHTMLProps<
    React.TextareaHTMLAttributes<HTMLTextAreaElement>,
    HTMLTextAreaElement
  > {
  label: string;
}

export const Input: FC<InputProps> = ({ label, ...props }) => {
  return (
    <div className="flex flex-col gap-y-[10px] w-[631px]">
      <span className="block font-medium text-[18px] leading-[20px] text-lightblack">
        {label}
      </span>
      <div className="w-full border rounded-[4px] px-4 py-[10px]">
        <input
          className="w-full placeholder-[#94A3B8] text-sm font-normal bg-[transparent] outline-none focus:outline-none border-none focus:border-none"
          placeholder="Give your post a title"
          {...props}
        />
      </div>
    </div>
  );
};

export const TextArea: FC<TextAreaProps> = ({ label, ...props }) => {
  return (
    <div className="flex flex-col gap-y-[10px] w-[631px]">
      <span className="block font-medium text-[18px] leading-[20px] text-lightblack">
        {label}
      </span>
      <div className="w-full border rounded-[4px] px-4 py-[10px]">
        <textarea
          className="w-full placeholder-[#94A3B8] text-sm font-normal bg-[transparent] outline-none focus:outline-none border-none focus:border-none resize-none"
          placeholder="Write something mind-blowing"
          rows={5}
          {...props}
        />
      </div>
    </div>
  );
};
