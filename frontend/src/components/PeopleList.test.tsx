import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import PeopleList from './PeopleList';

describe('PeopleList', () => {
  it('renders online people', () => {
    render(<PeopleList users={[{ id: 'u1', name: 'Alice' }, { id: 'u2', name: 'Bob' }]} />);

    expect(screen.getByText('Pessoas online')).toBeInTheDocument();
    expect(screen.getByText('Alice')).toBeInTheDocument();
    expect(screen.getByText('Bob')).toBeInTheDocument();
  });
});
