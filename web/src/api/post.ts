import axios from "axios";
import { Post } from "@/models";
import { BaseResponse } from "@/types";

const createPost = async (payload: {
  userId: number;
  title: string;
  body: string;
}): Promise<Post> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.post(`${baseUrl}/posts`, payload, {
    headers: {
      "Content-Type": "application/json",
    },
  });

  const data = response.data as BaseResponse<Post>;
  if (!data.success) {
    throw new Error(data.message);
  }

  return data.data!;
};

const getPosts = async (userId: number): Promise<Post[]> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.get(`${baseUrl}/posts?user_id=${userId}`);
  const data = response.data as BaseResponse<Post[]>;

  return data.data!;
};

const deletePost = async (postId: number): Promise<void> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  await axios.delete(`${baseUrl}/posts/${postId}`, {
    method: "DELETE",
  });
};

export const postService = {
  createPost,
  getPosts,
  deletePost,
};
