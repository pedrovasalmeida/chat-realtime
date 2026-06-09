import { describe, expect, it } from 'vitest';
import { chatReducer, initialChatState } from './chatReducer';

describe('chatReducer', () => {
  it('stores presence snapshots', () => {
    const state = chatReducer(initialChatState, {
      type: 'server-event',
      event: {
        type: 'presence',
        users: [
          { id: 'u1', name: 'Alice' },
          { id: 'u2', name: 'Bob' },
        ],
      },
    });

    expect(state.users).toEqual([
      { id: 'u1', name: 'Alice' },
      { id: 'u2', name: 'Bob' },
    ]);
  });

  it('appends broadcast messages', () => {
    const state = chatReducer(initialChatState, {
      type: 'server-event',
      event: {
        type: 'message',
        message: {
          id: 'm1',
          userId: 'u1',
          userName: 'Alice',
          content: 'hello',
          sentAt: '2026-06-09T18:30:00Z',
        },
      },
    });

    expect(state.messages).toHaveLength(1);
    expect(state.messages[0].content).toBe('hello');
  });

  it('stores server errors and connection state changes', () => {
    const connected = chatReducer(initialChatState, { type: 'status', status: 'connected' });
    const failed = chatReducer(connected, {
      type: 'server-event',
      event: { type: 'error', error: 'message content is required' },
    });

    expect(connected.status).toBe('connected');
    expect(failed.error).toBe('message content is required');
  });
});
