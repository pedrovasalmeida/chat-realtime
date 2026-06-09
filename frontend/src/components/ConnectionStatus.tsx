import type { ConnectionStatus as Status } from '../types/chat';

type Props = {
  status: Status;
  error: string | null;
};

const labels: Record<Status, string> = {
  idle: 'Aguardando entrada',
  connecting: 'Conectando',
  connected: 'Conectado',
  disconnected: 'Desconectado',
  error: 'Erro',
};

export default function ConnectionStatus({ status, error }: Props) {
  return (
    <div className="rounded-full bg-white/80 px-3 py-1 text-sm font-medium text-slate-700 shadow-sm">
      <span>{labels[status]}</span>
      {error ? <span className="ml-2 text-red-600">Erro: {error}</span> : null}
    </div>
  );
}
