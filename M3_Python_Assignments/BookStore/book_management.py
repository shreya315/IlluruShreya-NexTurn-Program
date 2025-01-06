class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_book(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

def add_book(books, title, author, price, quantity):
    book = Book(title, author, price, quantity)
    books.append(book)
    print(f"{title}' has added to the inventory successfully!")

def view_book(books):
    for book in books:
        print(book.display_book())

def search_book(books, search_term):
    found_book = [book for book in books if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()]
    if not found_book:
        print("No books found!")
    else:
        for book in found_book:
            print(book.display_book())
