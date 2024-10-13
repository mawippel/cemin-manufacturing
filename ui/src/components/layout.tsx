// components/Layout.js
import { PropsWithChildren } from "react";
import Sidebar from "./sidebar";

export default function Layout({ children }: PropsWithChildren) {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="flex-grow text-white flex-1 p-6 bg-gray-800">
        {children}
      </main>
    </div>
  );
}
