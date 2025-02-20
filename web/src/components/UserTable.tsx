"use client";

import { useQuery } from "@tanstack/react-query";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { Loader } from "@/components";
import { userService } from "@/api";
import { useEffect, useState } from "react";
import { useAppStoreDispatch } from "@/hooks";
import { Paginator } from "./Pagination";

export const UserTable = () => {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const [page, setPage] = useState(
    searchParams.get("page") ? parseInt(searchParams.get("page") as string) : 1
  );
  const dispatch = useAppStoreDispatch();
  const limit = 4;

  useEffect(() => {
    if (!searchParams.has("page")) {
      router.replace(`${pathname}?page=1`);
      setPage(1);
    } else {
      const newPage = parseInt(searchParams.get("page") as string);
      setPage(newPage);
    }
  }, [searchParams, pathname, router]);

  const { data, isLoading, error } = useQuery({
    queryKey: ["users", page],
    queryFn: () => userService.getUsers(limit, page),
  });

  const handlePageClick = (page: number) => {
    setPage(page);
    router.push(`${pathname}?page=${page}`);
  };

  if (error) {
    dispatch({
      type: "ADD_NOTIFICATION",
      payload: {
        text: error.message,
        isSuccess: false,
      },
    });
  }

  return (
    <div className="flex flex-col gap-y-6 justify-center w-full">
      <div className="border rounded-lg shadow-md text-lightblack min-h-[332px] overflow-x-scroll">
        <table className="w-full border-collapse">
          <thead>
            <tr className="text-xs font-medium h-[72px] text-center">
              <th className="pl-4 py-6 text-left">Full Name</th>
              <th className="pl-4 py-6 text-left">Email Address</th>
              <th className="pl-4 py-6 text-left">Address</th>
            </tr>
          </thead>
          <tbody className="text-sm font-normal">
            {isLoading ? (
              <tr>
                <td colSpan={3}>
                  <div className="w-full h-[332px] flex items-center justify-center">
                    <Loader />
                  </div>
                </td>
              </tr>
            ) : (
              data?.users?.map((user) => (
                <tr
                  key={user.id}
                  className="border-b text-left cursor-pointer h-[72px]"
                  onClick={() => {
                    router.push(`/posts?userId=${user.id}`);
                  }}
                >
                  <td className="pl-4 py-6 h-full whitespace-nowrap font-medium">
                    {user.name}
                  </td>
                  <td className="pl-4 py-6 h-full">{user.email}</td>
                  <td className="pl-4 md:max-w-[392px] py-6 h-full truncate">{`${user.address.street}, ${user.address.state}, ${user.address.city}, ${user.address.zipcode}`}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
      {data && (
        <Paginator
          currentIndex={page}
          lastIndex={data!.total_pages}
          isLeftButtonEnabled={data!.has_prev}
          isRightButtonEnabled={data!.has_next}
          onLeftButtonClick={() => handlePageClick(page - 1)}
          onRightButtonClick={() => handlePageClick(page + 1)}
          onPaginatorItemClick={handlePageClick}
        />
      )}
    </div>
  );
};
