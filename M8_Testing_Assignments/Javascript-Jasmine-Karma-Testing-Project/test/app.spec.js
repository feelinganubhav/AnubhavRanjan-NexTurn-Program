import { add } from '../src/app';

describe('Addition function', () => {
  it('should add two numbers correctly', () => {
    expect(add(2, 3)).toBe(5);
  });

  it('should handle negative numbers', () => {
    expect(add(-2, -3)).toBe(-5);
  });
});
