import { type FormEvent, useState } from 'react';

type Props = {
  disabled: boolean;
  onSend: (content: string) => void;
};

export default function MessageComposer({ disabled, onSend }: Props) {
  const [content, setContent] = useState('');
  const trimmedContent = content.trim();

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (disabled || !trimmedContent) {
      return;
    }
    onSend(trimmedContent);
    setContent('');
  }

  return (
    <form onSubmit={handleSubmit} className="flex gap-3 border-t border-slate-200 bg-white p-4">
      <label className="sr-only" htmlFor="message-input">
        Mensagem
      </label>
      <input
        id="message-input"
        className="min-w-0 flex-1 rounded-2xl border border-slate-300 px-4 py-3 outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
        value={content}
        onChange={(event) => setContent(event.target.value)}
        disabled={disabled}
        maxLength={500}
      />
      <button
        className="rounded-2xl bg-indigo-600 px-5 py-3 font-semibold text-white transition hover:bg-indigo-500 disabled:cursor-not-allowed disabled:bg-slate-300"
        disabled={disabled || !trimmedContent}
        type="submit"
      >
        Enviar
      </button>
    </form>
  );
}
