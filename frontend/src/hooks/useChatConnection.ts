import { useCallback, useEffect, useReducer, useRef } from 'react';
import { chatReducer, initialChatState } from '../chat/chatReducer';
import type { ClientEvent, ServerEvent } from '../types/chat';

const reconnectDelayMs = 1200;

type Options = {
  url: string;
};

export function useChatConnection({ url }: Options) {
  const [state, dispatch] = useReducer(chatReducer, initialChatState);
  const socketRef = useRef<WebSocket | null>(null);
  const nameRef = useRef<string | null>(null);
  const reconnectTimerRef = useRef<number | null>(null);
  const manualCloseRef = useRef(false);

  const clearReconnectTimer = useCallback(() => {
    if (reconnectTimerRef.current !== null) {
      window.clearTimeout(reconnectTimerRef.current);
      reconnectTimerRef.current = null;
    }
  }, []);

  const sendClientEvent = useCallback((event: ClientEvent) => {
    const socket = socketRef.current;
    if (!socket || socket.readyState !== WebSocket.OPEN) {
      dispatch({ type: 'client-error', error: 'conexao indisponivel' });
      return;
    }
    socket.send(JSON.stringify(event));
  }, []);

  const connect = useCallback(
    (name: string) => {
      clearReconnectTimer();
      manualCloseRef.current = false;
      nameRef.current = name;
      dispatch({ type: 'status', status: 'connecting' });

      const socket = new WebSocket(url);
      socketRef.current = socket;

      socket.onopen = () => {
        dispatch({ type: 'status', status: 'connected' });
        socket.send(JSON.stringify({ type: 'join', name } satisfies ClientEvent));
      };

      socket.onmessage = (event) => {
        try {
          const serverEvent = JSON.parse(event.data) as ServerEvent;
          dispatch({ type: 'server-event', event: serverEvent });
        } catch {
          dispatch({ type: 'client-error', error: 'evento invalido recebido do servidor' });
        }
      };

      socket.onerror = () => {
        dispatch({ type: 'status', status: 'error' });
        dispatch({ type: 'client-error', error: 'falha na conexao' });
      };

      socket.onclose = () => {
        socketRef.current = null;
        dispatch({ type: 'status', status: 'disconnected' });
        if (!manualCloseRef.current && nameRef.current) {
          reconnectTimerRef.current = window.setTimeout(() => connect(nameRef.current!), reconnectDelayMs);
        }
      };
    },
    [clearReconnectTimer, url],
  );

  const disconnect = useCallback(() => {
    manualCloseRef.current = true;
    clearReconnectTimer();
    socketRef.current?.close();
    socketRef.current = null;
    dispatch({ type: 'status', status: 'disconnected' });
  }, [clearReconnectTimer]);

  const sendMessage = useCallback(
    (content: string) => {
      sendClientEvent({ type: 'message', content });
    },
    [sendClientEvent],
  );

  useEffect(() => {
    return () => {
      manualCloseRef.current = true;
      clearReconnectTimer();
      socketRef.current?.close();
    };
  }, [clearReconnectTimer]);

  return {
    state,
    connect,
    disconnect,
    sendMessage,
  };
}
