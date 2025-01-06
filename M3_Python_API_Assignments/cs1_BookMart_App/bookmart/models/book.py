class Book:
    def __init__(self, book_id, title, author, price, quantity):
        if not title or not author:
            raise ValueError("Invalid book details: Title or Author cannot be empty...")
        if price <= 0:
            raise ValueError("price must be positive...")
        if quantity < 0:
            raise ValueError("quantity cannot be negative...")
        
        self.book_id = book_id
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def update_stock(self, quantity):
        if quantity < 0 and abs(quantity) > self.quantity:
            raise ValueError(f"Cannot reduce stock below 0 for {self.title}. Current stock: {self.quantity}.")
        self.quantity += quantity

    def __str__(self):
        return (f"Book({self.book_id}): '{self.title}' by {self.author}, "
                f"Price: Rs.{self.price:.2f}, Quantity: {self.quantity})")
