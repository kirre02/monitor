import { render, screen } from '@testing-library/react';
import {expect, test, vi} from 'vitest';
import SiteList from "./SiteList";

test('renders site list correctly', async () => {
  vi.mock('@tanstack/react-query', () => {
    const data = {
        sites: [
            { id: 1, name: 'Site 1' },
            { id: 2, name: 'Site 2' },
        ],
    };
    
    return {
        useQuery: () => ({
            isLoading: false,
            error: null,
            data: data,
        }),
        useMutation: vi.fn(),
        useQueryClient: vi.fn(),
    };
  });

    // Render the SiteList component
    render(<SiteList />);

    // Verify that site names are rendered correctly
    const siteNames = screen.getAllByText(/Site \d/);
    expect(siteNames).toHaveLength(2);
})
