import Link from "next/link";

export default function Sidebar() {
    return (
        <aside className="w-64 bg-gray-900 text-white">
        <div className="flex items-center justify-center text-center">
            <h1 className="font-custom text-4xl m-4">Cemin Confecções</h1>
        </div>
        <nav className="p-4">
          <ul>
            <li>
              <Link href="/colaboradores">
              Colaboradores
              </Link>
            </li>
            <li>
              <Link href="/modelos">
              Modelos
              </Link>
            </li>
            <li>
              <Link href="/operacoes">
              Operacoes
              </Link>
            </li>
            <li>
              <Link href="/execucoes">
                Execucoes
              </Link>
            </li>
          </ul>
        </nav>
      </aside>
    );
}