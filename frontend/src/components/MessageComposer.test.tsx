import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { describe, expect, it, vi } from 'vitest';
import MessageComposer from './MessageComposer';

describe('MessageComposer', () => {
  it('sends trimmed messages and clears the input', async () => {
    const onSend = vi.fn();
    const user = userEvent.setup();
    render(<MessageComposer disabled={false} onSend={onSend} />);

    await user.type(screen.getByLabelText('Mensagem'), '  ola  ');
    await user.click(screen.getByRole('button', { name: 'Enviar' }));

    expect(onSend).toHaveBeenCalledWith('ola');
    expect(screen.getByLabelText('Mensagem')).toHaveValue('');
  });

  it('disables send button when disconnected', () => {
    render(<MessageComposer disabled={true} onSend={() => undefined} />);
    expect(screen.getByRole('button', { name: 'Enviar' })).toBeDisabled();
  });
});
