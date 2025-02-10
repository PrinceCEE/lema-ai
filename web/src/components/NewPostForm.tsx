"use client";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { object, string } from "yup";
import { Input, TextArea } from "./Input";
import { Button } from "./Button";
import { FC } from "react";
import { useSearchParams } from "next/navigation";
import { postService } from "@/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";

const createPostSchema = object({
  title: string().required(),
  body: string().required(),
});

export const NewPostForm: FC<{ onClose: () => void }> = ({ onClose }) => {
  const searchParams = useSearchParams();
  const userId = searchParams.get("userId");

  const queryClient = useQueryClient();
  const { mutate, isPending, isError } = useMutation({
    mutationFn: postService.createPost,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["posts", userId],
      });
    },
  });

  if (isError) {
    alert("Failed to create post");
  }

  const { register, handleSubmit } = useForm({
    resolver: yupResolver(createPostSchema),
  });

  return (
    <form
      className="flex flex-col z-50 p-6 gap-6 bg-white border rounded-lg shadow-sm"
      onClick={(e) => e.stopPropagation()}
      onSubmit={handleSubmit(async (data) => {
        mutate({ ...data, userId: Number(userId!) });
        onClose();
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
          isDisabled={isPending}
        />
      </div>
    </form>
  );
};
