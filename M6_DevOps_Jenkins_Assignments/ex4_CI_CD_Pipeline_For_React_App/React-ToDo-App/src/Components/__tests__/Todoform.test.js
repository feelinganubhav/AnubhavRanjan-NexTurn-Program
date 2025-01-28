import React from 'react';
import { render, screen, fireEvent } from "@testing-library/react";
import Todoform from "../Todoform";
import UseTodocontext from "../../Context/Todocontext";

// Mock the context
jest.mock("../../Context/Todocontext");

describe("TodoForm Component", () => {
  it("should render the form", () => {
    UseTodocontext.mockReturnValue({ addtodo: jest.fn() });

    render(<Todoform />);

    expect(screen.getByPlaceholderText(/write todo/i)).toBeInTheDocument();
    expect(screen.getByPlaceholderText(/write date/i)).toBeInTheDocument();
    expect(screen.getByText(/add/i)).toBeInTheDocument();
  });

  it("should call addtodo when the form is submitted", () => {
    const addtodoMock = jest.fn();
    UseTodocontext.mockReturnValue({ addtodo: addtodoMock });

    render(<Todoform />);

    const todoInput = screen.getByPlaceholderText(/write todo/i);
    const dateInput = screen.getByPlaceholderText(/write date/i);
    const addButton = screen.getByText(/add/i);

    fireEvent.change(todoInput, { target: { value: "New Todo" } });
    fireEvent.change(dateInput, { target: { value: "2025-01-31" } });
    fireEvent.click(addButton);

    expect(addtodoMock).toHaveBeenCalledWith({
      todo: "New Todo",
      completed: false,
      targetDate: "2025-01-31",
    });
  });
});
