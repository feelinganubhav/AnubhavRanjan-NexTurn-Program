from .customer import Customer
from datetime import datetime

class Transaction(Customer):
    def __init__(self, customer, book_title, quantity_sold, total_price, transaction_date = datetime.now()):
        super().__init__(customer.customer_id, customer.name, customer.email, customer.phone)
        if quantity_sold <= 0 or total_price <= 0:
            raise ValueError("Quantity sold and total price must be positive.")

        self.book_title = book_title
        self.quantity_sold = quantity_sold
        self.total_price = total_price
        self.transaction_date = transaction_date

    def __str__(self):
        return (f"Date: {self.transaction_date.strftime('%Y-%m-%d %H:%M:%S')}\n"
                f"Transaction(Customer ID: {self.customer_id}, Name: {self.name}, "
                f"Book: '{self.book_title}', Quantity Sold: {self.quantity_sold}, "
                f"Total Price: ${self.total_price:.2f})")
