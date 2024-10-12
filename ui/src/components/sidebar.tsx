import Link from "next/link";

export default function Sidebar() {
    return (
        <aside className="w-64 bg-gray-800 text-gray-200">
        <div className="p-4 text-xl font-bold">Cemin Confeccoes</div>
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