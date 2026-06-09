import type { User } from '../types/chat';

type Props = {
  users: User[];
};

export default function PeopleList({ users }: Props) {
  return (
    <aside className="rounded-3xl bg-white p-5 shadow-sm lg:w-72">
      <h2 className="text-lg font-semibold text-slate-950">Pessoas online</h2>
      <p className="mt-1 text-sm text-slate-500">{users.length} conectado(s)</p>
      <ul className="mt-4 space-y-2">
        {users.map((user) => (
          <li key={user.id} className="flex items-center gap-3 rounded-2xl bg-slate-50 px-3 py-2 text-slate-800">
            <span className="h-2.5 w-2.5 rounded-full bg-emerald-500" aria-hidden="true" />
            <span>{user.name || user.id}</span>
          </li>
        ))}
      </ul>
    </aside>
  );
}
