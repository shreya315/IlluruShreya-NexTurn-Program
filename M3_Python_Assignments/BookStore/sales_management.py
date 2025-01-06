from customer_management import Customer
from book_management import Book

class Transaction(Customer):
    def __init__(self, name, email, phone, book_name, quantity_sold):
        super().__init__(name, email, phone)
        self.book_name = book_name
        self.quantity_sold = quantity_sold

    def transaction_log(self):
        return f"Customer: {self.name}, Book: {self.book_name}, Quantity Sold: {self.quantity_sold}"

sales_records = []

def sell_book(customers, books, customer_name, book_name, quantity):
    customer = next((customer for customer in customers if customer.name == customer_name), None)
    if not customer:
        print("Customer not found.")
        return

    book = next((book for book in books if book.name == book_name), None)
    if not book:
        print("Book not found.")
        return

    if book.quantity < quantity:
        print(f"Only {book.quantity} copies of {book_name} are available. So the book can't be sold.")
        return

    book.quantity -= quantity
    transaction = Transaction(customer_name, customer.email, customer.phone, book_name, quantity)
    sales_records.append(transaction)
    print(f"The book {book_name} has been successfully")

def view_sales():
    for transaction in sales_records:
        print(transaction.display_transaction())
