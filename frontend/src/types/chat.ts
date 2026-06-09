export type ConnectionStatus = 'idle' | 'connecting' | 'connected' | 'disconnected' | 'error';

export type User = {
  id: string;
  name: string;
};

export type ChatMessage = {
  id: string;
  userId: string;
  userName: string;
  content: string;
  sentAt: string;
};

export type ServerEvent =
  | { type: 'presence'; users: User[] }
  | { type: 'message'; message: ChatMessage }
  | { type: 'error'; error: string };

export type ClientEvent =
  | { type: 'join'; name: string }
  | { type: 'message'; content: string };

export type ChatState = {
  status: ConnectionStatus;
  users: User[];
  messages: ChatMessage[];
  error: string | null;
};
