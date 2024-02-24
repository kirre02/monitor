import { render } from '@testing-library/react';
import { describe, expect, test } from "vitest";
import StatusBadge from './badge';
import { Status } from 'monitor-sdk/models';

describe('StatusBadge', () => {
  test('renders Up badge when status is up', () => {
    const status: Status = {
      up: true,
      checkedAt: '2022-02-15T12:00:00Z', // This value can be changed based on your requirements
    };

    const { getByText } = render(<StatusBadge status={status} />);
    const badgeElement = getByText('Up');
    expect(badgeElement).toBeInTheDocument();
    expect(badgeElement).toHaveClass('bg-green-100');
    expect(badgeElement).toHaveClass('text-green-800');
  });

  test('renders Down badge when status is down', () => {
    const status: Status = {
      up: false,
      checkedAt: '2022-02-15T12:00:00Z', // This value can be changed based on your requirements
    };

    const { getByText } = render(<StatusBadge status={status} />);
    const badgeElement = getByText('Down');
    expect(badgeElement).toBeInTheDocument();
    expect(badgeElement).toHaveClass('bg-red-100');
    expect(badgeElement).toHaveClass('text-red-800');
  });

  test('renders Unknown badge when status is undefined', () => {
    const { getByText } = render(<StatusBadge status={undefined} />);
    const badgeElement = getByText('Unknown');
    expect(badgeElement).toBeInTheDocument();
    expect(badgeElement).toHaveClass('bg-gray-100');
    expect(badgeElement).toHaveClass('text-gray-800');
  });
});