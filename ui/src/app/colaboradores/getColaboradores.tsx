import { Colaborador } from "@/app/models";

export async function fetchColaboradores(): Promise<Colaborador[]> {
  const response = await fetch("http://localhost:8080/colaboradores", {
    cache: "no-cache",
  });
  let data = await response.json();
  console.log(data);

  // Format the admission_date for each Colaborador
  data = data.map((colaborador: Colaborador) => ({
    ...colaborador,
    admission_date: formatDate(colaborador.admission_date),
  }));
  return data;
}

export function formatDate(dateString: string): string {
  const date = new Date(dateString);
  const day = String(date.getDate()).padStart(2, "0");
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const year = date.getFullYear();
  return `${day}/${month}/${year}`; // Returns the date in dd/mm/yyyy format
}
