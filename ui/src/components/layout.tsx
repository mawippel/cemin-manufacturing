import { PropsWithChildren } from "react";
import Sidebar from "./sidebar";

export default function Layout({ children }: PropsWithChildren) {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="bg-gray-900 flex-grow flex-1 p-6">{children}</main>
    </div>
  );
}
