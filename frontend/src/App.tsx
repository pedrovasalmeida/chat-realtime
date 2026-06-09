import { useState } from 'react';
import { initialChatState } from './chat/chatReducer';
import ConnectionStatus from './components/ConnectionStatus';
import JoinForm from './components/JoinForm';
import MessageComposer from './components/MessageComposer';
import MessageList from './components/MessageList';
import PeopleList from './components/PeopleList';
import { useChatConnection } from './hooks/useChatConnection';

const wsUrl = import.meta.env.VITE_WS_URL ?? 'ws://localhost:8080/ws';

export default function App() {
  const [joined, setJoined] = useState(false);
  const { state, connect, sendMessage } = useChatConnection({ url: wsUrl });
  const chatState = joined ? state : initialChatState;

  function handleJoin(name: string) {
    setJoined(true);
    connect(name);
  }

  if (!joined) {
    return (
      <main className="flex min-h-screen items-center justify-center overflow-hidden bg-[radial-gradient(circle_at_top_left,#f8c86b_0,#f8c86b_12rem,transparent_24rem),linear-gradient(135deg,#182033_0%,#24324d_48%,#e9ddc7_49%,#f7efe2_100%)] p-6">
        <JoinForm onJoin={handleJoin} />
      </main>
    );
  }

  const canSend = chatState.status === 'connected';

  return (
    <main className="min-h-screen bg-[linear-gradient(135deg,#172033_0%,#22334f_40%,#f2e5cf_40%,#fbf4e8_100%)] p-4 text-slate-950 md:p-8">
      <div className="mx-auto flex max-w-6xl flex-col gap-4">
        <header className="flex flex-col gap-3 rounded-3xl border border-white/60 bg-white/85 p-5 shadow-[0_24px_80px_rgba(15,23,42,0.18)] backdrop-blur md:flex-row md:items-center md:justify-between">
          <div>
            <p className="text-sm font-semibold uppercase tracking-[0.28em] text-amber-700">Sala geral</p>
            <h1 className="text-3xl font-black tracking-tight text-slate-950">Realtime Chat</h1>
          </div>
          <ConnectionStatus status={chatState.status} error={chatState.error} />
        </header>

        <section className="grid min-h-[70vh] gap-4 lg:grid-cols-[1fr_18rem]">
          <div className="flex min-h-[34rem] flex-col overflow-hidden rounded-3xl border border-white/70 bg-slate-50/95 shadow-[0_32px_90px_rgba(15,23,42,0.22)]">
            <MessageList messages={chatState.messages} />
            <MessageComposer disabled={!canSend} onSend={sendMessage} />
          </div>
          <PeopleList users={chatState.users} />
        </section>
      </div>
    </main>
  );
}
