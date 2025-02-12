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
import { AxiosError } from "axios";

const createPostSchema = object({
  title: string()
    .required()
    .min(5)
    .max(50)
    .matches(
      /^[^{}&*%$#+=<>^\\|/]*$/,
      "Some special characters are not allowed"
    )
    .label("Title"),
  body: string()
    .required()
    .matches(
      /^[^{}&*%$#+=<>^\\|/]*$/,
      "Some special characters are not allowed"
    )
    .label("Post content"),
});

export const NewPostForm: FC<{ onClose: () => void }> = ({ onClose }) => {
  const searchParams = useSearchParams();
  const userId = searchParams.get("userId");

  const queryClient = useQueryClient();
  const { mutate, isPending } = useMutation({
    mutationFn: postService.createPost,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["posts", userId],
      });
    },
    onError: (err: Error | AxiosError) => {
      if (err instanceof AxiosError) {
        alert(err.response?.data.message || "Failed to create post");
      } else {
        alert(err.message || "Failed to create post");
      }
    },
  });

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(createPostSchema),
    defaultValues: { title: "", body: "" },
  });

  return (
    <form
      className="flex flex-col z-50 p-6 gap-6 md:w-[632px] sm:w-full mobilesm:w-full bg-white border rounded-lg shadow-sm"
      onClick={(e) => e.stopPropagation()}
      onSubmit={handleSubmit(async (data) => {
        mutate({ ...data, userId: Number(userId!) });
        onClose();
      })}
    >
      <h1 className="text-x font-medium text-black">New Post</h1>
      <Input
        label="Post title"
        {...register("title")}
        errorText={errors.title?.message}
      />
      <TextArea
        label="Post content"
        {...register("body")}
        errorText={errors.body?.message}
      />
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
