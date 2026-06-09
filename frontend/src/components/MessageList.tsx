import type { ChatMessage } from '../types/chat';

type Props = {
  messages: ChatMessage[];
};

export default function MessageList({ messages }: Props) {
  if (messages.length === 0) {
    return <div className="flex flex-1 items-center justify-center text-slate-500">Nenhuma mensagem ainda.</div>;
  }

  return (
    <ol className="flex flex-1 flex-col gap-3 overflow-y-auto p-4">
      {messages.map((message) => (
        <li key={message.id} className="rounded-2xl bg-white p-4 shadow-sm">
          <div className="flex items-center justify-between gap-3 text-sm text-slate-500">
            <strong className="text-slate-800">{message.userName || message.userId}</strong>
            <time dateTime={message.sentAt}>{formatTime(message.sentAt)}</time>
          </div>
          <p className="mt-2 whitespace-pre-wrap text-slate-900">{message.content}</p>
        </li>
      ))}
    </ol>
  );
}

function formatTime(value: string) {
  return new Intl.DateTimeFormat('pt-BR', {
    hour: '2-digit',
    minute: '2-digit',
  }).format(new Date(value));
}
