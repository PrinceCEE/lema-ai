import { UserTable } from "@/components";

export default function Users() {
  return (
    <div className="flex flex-col w-[856px] gap-6 h-max">
      <h1 className="text-xl font-medium text-black">Users</h1>
      <UserTable />
    </div>
  );
}
