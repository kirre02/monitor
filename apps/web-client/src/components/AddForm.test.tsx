import { render, screen, fireEvent } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, expect, test, vi } from "vitest";
import AddForm from "./AddForm";

describe("AddForm", () => {
  // Mock the ResizeObserver
  const ResizeObserverMock = vi.fn(() => ({
    observe: vi.fn(),
    unobserve: vi.fn(),
    disconnect: vi.fn(),
  }));

  // Stub the global ResizeObserver
  vi.stubGlobal("ResizeObserver", ResizeObserverMock);

  test("renders Add Site button", () => {
    render(<AddForm />);
    const addButton = screen.getByRole("button", { name: /Add site/i });
    expect(addButton).toBeInTheDocument();
  });

  test("opens dialog when Add Site button is clicked", async () => {
    render(<AddForm />);
    const addButton = screen.getByRole("button", { name: /Add site/i });
    fireEvent.click(addButton);
    const dialog = await screen.getByText("Add Site");
    expect(dialog).toBeInTheDocument();
  });

  test("submits form with valid input values", async () => {
    vi.mock("@tanstack/react-query", () => {
      return {
        useQuery: () => ({
          isLoading: false,
          error: null,
        }),
        useMutation: vi.fn(),
        useQueryClient: vi.fn(),
      };
    });
    render(<AddForm />);
    const addButton = screen.getByRole("button", { name: /Add Site/i });
    fireEvent.click(addButton);

    const nameInput = await screen.findByLabelText(/name/i);
    const urlInput = await screen.findByLabelText(/url/i);
    const submitButton = await screen.findByRole("button", { name: /submit/i });

    // Simulate user input
    fireEvent.change(nameInput, { target: { value: "Example Site" } });
    fireEvent.change(urlInput, { target: { value: "http://example.com" } });

    // Submit the form
    userEvent.click(submitButton);
  });
});
