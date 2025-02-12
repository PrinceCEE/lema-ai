import { UserTable } from "@/components";

export default function Users() {
  return (
    <div className="md:w-[856px] w-full flex flex-col items-center mx-auto gap-y-6 h-full">
      <h1 className="text-xl font-medium self-start text-left p-0 text-black">
        Users
      </h1>
      <UserTable />
    </div>
  );
}
