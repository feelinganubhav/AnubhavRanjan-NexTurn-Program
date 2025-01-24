import { fetchAndDisplayUser } from '../src/a6_Bonus_Challenge';

describe('Bonus Challenge: Integrate All Concepts', () => {
  let mockApiService;
  let mockElement;
  let textContentSpy;

  beforeEach(() => {
    // Mock the DOM element
    let textContentValue = '';
    mockElement = {
      get textContent() {
        return textContentValue;
      },
      set textContent(value) {
        textContentValue = value;
      },
    };

    // Spy on the textContent property
    textContentSpy = spyOnProperty(mockElement, 'textContent', 'set').and.callThrough();

    // Mock the apiService
    mockApiService = {
      getUser: jasmine.createSpy('getUser'),
    };
  });

  it('should display user name when fetch is successful', async () => {
    const userId = 1;
    const user = { name: 'Anubhav' };
    mockApiService.getUser.and.resolveTo(user);

    await fetchAndDisplayUser(mockApiService, userId, mockElement);

    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(textContentSpy).toHaveBeenCalledTimes(1);
    expect(mockElement.textContent).toBe('Hello, Anubhav');
  });

  it('should display error message when user data is invalid', async () => {
    const userId = 2;
    const user = { name: '' };
    mockApiService.getUser.and.resolveTo(user);

    await fetchAndDisplayUser(mockApiService, userId, mockElement);

    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(textContentSpy).toHaveBeenCalledTimes(1);
    expect(mockElement.textContent).toBe('Invalid user data');
  });

  it('should display error message when fetch fails', async () => {
    const userId = 3;
    const errorMessage = 'User not found';
    mockApiService.getUser.and.rejectWith(new Error(errorMessage));

    await fetchAndDisplayUser(mockApiService, userId, mockElement);

    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(textContentSpy).toHaveBeenCalledTimes(1);
    expect(mockElement.textContent).toBe(errorMessage);
  });

  it('should handle network timeout', async () => {
    const userId = 4;
    mockApiService.getUser.and.callFake(
      () =>
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error('Network timeout')), 100)
        )
    );

    await fetchAndDisplayUser(mockApiService, userId, mockElement);

    expect(mockApiService.getUser).toHaveBeenCalledWith(userId);
    expect(textContentSpy).toHaveBeenCalledTimes(1);
    expect(mockElement.textContent).toBe('Network timeout');
  });

  it('should handle multiple consecutive calls', async () => {
    const userId = 5;
    mockApiService.getUser
      .and.resolveTo({ name: 'Anubhav' })
      .and.rejectWith(new Error('API Error'))
      .and.resolveTo({ name: 'Jonty' });

    await fetchAndDisplayUser(mockApiService, userId, mockElement);
    await fetchAndDisplayUser(mockApiService, userId, mockElement);
    await fetchAndDisplayUser(mockApiService, userId, mockElement);

    expect(mockApiService.getUser).toHaveBeenCalledTimes(3);
    expect(textContentSpy).toHaveBeenCalledTimes(3);
    expect(mockElement.textContent).toBe('Hello, Jonty');
  });
});
