import React from 'react';
import { render, screen, fireEvent } from "@testing-library/react";
import Todoitem from "../Todoitem";
import UseTodocontext from "../../Context/Todocontext";

// Mock the context
jest.mock("../../Context/Todocontext");

describe("TodoItem Component", () => {
  it("should render the todo item", () => {
    UseTodocontext.mockReturnValue({
      updatetodo: jest.fn(),
      deletetodo: jest.fn(),
      togglecomplete: jest.fn(),
    });

    const todo = {
      id: 1,
      todo: "Sample Todo",
      completed: false,
      targetDate: "2025-01-31",
    };

    render(<Todoitem todo={todo} />);

    expect(screen.getByDisplayValue(/sample todo/i)).toBeInTheDocument();
    expect(screen.getByDisplayValue("2025-01-31")).toBeInTheDocument();
  });

  it("should call togglecomplete when checkbox is clicked", () => {
    const togglecompleteMock = jest.fn();
    UseTodocontext.mockReturnValue({
      togglecomplete: togglecompleteMock,
    });

    const todo = {
      id: 1,
      todo: "Sample Todo",
      completed: false,
      targetDate: "2025-01-31",
    };

    render(<Todoitem todo={todo} />);

    const checkbox = screen.getByRole("checkbox");
    fireEvent.click(checkbox);

    expect(togglecompleteMock).toHaveBeenCalledWith(1);
  });

  it("should call deletetodo when delete button is clicked", () => {
    const deletetodoMock = jest.fn();
    UseTodocontext.mockReturnValue({
      deletetodo: deletetodoMock,
    });

    const todo = {
      id: 1,
      todo: "Sample Todo",
      completed: false,
      targetDate: "2025-01-31",
    };

    render(<Todoitem todo={todo} />);

    const deleteButton = screen.getByText(/❌/i);
    fireEvent.click(deleteButton);

    expect(deletetodoMock).toHaveBeenCalledWith(1);
  });
});
