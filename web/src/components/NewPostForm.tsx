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
import { useAppStoreDispatch } from "@/hooks";
import { Loader } from "./Loaders";

const createPostSchema = object({
  title: string()
    .required()
    .min(5)
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
  const dispatch = useAppStoreDispatch();

  const queryClient = useQueryClient();
  const { mutate, isPending } = useMutation({
    mutationFn: postService.createPost,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["posts", userId],
      });

      dispatch({
        type: "ADD_NOTIFICATION",
        payload: {
          text: "Post added successfully",
          isSuccess: true,
        },
      });
      onClose();
    },
    onError: (err: Error | AxiosError) => {
      if (err instanceof AxiosError) {
        dispatch({
          type: "ADD_NOTIFICATION",
          payload: {
            text: err.response?.data.message || "Failed to create post",
            isSuccess: false,
          },
        });
      } else {
        dispatch({
          type: "ADD_NOTIFICATION",
          payload: {
            text: err.message || "Failed to create post",
            isSuccess: false,
          },
        });
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
        mutate({ ...data, userId: userId! });
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
          className="flex justify-center items-center px-4 py-[11.5px] border rounded-[4px] text-[14px] leading-[16.94px] font-normal text-[#334155] bg-white"
          onClick={() => onClose()}
        >
          Cancel
        </Button>
        <Button
          className="flex justify-center gap-x-2 items-center px-4 py-[11.5px] border rounded-[4px] bg-[#334155] font-semibold text-sm text-white"
          isSubmit
          isDisabled={isPending}
        >
          <span className="inline-block">Publish</span>
          {isPending && <Loader bg="bg-white" isSmall />}
        </Button>
      </div>
    </form>
  );
};
