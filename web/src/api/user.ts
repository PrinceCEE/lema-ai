import { User } from "@/models";
import { BaseResponse, GetUsersResponse } from "@/types";

const getUsers = async (limit = 5, page = 1): Promise<GetUsersResponse> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  console.log(baseUrl);
  const response = await fetch(`${baseUrl}/users?limit=${limit}&page=${page}`);
  if (!response.ok) {
    throw new Error("Failed to fetch users");
  }

  const data = (await response.json()) as BaseResponse<GetUsersResponse>;
  if (!data.success) {
    throw new Error(data.message);
  }

  return data.data!;
};

const getUser = async (userId: number): Promise<User> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await fetch(`${baseUrl}/users/${userId}`);
  const data = (await response.json()) as BaseResponse<User>;
  return data.data!;
};

const getUsersCount = async (): Promise<number> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await fetch(`${baseUrl}/users/count`);
  const data = (await response.json()) as BaseResponse<number>;
  return data.data!;
};

export const userService = {
  getUsers,
  getUser,
  getUsersCount,
};
