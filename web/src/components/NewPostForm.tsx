"use client";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { object, string } from "yup";
import { Input, TextArea } from "./Input";
import { Button } from "./Button";
import { FC, useState } from "react";
import { useSearchParams } from "next/navigation";
import { postService } from "@/api";

const createPostSchema = object({
  title: string().required(),
  body: string().required(),
});

export const NewPostForm: FC<{ onClose: () => void }> = ({ onClose }) => {
  const searchParams = useSearchParams();
  const [isLoading, setIsLoading] = useState(false);
  const userId = searchParams.get("userId");

  const { register, handleSubmit } = useForm({
    resolver: yupResolver(createPostSchema),
  });

  return (
    <form
      className="flex flex-col z-50 p-6 gap-6 bg-white border rounded-lg shadow-sm"
      onClick={(e) => e.stopPropagation()}
      onSubmit={handleSubmit(async (data) => {
        try {
          setIsLoading(true);
          await postService.createPost({
            ...data,
            userId: Number(userId!),
          });

          setIsLoading(false);
          onClose();
          window.location.reload();
        } catch (err: any) {
          alert(err.message);
        }
      })}
    >
      <h1 className="text-x font-medium text-black">New Post</h1>
      <Input label="Post title" {...register("title")} />
      <TextArea label="Post content" {...register("body")} />
      <div className="flex justify-end gap-2">
        <Button
          className="text-[14px] leading-[16.94px] font-normal text-[#334155] bg-white"
          text="Cancel"
          onClick={() => onClose()}
        />
        <Button
          className="bg-[#334155] font-semibold text-sm text-white"
          text="Publish"
          isSubmit
          isDisabled={isLoading}
        />
      </div>
    </form>
  );
};
