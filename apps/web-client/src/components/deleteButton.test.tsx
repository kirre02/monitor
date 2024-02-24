import { fireEvent, render } from "@testing-library/react";
import { describe, expect, test, vi } from "vitest";
import DeleteButton from "./deleteButton";

describe('delete button', () => {
    test('renders button with trash icon', () => {
        const onClickMock = vi.fn() // Create a mock function for onClick

        const {getByRole} = render(<DeleteButton onClick={onClickMock}/>)

        const buttonElement =  getByRole('button')

        expect(buttonElement).toContainHTML('<svg class="h-5 w-5"')

        expect(buttonElement).toHaveClass('text-red-400 hover:text-red-800');
    })

    test('calls onClick function when clicked', () => {
        const onClickMock = vi.fn()

        const { getByRole } = render(<DeleteButton onClick={onClickMock} />);

        // Find the button element by its role
        const buttonElement = getByRole('button');
    
        // Simulate a click event on the button
        fireEvent.click(buttonElement);
    
        // Check if the onClick function is called
        expect(onClickMock).toHaveBeenCalled();
    })
})