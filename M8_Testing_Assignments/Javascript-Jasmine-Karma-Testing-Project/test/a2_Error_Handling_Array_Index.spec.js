
import { getElement } from '../src/a2_Error_Handling_Array_Index';

describe('Array Index Error Handling', () => {
  let testArray;

  beforeEach(() => {
    testArray = [10, 20, 30, 40, 50];
  });


  it('should return the correct element for a valid index', () => {
    expect(getElement(testArray, 0)).toBe(10);
    expect(getElement(testArray, 2)).toBe(30);
    expect(getElement(testArray, 4)).toBe(50);
  });


  it('should throw an error for a negative index', () => {
    expect(() => getElement(testArray, -1)).toThrowError("Index out of bounds");
  });

  it('should throw an error for an index greater than or equal to the array length', () => {
    expect(() => getElement(testArray, 5)).toThrowError("Index out of bounds");
    expect(() => getElement(testArray, 100)).toThrowError("Index out of bounds");
  });

  it('should throw an error if the array is empty', () => {
    const emptyArray = [];
    expect(() => getElement(emptyArray, 0)).toThrowError("Index out of bounds");
  });

  it('should throw an error if the index is NaN', () => {
    expect(() => getElement(testArray, NaN)).toThrowError("Index must be an integer");
  });

  it('should throw an error for non-integer indices', () => {
    expect(() => getElement(testArray, 1.5)).toThrowError("Index must be an integer");
    expect(() => getElement(testArray, "1")).toThrowError("Index must be an integer");
  });
});
