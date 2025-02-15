import { User } from "@/models";
import axios from "axios";
import { BaseResponse, GetUsersResponse } from "@/types";

const getUsers = async (limit = 5, page = 1): Promise<GetUsersResponse> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.get(
    `${baseUrl}/users?limit=${limit}&page=${page}`
  );

  const data = response.data as BaseResponse<GetUsersResponse>;
  if (!data.success) {
    throw new Error(data.message);
  }
  console.log(data.data);
  return data.data!;
};

const getUser = async (userId: number): Promise<User> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.get(`${baseUrl}/users/${userId}`);
  const data = response.data as BaseResponse<User>;
  return data.data!;
};

const getUsersCount = async (): Promise<number> => {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
  const response = await axios.get(`${baseUrl}/users/count`);
  const data = response.data as BaseResponse<number>;
  return data.data!;
};

export const userService = {
  getUsers,
  getUser,
  getUsersCount,
};
