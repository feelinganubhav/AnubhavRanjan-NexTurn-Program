import { capitalize, reverseString } from '../src/a1_Fundamentals_String_Utilities';

describe('String Utility Functions', () => {
  describe('capitalize Function', () => {
    it('should capitalize the first letter of a word', () => {
      expect(capitalize('hello')).toBe('Hello');
    });

    it('should return an empty string if input is empty', () => {
      expect(capitalize('')).toBe('');
    });

    it('should handle single-character words', () => {
      expect(capitalize('a')).toBe('A');
      expect(capitalize('z')).toBe('Z');
    });

    it('should not modify already capitalized words', () => {
      expect(capitalize('World')).toBe('World');
    });

    it('should handle non-alphabetic characters at the start', () => {
      expect(capitalize('1world')).toBe('1world');
      expect(capitalize('!hello')).toBe('!hello');
    });

    it('should handle null or undefined gracefully', () => {
      expect(capitalize(null)).toBe('');
      expect(capitalize(undefined)).toBe('');
    });
  });

  
  describe('reverseString Function', () => {
    it('should reverse a given string', () => {
      expect(reverseString('hello')).toBe('olleh');
    });

    it('should return an empty string when input is empty', () => {
      expect(reverseString('')).toBe('');
    });

    it('should handle single-character strings', () => {
      expect(reverseString('a')).toBe('a');
    });

    it('should correctly reverse palindromes', () => {
      expect(reverseString('madam')).toBe('madam');
    });

    it('should handle strings with spaces and special characters', () => {
      expect(reverseString('a b c')).toBe('c b a');
      expect(reverseString('!hello')).toBe('olleh!');
    });

    it('should not throw errors for null or undefined inputs', () => {
      expect(() => reverseString(null)).not.toThrow();
      expect(() => reverseString(undefined)).not.toThrow();
    });
  });
});
