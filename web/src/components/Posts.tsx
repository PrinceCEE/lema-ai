"use client";
import { Post } from "@/models";
import { FC, useState } from "react";
import { IoAddCircleOutline } from "react-icons/io5";
import { RiDeleteBinLine } from "react-icons/ri";
import { ModalOverlay } from "./ModalOverlay";
import { postService } from "@/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useSearchParams } from "next/navigation";

export const PostContainer: FC<{ post?: Post }> = ({ post }) => {
  return (
    <div className="w-[270px] h-[293px] border p-6 rounded-lg border-[#D5D7DA]">
      {post ? <DisplayPost post={post} /> : <NewPost />}
    </div>
  );
};

export const NewPost = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  const openModal = () => setIsModalOpen(true);
  const closeModal = () => setIsModalOpen(false);

  return (
    <div className="flex flex-col justify-center items-center h-full w-full text-[#717680]">
      <div
        className="w-max flex flex-col justify-center items-center gap-2 cursor-pointer"
        onClick={(e) => {
          e.stopPropagation();
          openModal();
        }}
      >
        <IoAddCircleOutline className="w-6 h-6" />
        <span className="inline-block font-semibold text-sm">New Post</span>
      </div>
      <ModalOverlay isOpen={isModalOpen} onClose={closeModal} />
    </div>
  );
};

export const DisplayPost: FC<{ post: Post }> = ({ post }) => {
  const queryClient = useQueryClient();
  const searchParams = useSearchParams();
  const userId = searchParams.get("userId");

  const { mutate, isError } = useMutation({
    mutationFn: postService.deletePost,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["posts", userId],
      });
    },
  });

  if (isError) {
    alert("Failed to delete post");
  }

  const handleClick = async () => {
    mutate(post.id);
  };

  return (
    <div className="w-full h-full flex flex-col gap-4 text-lightblack relative">
      <h1 className="text-[18px] leading-5 font-medium">{post.title}</h1>
      <p className="text-sm font-normal overflow-hidden flex-auto overflow-ellipsis">
        {post.body}
      </p>
      <RiDeleteBinLine
        className="absolute top-[-15px] right-[-15px] text-red-400 cursor-pointer"
        onClick={handleClick}
      />
    </div>
  );
};
