import json
from models import Book

BOOKS_FILE = 'bookmart/data/books.json'

class BookManager:
    def __init__(self):
        self.books = self.load_books()

    def load_books(self):
        """Load books from the JSON file."""
        try:
            with open(BOOKS_FILE, 'r') as file:
                books_data = json.load(file)
                return [Book(**data) for data in books_data]
        except FileNotFoundError:
            return []
        except Exception as e:
            print(f"Error loading books: {e}")
            return []

    def save_books(self):
        """Save books to the JSON file."""
        with open(BOOKS_FILE, 'w') as file:
            json.dump([book.__dict__ for book in self.books], file, indent=4)

    def add_book(self, title, author, price, quantity):
        """Add a new book to the inventory."""
        book_id = len(self.books) + 1
        try:
            existing_book = self.get_book(title)
            if existing_book:
                existing_book.quantity += quantity
                print(f"Updated quantity for existing book: {title}")
                return
            
            new_book = Book(book_id, title, author, price, quantity)
            self.books.append(new_book)
            self.save_books()
            print(f"Book '{title}' added successfully.")
        except ValueError as e:
            print(f"Error in Adding Book: {e}")

    def view_books(self):
        """Display all books in the inventory."""
        if not self.books:
            print("No books available in the inventory.")
        else:
            print("\nAvailable Books:\n")
            for book in self.books:
                print(book)


    def get_book(self, keyword):
        """Get a book by title or author."""
        keyword = keyword.lower()
        for book in self.books:
            if keyword in book.title.lower() or keyword in book.author.lower():
                return book
        return None
    
    def search_book(self, keyword):
        """Search for a book by title or author."""
        results = [book for book in self.books if keyword.lower() in book.title.lower() or keyword.lower() in book.author.lower()]
        if results:
            for book in results:
                print(book)
        else:
            print(f"No books found for the search keyword: {keyword}")

    def get_book_by_id(self, book_id):
        """Retrieve a Book by ID."""
        book = next((b for b in self.books if b.book_id == book_id), None)
        if not book:
            raise ValueError(f"Book with ID {book_id} not found.")
        return book

    def update_book_stock(self, book_id, quantity):
        """Update the stock of a book."""
        try:
            book = self.get_book_by_id(book_id)
            book.update_stock(quantity)
            self.save_books()
            print(f"Stock updated successfully for book ID {book_id}.")
        except ValueError as e:
            print(f"Error: {e}")
