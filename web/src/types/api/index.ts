import { User } from "@/models";

export interface BaseResponse<T = null> {
  message: string;
  success: boolean;
  data?: T;
}

export type GetUsersResponse = {
  users: User[];
  count: number;
  total_pages: number;
  page: number;
  limit: number;
  has_next: boolean;
  has_prev: boolean;
};
