class Customer:
    def __init__(self, customer_id, name, email, phone):
        if not name or not email or not phone:
            raise ValueError("Name, email, and phone cannot be empty.")

        if not ("@" in email and "." in email):
            raise ValueError("Invalid email format.")

        if not phone.isdigit() or len(phone) < 10:
            raise ValueError("Invalid phone number. It must contain at least 10 digits.")

        self.customer_id = customer_id
        self.name = name
        self.email = email
        self.phone = phone

    def __str__(self):
        return f"Customer({self.customer_id}): {self.name}, Email: {self.email}, Phone: {self.phone}"
