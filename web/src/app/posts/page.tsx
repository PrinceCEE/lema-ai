"use client";
import { Suspense } from "react";
import { postService, userService } from "@/api";
import { BackButton, Loader, PostContainer } from "@/components";
import { useQuery } from "@tanstack/react-query";
import { useRouter, useSearchParams } from "next/navigation";

const Users = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const userId = searchParams.get("userId");

  const { data: posts, isLoading: postsLoading } = useQuery({
    queryKey: ["posts", userId],
    queryFn: () => postService.getPosts(userId!),
  });

  const { data: user, isLoading: userLoading } = useQuery({
    queryKey: ["user", userId],
    queryFn: () => userService.getUser(userId!),
  });

  if (postsLoading || userLoading) {
    return (
      <div className="w-full h-[332px] flex items-center justify-center">
        <Loader />
      </div>
    );
  }

  return (
    <div className="md:w-[856px] w-full flex flex-col items-center mx-auto gap-y-6 h-full">
      <div className="w-full flex flex-col gap-4 sticky z-10 bg-white top-0 py-6">
        <BackButton
          text="Back to Users"
          handleClick={() => {
            router.back();
          }}
        />
        <h1 className="text-x font-medium text-[#181D27]">{user?.name}</h1>
        <div className="flex gap-2 items-center flex-wrap text-sm text-lightblack font-normal">
          <span>{user?.email}</span>
          <span>•</span>
          <span className="font-medium">{`${posts?.length} Posts`}</span>
        </div>
      </div>
      <div className="flex gap-4 flex-wrap">
        <>
          <PostContainer />
          {posts?.map((post) => (
            <PostContainer post={post} key={post.id} />
          ))}
        </>
      </div>
    </div>
  );
};

export default function Page() {
  return (
    <Suspense fallback={<Loader />}>
      <Users />
    </Suspense>
  );
}
