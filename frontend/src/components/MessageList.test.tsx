import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import MessageList from './MessageList';

describe('MessageList', () => {
  it('renders author, content, and time', () => {
    render(
      <MessageList
        messages={[
          {
            id: 'm1',
            userId: 'u1',
            userName: 'Alice',
            content: 'ola',
            sentAt: '2026-06-09T18:30:00Z',
          },
        ]}
      />,
    );

    expect(screen.getByText('Alice')).toBeInTheDocument();
    expect(screen.getByText('ola')).toBeInTheDocument();
    expect(screen.getByText(/\d{2}:\d{2}/)).toBeInTheDocument();
  });
});
