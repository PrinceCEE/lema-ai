import { baseUrl } from "@/constants";
import { Post } from "@/models";
import { BaseResponse } from "@/types";

const createPost = async (payload: {
  userId: number;
  title: string;
  body: string;
}): Promise<Post> => {
  console.log(payload);
  const response = await fetch(`${baseUrl}/posts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  const data = (await response.json()) as BaseResponse<Post>;
  return data.data!;
};

const getPosts = async (userId: number): Promise<Post[]> => {
  const response = await fetch(`${baseUrl}/posts?user_id=${userId}`);
  const data = (await response.json()) as BaseResponse<Post[]>;

  return data.data!;
};

const deletePost = async (postId: number): Promise<void> => {
  await fetch(`${baseUrl}/posts/${postId}`, {
    method: "DELETE",
  });
};

export const postService = {
  createPost,
  getPosts,
  deletePost,
};
