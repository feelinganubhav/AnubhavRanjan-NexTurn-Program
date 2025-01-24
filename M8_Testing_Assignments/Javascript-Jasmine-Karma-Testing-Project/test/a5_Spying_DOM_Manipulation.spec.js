import { toggleVisibility } from '../src/a5_Spying_DOM_Manipulation';

describe('Spying: DOM Manipulation', () => {
  let mockElement;
  let displaySpy;

  beforeEach(() => {
    // Mock DOM element with a proper getter and setter for `style.display`
    let displayValue = 'block'; 
    mockElement = {
      style: {
        get display() {
          return displayValue;
        },
        set display(value) {
          displayValue = value;
        },
      },
    };

    // Spy on the `display` property
    displaySpy = spyOnProperty(mockElement.style, 'display', 'set').and.callThrough(); 
  });

  it('should hide the element when it is initially visible', () => {
    toggleVisibility(mockElement);

    expect(mockElement.style.display).toBe('none');
    expect(displaySpy).toHaveBeenCalledWith('none');
    expect(displaySpy).toHaveBeenCalledTimes(1);
  });

  it('should show the element when it is initially hidden', () => {
    mockElement.style.display = 'none';
    expect(displaySpy).toHaveBeenCalledTimes(1);

    displaySpy.calls.reset();

    toggleVisibility(mockElement);

    expect(mockElement.style.display).toBe('block');
    expect(displaySpy).toHaveBeenCalledTimes(1);
    expect(displaySpy).toHaveBeenCalledWith('block');
  });

  it('should toggle visibility twice correctly', () => {
    // First toggle (visible -> hidden)
    toggleVisibility(mockElement);
    expect(mockElement.style.display).toBe('none');
    expect(displaySpy).toHaveBeenCalledTimes(1);
    expect(displaySpy).toHaveBeenCalledWith('none');

    // Reset spy for second toggle
    displaySpy.calls.reset();

    // Second toggle (hidden -> visible)
    toggleVisibility(mockElement);
    expect(mockElement.style.display).toBe('block');
    expect(displaySpy).toHaveBeenCalledTimes(1);
    expect(displaySpy).toHaveBeenCalledWith('block');
  });

  it('should not throw an error if the element has no display style set', () => {
    const mockElementWithoutStyle = {
      style: {},
    };

    expect(() => toggleVisibility(mockElementWithoutStyle)).not.toThrow();
    expect(mockElementWithoutStyle.style.display).toBe('none');
  });

  it("should track number of style changes", () => {
    // Toggle multiple times
    toggleVisibility(mockElement); 
    toggleVisibility(mockElement); 
    toggleVisibility(mockElement); 

    // Verify the spy was called exactly 3 times
    expect(displaySpy).toHaveBeenCalledTimes(3);
  });
});