"use client";

import { PropsWithChildren } from "react";
import Sidebar from "./sidebar";

export default function Layout({ children }: PropsWithChildren) {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="flex-grow flex-1 p-6">{children}</main>
    </div>
  );
}
