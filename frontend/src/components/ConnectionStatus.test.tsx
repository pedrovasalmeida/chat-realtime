import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import ConnectionStatus from './ConnectionStatus';

describe('ConnectionStatus', () => {
  it('shows connected state', () => {
    render(<ConnectionStatus status="connected" error={null} />);
    expect(screen.getByText('Conectado')).toBeInTheDocument();
  });

  it('shows error text', () => {
    render(<ConnectionStatus status="error" error="falha de conexao" />);
    expect(screen.getByText('Erro: falha de conexao')).toBeInTheDocument();
  });
});
