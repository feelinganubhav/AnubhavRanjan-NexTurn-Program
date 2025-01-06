from management import BookManager
from management import CustomerManager
from management import SalesManager

def main_menu():
    print("\nWelcome to BookMart!")
    print("1. Book Management")
    print("2. Customer Management")
    print("3. Sales Management")
    print("4. Exit")

def book_management_menu(book_manager):
    while True:
        print("\nBook Management")
        print("1. Add a New Book")
        print("2. View All Books")
        print("3. Search for a Book")
        print("4. Update Book Stock")
        print("5. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            title = input("Enter book title: ")
            author = input("Enter book author: ")
            try:
                price = float(input("Enter book price: "))
                quantity = int(input("Enter book quantity: "))
                book_manager.add_book(title, author, price, quantity)
            except ValueError:
                print("Invalid input. Please enter numeric values for price and quantity.")

        elif choice == '2':
            book_manager.view_books()

        elif choice == '3':
            keyword = input("Enter title or author to search: ")
            book_manager.search_book(keyword)

        elif choice == '4':
            try:
                book_id = int(input("Enter book ID: "))
                quantity = int(input("Enter quantity to add or remove: "))
                book_manager.update_book_stock(book_id, quantity)
            except ValueError:
                print("Invalid input. Please enter numeric values.")

        elif choice == '5':
            break

        else:
            print("Invalid choice. Please try again.")

def customer_management_menu(customer_manager):
    while True:
        print("\nCustomer Management")
        print("1. Add a New Customer")
        print("2. View All Customers")
        print("3. Search for a Customer")
        print("4. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            name = input("Enter customer name: ")
            email = input("Enter customer email: ")
            phone = input("Enter customer phone number: ")
            customer_manager.add_customer(name, email, phone)

        elif choice == '2':
            customer_manager.view_customers()
        
        elif choice == '3':
            keyword = input("Enter Name or Email to search: ")
            customer_manager.search_customer(keyword)

        elif choice == '4':
            break

        else:
            print("Invalid choice. Please try again.")

def sales_management_menu(sales_manager, customer_manager):
    while True:
        print("\nSales Management")
        print("1. Sell a Book")
        print("2. View Sales Records")
        print("3. View Sales Summary")
        print("4. Back to Main Menu")
        choice = input("Enter your choice: ")

        if choice == '1':
            try:
                search_choice = input("Search customer by (1) ID, (2) Name, or (3) Email: ").strip()
                
                if search_choice == '1':
                    customer_id = int(input("Enter customer ID: "))
                    customer = customer_manager.get_customer_by_id(customer_id)
                elif search_choice == '2':
                    customer_name = input("Enter customer name: ").strip()
                    customer = customer_manager.get_customer_by_name(customer_name)
                elif search_choice == '3':
                    customer_email = input("Enter customer email: ").strip()
                    customer = customer_manager.get_customer_by_email(customer_email)
                else:
                    print("Invalid choice. Please try again.")
                    continue

                book_id = int(input("Enter book ID: "))
                quantity = int(input("Enter quantity: "))
                sales_manager.sell_book(customer.customer_id, book_id, quantity)
            except ValueError:
                print("Invalid input. Please Enter the Inputs values Accordingly...")
            except Exception as e:
                print(f"Error: {e}")

        elif choice == '2':
            sales_manager.view_sales()

        elif choice == '3':
            sales_manager.get_sales_summary()

        elif choice == '4':
            break

        else:
            print("Invalid choice. Please try again.")

def main():
    book_manager = BookManager()
    customer_manager = CustomerManager()
    sales_manager = SalesManager(book_manager, customer_manager)

    while True:
        main_menu()
        choice = input("Enter your choice: ")

        if choice == '1':
            book_management_menu(book_manager)

        elif choice == '2':
            customer_management_menu(customer_manager)

        elif choice == '3':
            sales_management_menu(sales_manager, customer_manager)

        elif choice == '4':
            print("Exiting BookMart Management System. Goodbye!")
            break

        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main()
