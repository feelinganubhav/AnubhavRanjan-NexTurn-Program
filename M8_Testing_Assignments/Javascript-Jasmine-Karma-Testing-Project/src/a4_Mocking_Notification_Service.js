
export function sendNotification(notificationService, message) {
    if(message === '')
    {
        return "Failed to Send: Message is Empty...";
    }
    const status = notificationService.send(message);
    return status ? "Notification Sent" : "Failed to Send";
  }
  