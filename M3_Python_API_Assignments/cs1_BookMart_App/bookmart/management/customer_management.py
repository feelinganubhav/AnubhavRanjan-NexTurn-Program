import json
from models import Customer

CUSTOMERS_FILE = 'bookmart/data/customers.json'

class CustomerManager:
    def __init__(self):
        self.customers = self.load_customers()

    def load_customers(self):
        """Load customers from the JSON file."""
        try:
            with open(CUSTOMERS_FILE, 'r') as file:
                customers_data = json.load(file)
                return [Customer(**data) for data in customers_data]
        except FileNotFoundError:
            return []
        except Exception as e:
            print(f"Error loading customers: {e}")
            return []

    def save_customers(self):
        """Save customers to the JSON file."""
        with open(CUSTOMERS_FILE, 'w') as file:
            json.dump([customer.__dict__ for customer in self.customers], file, indent=4)

    def add_customer(self, name, email, phone):
        """Add a new customer."""
        customer_id = len(self.customers) + 1
        try:
            if any(customer.email == email for customer in self.customers):
                raise ValueError(f"Customer with email '{email}' already exists.")

            new_customer = Customer(customer_id, name, email, phone)
            self.customers.append(new_customer)
            self.save_customers()
            print(f"Customer '{name}' added successfully.")
        except ValueError as e:
            print(f"Error: {e}")

    def view_customers(self):
        """Display all customers."""
        if not self.customers:
            print("No customers available.")
        else:
            print("\nRegistered Customers:\n")
            for customer in self.customers:
                print(customer)

    def search_customer(self, keyword):
        """Search for a customer by name or email."""
        results = [customer for customer in self.customers if keyword.lower() in customer.name.lower() or keyword.lower() in customer.email.lower()]
        if results:
            for customer in results:
                print(customer)
        else:
            print(f"No customers found for the search keyword: {keyword}")

    def get_customer_by_id(self, customer_id):
        """Retrieve a customer by ID."""
        customer = next((c for c in self.customers if c.customer_id == customer_id), None)
        if not customer:
            raise ValueError(f"Customer with ID {customer_id} not found.")
        return customer

    def get_customer_by_name(self, name):
        """Retrieve a customer by name."""
        customer = next((c for c in self.customers if c.name.lower() == name.lower()), None)
        if not customer:
            raise ValueError(f"Customer with name '{name}' not found.")
        return customer

    def get_customer_by_email(self, email):
        """Retrieve a customer by email."""
        customer = next((c for c in self.customers if c.email.lower() == email.lower()), None)
        if not customer:
            raise ValueError(f"Customer with email '{email}' not found.")
        return customer

