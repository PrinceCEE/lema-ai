"use client";

import { useQuery } from "@tanstack/react-query";
import ReactPaginate from "react-paginate";
import { useRouter } from "next/navigation";
import { NextLabel, PreviousLabel, Loader } from "@/components";
import { userService } from "@/api";
import { useState } from "react";

export const UserTable = () => {
  const router = useRouter();
  const [page, setPage] = useState(0);
  const limit = 5;

  const { data, isLoading, error } = useQuery({
    queryKey: ["users", page],
    queryFn: () => userService.getUsers(limit, page + 1),
  });

  const handlePageClick = (event: { selected: number }) => {
    setPage(event.selected);
  };

  if (error) {
    alert(error.message);
  }

  return (
    <>
      <div className="border rounded-lg overflow-hidden shadow-md text-lightblack min-h-[332px]">
        <table className="w-full border-collapse">
          <thead>
            <tr className="text-xs font-medium h-[44px] text-center">
              <th className="pl-4 pt-[13px] text-left">Full Name</th>
              <th className="pl-4 pt-[13px] text-left">Email Address</th>
              <th className="pl-4 pt-[13px] text-left">Address</th>
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
                  className="border-b h-[72px] text-left cursor-pointer"
                  onClick={() => {
                    router.push(`/posts?userId=${user.id}`);
                  }}
                >
                  <td className="pl-4">{`${user.first_name} ${user.last_name}`}</td>
                  <td className="pl-4">{user.email}</td>
                  <td className="pl-4 w-[392px] truncate">{`${user.address.street}, ${user.address.state}, ${user.address.city}, ${user.address.zipcode}`}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
      <div className="flex flex-row self-end">
        <ReactPaginate
          breakLabel="..."
          nextLabel={<NextLabel />}
          previousLabel={<PreviousLabel />}
          className="flex items-center justify-center gap-2" // Added gap-2 class
          onPageChange={handlePageClick}
          pageRangeDisplayed={3}
          pageCount={data?.total_pages || 0}
          renderOnZeroPageCount={null}
          pageClassName="w-[40px] h-[40px] flex items-center justify-center text-sm font-medium text-lightblack hover:text-paginationBtnHoverText hover:bg-paginationBtnBg active:text-paginationBtnHoverText active:bg-paginationBtnBg cursor-pointer"
          activeClassName="bg-paginationBtnBg text-paginationBtnHoverText"
          forcePage={page}
        />
      </div>
    </>
  );
};
