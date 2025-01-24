
import { sendNotification } from '../src/a4_Mocking_Notification_Service';

describe('Mocking: Notification Service', () => {
  let mockNotificationService;

  beforeEach(() => {
    mockNotificationService = {
      send: jasmine.createSpy('send'),
    };
  });

  it('should return "Notification Sent" when notification is successfully sent', () => {
    mockNotificationService.send.and.returnValue(true);

    const result = sendNotification(mockNotificationService, 'Hello, World!');

    expect(result).toBe('Notification Sent');
    expect(mockNotificationService.send).toHaveBeenCalledWith('Hello, World!');
  });

  it('should return "Failed to Send" when notification fails to send', () => {
    mockNotificationService.send.and.returnValue(false);

    const result = sendNotification(mockNotificationService, 'Error message');

    expect(result).toBe('Failed to Send');
    expect(mockNotificationService.send).toHaveBeenCalledWith('Error message');
  });

  it('should not send notification if the message is empty', () => {
    mockNotificationService.send.and.returnValue(true);

    const result = sendNotification(mockNotificationService, '');

    expect(result).toBe('Failed to Send: Message is Empty...'); 
    expect(mockNotificationService.send).not.toHaveBeenCalledWith('');
  });

  it('should handle a null notification service gracefully', () => {
    expect(() => sendNotification(null, 'Hello!')).toThrow();
  });
});
