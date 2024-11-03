export default function PageTitle({ title }: { title: string }) {
  return (
    <div className="mb-5  text-white">
      <h1 className="text-3xl font-bold">{title}</h1>
    </div>
  );
}
