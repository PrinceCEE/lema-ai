import axios from "axios";
import { Post } from "@/models";
import { BaseResponse } from "@/types";

const createPost = async (payload: {
  userId: string;
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

const getPosts = async (userId: string): Promise<Post[]> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.get(`${baseUrl}/posts?user_id=${userId}`);
  const data = response.data as BaseResponse<Post[]>;
  if (!data.success) {
    throw new Error(data.message);
  }

  return data.data!;
};

const deletePost = async (postId: string): Promise<void> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.delete(`${baseUrl}/posts/${postId}`, {
    method: "DELETE",
  });

  const data = response.data as BaseResponse<null>;
  if (!data.success) {
    throw new Error(data.message);
  }
};

export const postService = {
  createPost,
  getPosts,
  deletePost,
};
