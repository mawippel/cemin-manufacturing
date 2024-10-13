"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { buttonVariants } from "@/components/ui/button";
import { cn } from "@/lib/utils";

export default function Sidebar() {
  const pathname = usePathname();

  return (
    <aside className="p-5 w-64">
      <div className="flex items-center justify-center text-center">
        <h1 className="font-custom text-4xl m-4">Cemin Confecções</h1>
      </div>
      <div
        data-orientation="horizontal"
        className="shrink-0 bg-border h-[1px] w-full my-6"
      />
      <nav className="flex space-x-2 lg:flex-col lg:space-x-0 lg:space-y-1">
        <Link
          href="/colaboradores"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            pathname === "/colaboradores"
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          Colaboradores
        </Link>
        <Link
          href="/modelos"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            pathname === "/modelos"
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          Modelos
        </Link>
        <Link
          href="/operacoes"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            pathname === "/operacoes"
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          Operacoes
        </Link>
        <Link
          href="/execucoes"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            pathname === "/execucoes"
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          Execucoes
        </Link>
        <Link
          href="/dashboards"
          className={cn(
            buttonVariants({ variant: "ghost" }),
            pathname === "/dashboards"
              ? "bg-muted hover:bg-muted"
              : "hover:bg-transparent hover:underline",
            "justify-start"
          )}
        >
          Dashboards
        </Link>
      </nav>
    </aside>
  );
}
