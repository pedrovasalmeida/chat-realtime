import type { ChatState, ConnectionStatus, ServerEvent } from '../types/chat';

export const initialChatState: ChatState = {
  status: 'idle',
  users: [],
  messages: [],
  error: null,
};

export type ChatAction =
  | { type: 'status'; status: ConnectionStatus }
  | { type: 'server-event'; event: ServerEvent }
  | { type: 'client-error'; error: string | null }
  | { type: 'reset' };

export function chatReducer(state: ChatState, action: ChatAction): ChatState {
  switch (action.type) {
    case 'status':
      return { ...state, status: action.status, error: action.status === 'connected' ? null : state.error };
    case 'server-event':
      return applyServerEvent(state, action.event);
    case 'client-error':
      return { ...state, error: action.error };
    case 'reset':
      return initialChatState;
    default:
      return state;
  }
}

function applyServerEvent(state: ChatState, event: ServerEvent): ChatState {
  switch (event.type) {
    case 'presence':
      return { ...state, users: event.users };
    case 'message':
      return { ...state, messages: [...state.messages, event.message] };
    case 'error':
      return { ...state, error: event.error };
    default:
      return state;
  }
}
