import { type FormEvent, useState } from 'react';

type Props = {
  onJoin: (name: string) => void;
};

export default function JoinForm({ onJoin }: Props) {
  const [name, setName] = useState('');
  const trimmedName = name.trim();

  function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (trimmedName) {
      onJoin(trimmedName);
    }
  }

  return (
    <form onSubmit={handleSubmit} className="mx-auto flex max-w-md flex-col gap-4 rounded-3xl bg-white p-8 shadow-xl">
      <div>
        <h1 className="text-3xl font-bold text-slate-950">Realtime Chat</h1>
        <p className="mt-2 text-slate-600">Informe seu nome para entrar na sala geral.</p>
      </div>
      <label className="flex flex-col gap-2 text-sm font-medium text-slate-700">
        Seu nome
        <input
          className="rounded-2xl border border-slate-300 px-4 py-3 text-base outline-none focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
          value={name}
          onChange={(event) => setName(event.target.value)}
          maxLength={40}
        />
      </label>
      <button
        className="rounded-2xl bg-indigo-600 px-5 py-3 font-semibold text-white transition hover:bg-indigo-500 disabled:cursor-not-allowed disabled:bg-slate-300"
        disabled={!trimmedName}
        type="submit"
      >
        Entrar no chat
      </button>
    </form>
  );
}
