import { baseUrl } from "@/constants";
import { User } from "@/models";
import { BaseResponse, GetUsersResponse } from "@/types";

const getUsers = async (limit = 5, page = 1): Promise<GetUsersResponse> => {
  const response = await fetch(`${baseUrl}/users?limit=${limit}&page=${page}`);
  const data = (await response.json()) as BaseResponse<GetUsersResponse>;

  return data.data!;
};

const getUser = async (userId: number): Promise<User> => {
  const response = await fetch(`${baseUrl}/users/${userId}`);
  const data = (await response.json()) as BaseResponse<User>;
  return data.data!;
};

const getUsersCount = async (): Promise<number> => {
  const response = await fetch(`${baseUrl}/users/count`);
  const data = (await response.json()) as BaseResponse<number>;
  return data.data!;
};

export const userService = {
  getUsers,
  getUser,
  getUsersCount,
};
