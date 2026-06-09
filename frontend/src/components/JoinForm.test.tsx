import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import JoinForm from './JoinForm';

describe('JoinForm', () => {
  it('submits trimmed name', async () => {
    const onJoin = vi.fn();
    const user = userEvent.setup();
    render(<JoinForm onJoin={onJoin} />);

    await user.type(screen.getByLabelText('Seu nome'), '  Pedro  ');
    await user.click(screen.getByRole('button', { name: 'Entrar no chat' }));

    expect(onJoin).toHaveBeenCalledWith('Pedro');
  });
});
