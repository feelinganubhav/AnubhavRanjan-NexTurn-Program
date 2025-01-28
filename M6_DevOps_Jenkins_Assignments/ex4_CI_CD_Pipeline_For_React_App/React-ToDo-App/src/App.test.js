import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import App from './App';

describe('App Component Tests', () => {
  beforeEach(() => {
    localStorage.clear(); 
  });

  test('renders the app with the correct title', () => {
    render(<App />);
    expect(screen.getByText(/Manage Your Todos/i)).toBeInTheDocument();
  });

  test('adds a new todo', () => {
    render(<App />);
    const input = screen.getByPlaceholderText(/Write Todo\.\.\./i); 
    const dateInput = screen.getByPlaceholderText(/Write Date\.\.\./i); 
    const button = screen.getByText(/Add/i); 

    fireEvent.change(input, { target: { value: 'Test Todo' } });
    fireEvent.change(dateInput, { target: { value: '2025-01-31' } });
    fireEvent.click(button);

    expect(screen.getByDisplayValue('Test Todo')).toBeInTheDocument();
    expect(screen.getByDisplayValue('2025-01-31')).toBeInTheDocument();
  });

  test('deletes a todo', () => {
    render(<App />);
    const input = screen.getByPlaceholderText(/Write Todo\.\.\./i);
    const dateInput = screen.getByPlaceholderText(/Write Date\.\.\./i);
    const button = screen.getByText(/Add/i);

    // Add a todo
    fireEvent.change(input, { target: { value: 'Todo to Delete' } });
    fireEvent.change(dateInput, { target: { value: '2025-01-31' } });
    fireEvent.click(button);
    screen.findByText(/Todo to Delete/i);

    expect(screen.getByDisplayValue(/Todo to Delete/i)).toBeInTheDocument();
    // screen.debug()

    // Delete the todo
    const deleteButton = screen.getByText(/❌/i); 
    fireEvent.click(deleteButton);

    // screen.debug();
    expect(screen.queryByDisplayValue(/Todo to Delete/i)).not.toBeInTheDocument();
  });

  test('toggles a todo as completed', () => {
    render(<App />);
    const input = screen.getByPlaceholderText(/Write Todo\.\.\./i);
    const dateInput = screen.getByPlaceholderText(/Write Date\.\.\./i);
    const button = screen.getByText(/Add/i);

    // Add a todo
    fireEvent.change(input, { target: { value: 'Todo to Complete' } });
    fireEvent.change(dateInput, { target: { value: '2025-01-31' } });
    fireEvent.click(button);

    // Mark as completed
    const checkbox = screen.getByRole('checkbox'); 
    fireEvent.click(checkbox);

    // Check if it's marked as completed
    const completedTodo = screen.getByDisplayValue('Todo to Complete');
    expect(completedTodo).toHaveClass('line-through'); 
  });

  test('loads todos from local storage', () => {
    const mockTodos = [
      { id: 1, todo: 'Stored Todo', completed: false, targetDate: '2025-01-31' },
    ];
    localStorage.setItem('key', JSON.stringify(mockTodos)); 

    render(<App />);
    // screen.debug()
    expect(screen.getByDisplayValue(/Stored Todo/i)).toBeInTheDocument();
    expect(screen.getByDisplayValue('2025-01-31')).toBeInTheDocument();
  });
});
