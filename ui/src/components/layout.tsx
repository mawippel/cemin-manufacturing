// components/Layout.js
import { PropsWithChildren } from 'react';
import Sidebar from './sidebar';

export default function Layout({ children }: PropsWithChildren) {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <main className="flex-1 p-6 bg-gray-100">{children}</main>
    </div>
  );
}
