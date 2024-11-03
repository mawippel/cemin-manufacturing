"use client";

import Layout from "@/components/Layout";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { fetchColaboradores } from "./getColaboradores";
import Divider from "@/components/Divider";
import PageTitle from "@/components/PageTitle";
import { format } from "date-fns";
import { useEffect, useState } from "react";
import { Colaborador } from "../models";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { z } from "zod";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Calendar } from "@/components/ui/calendar";

const formSchema = z.object({
  nome: z.string().min(5).max(50),
  dataAdmissao: z.date().max(new Date()),
  salario: z.coerce.number().min(0),
});

export default function ColaboradoresPage() {
  const [open, setOpen] = useState(false);
  const [colaboradores, setColaboradores] = useState<Colaborador[]>([]);

  useEffect(() => {
    const loadColaboradores = async () => {
      const data = await fetchColaboradores();
      setColaboradores(data);
    };

    loadColaboradores();
  }, [open]);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      nome: "",
      dataAdmissao: new Date(),
      salario: 1412.0,
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const formattedValues = {
      name: values.nome,
      admission_date: values.dataAdmissao.toISOString(),
      salary: values.salario,
    };

    try {
      const response = await fetch("http://localhost:8080/colaboradores", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formattedValues),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      const data = await response.json();
      console.log("Success:", data);
    } catch (error) {
      console.error("Error:", error);
    }

    setOpen(false);
  }

  return (
    <Layout>
      <PageTitle title="Colaboradores" />
      <Divider />
      <main className="p-5 flex flex-col">
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button className="self-end" onClick={() => setOpen(true)}>
              Adicionar
            </Button>
          </DialogTrigger>

          <DialogContent className="sm:max-w-[425px]">
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="space-y-8"
              >
                <DialogHeader>
                  <DialogTitle>Adicionar Colaborador</DialogTitle>
                </DialogHeader>
                <Divider />

                <FormField
                  control={form.control}
                  name="nome"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Nome:</FormLabel>
                      <FormControl>
                        <Input placeholder="Fulano" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="dataAdmissao"
                  render={({ field }) => (
                    <FormItem className="w-full">
                      <FormLabel className="block">Data de Admissao:</FormLabel>
                      <Popover>
                        <PopoverTrigger asChild>
                          <FormControl>
                            <Button variant={"outline"}>
                              {field.value ? (
                                format(field.value, "dd/MM/yyyy")
                              ) : (
                                <span>Pick a date</span>
                              )}
                              <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                            </Button>
                          </FormControl>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0" align="start">
                          <Calendar
                            mode="single"
                            selected={field.value}
                            onSelect={field.onChange}
                            initialFocus
                          />
                        </PopoverContent>
                      </Popover>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="salario"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Salario:</FormLabel>
                      <FormControl>
                        <Input type="number" placeholder="1000.00" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <DialogFooter>
                  <Button type="submit">Salvar</Button>
                </DialogFooter>
              </form>
            </Form>
          </DialogContent>
        </Dialog>

        <Table className="mt-3 w-full">
          <TableHeader>
            <TableRow>
              <TableHead>ID</TableHead>
              <TableHead>Nome</TableHead>
              <TableHead>Data Admissao</TableHead>
              <TableHead className="text-right">Salario</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {colaboradores.map((col) => (
              <TableRow key={col.id}>
                <TableCell>{col.id}</TableCell>
                <TableCell>{col.name}</TableCell>
                <TableCell>{col.admission_date}</TableCell>
                <TableCell className="text-right">{col.salary}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </main>
    </Layout>
  );
}
