import json
from models import Transaction
from models import Book
from models import Customer
from datetime import datetime

SALES_FILE = 'bookmart/data/sales.json'

class SalesManager:
    def __init__(self, book_manager, customer_manager):
        self.sales = self.load_sales()
        self.book_manager = book_manager
        self.customer_manager = customer_manager

    def load_sales(self):
        """Load sales records from the JSON file."""
        try:
            with open(SALES_FILE, 'r') as file:
                sales_data = json.load(file)
                # print(sales_data)
                # customer = Customer()
                # for sale_data in sales_data:
                #     customer.customer_id=sale_data['customer_id']
                #     customer.name = sale_data['name']
                #     customer.email = sale_data['email']
                #     customer.phone = sale_data['phone']
                return [
                    Transaction(
                        customer=Customer(sale_data['customer_id'], sale_data['name'], sale_data['email'], sale_data['phone']),
                        book_title=sale_data['book_title'],
                        quantity_sold=sale_data['quantity_sold'],
                        total_price = sale_data['total_price'],
                        transaction_date=datetime.fromisoformat(sale_data['transaction_date']) if 'transaction_date' in sale_data else None
                    )
                    for sale_data in sales_data
                ]
        except FileNotFoundError:
            return []
        except Exception as e:
            print(f"Error loading sales: {e}")
            return []

    def save_sales(self):
        """Save sales records to the JSON file."""
        with open(SALES_FILE, 'w') as file:
            sales_data = [
                {
                    **sale.__dict__,
                    'transaction_date': sale.transaction_date.isoformat() if isinstance(sale.transaction_date, datetime) else sale.transaction_date
                }
                for sale in self.sales
            ]
            json.dump(sales_data, file, indent=4)

    def sell_book(self, customer_id, book_id, quantity):
        """Sell a book to a customer."""
        try:
            customer = self.customer_manager.get_customer_by_id(customer_id)
            book = self.book_manager.get_book_by_id(book_id)

            if book.quantity < quantity:
                raise ValueError(f"Not enough stock for '{book.title}'. Only {book.quantity} left.")

            # book.update_stock(-quantity)
            self.book_manager.update_book_stock(book_id, -quantity)
            transaction = Transaction(customer, book.title, quantity, book.price * quantity)
            self.sales.append(transaction)
            self.save_sales()
            self.book_manager.save_books()

            print(f"Sale successful! Book '{book.title}' sold to {customer.name}.")
        except ValueError as e:
            print(f"Error Processing Sale: {e}")

    def view_sales(self):
        """Display all sales records."""
        if not self.sales:
            print("No sales records available.")
        else:
            print("\nSales Records:\n")
            for sale in self.sales:
                print(sale)

    def get_sales_summary(self):
        """Provide a summary of total sales and revenue."""
        total_books_sold = sum(sale.quantity_sold for sale in self.sales)
        total_revenue = sum(sale.total_price for sale in self.sales)
        print(f"Total Books Sold: {total_books_sold}")
        print(f"Total Revenue: Rs.{total_revenue:.2f}")
