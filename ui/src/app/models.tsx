export interface BaseEntity {
  id: number;
  createdAt: string;
  updatedAt: string;
}

export interface Modelo extends BaseEntity {
  name: string;
}

export interface Operacao extends BaseEntity {
  name: string;
  time: number; // in seconds
  modeloId: number;
  modelo: Modelo;
}

export interface Colaborador extends BaseEntity {
  name: string;
  admission_date: string;
  salary: number;
}

export interface Execucao extends BaseEntity {
  modeloId: number;
  modelo: Modelo;
  operacaoId: number;
  operacao: Operacao;
  colaboradorId: number;
  colaborador: Colaborador;
  percProdutivo: number; // assuming float32 maps to number in TypeScript
}
