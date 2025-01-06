import book_management
import customer_management
import sales_management

def menu():
    books = []
    customers = []
    
    while True:
        print("Welcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            bookManagement(books)
        elif choice == "2":
            customerManagement(customers)
        elif choice == "3":
            salesManagement(customers, books)
        elif choice == "4":
            print("Exit.")
            break
        else:
            print("Invalid choice")

def bookManagement(books):
    while True:
        print("Book Management")
        print("1. Add Book")
        print("2. View Books")
        print("3. Search Book")
        print("4. Go to Home")
        choice = input("Enter your choice: ")

        if choice == "1":
            title = input("Enter book name: ")
            author = input("Enter author name: ")
            try:
                price = float(input("Enter price of the book: "))
                quantity = int(input("Enter book quantity: "))
                if price <= 0 or quantity <= 0:
                    print("Price and quantity must be positive.")
                else:
                    book_management.add_book(books, title, author, price, quantity)
            except ValueError:
                print("Invalid input! Price must be a number and quantity must be an integer.")
        elif choice == "2":
            book_management.view_books(books)
        elif choice == "3":
            search_term = input("Enter title of the book or author of the book: ")
            book_management.search_book(books, search_term)
        elif choice == "4":
            break
        else:
            print("Invalid choice.")

def customerManagement(customers):
    while True:
        print("Customer Management")
        print("1. Add Customer")
        print("2. View Customers")
        print("3. Go to Home")
        choice = input("Enter your choice: ")

        if choice == "1":
            name = input("Enter name of the customer: ")
            email = input("Enter customer email: ")
            phone = input("Enter customer phone no: ")
            customer_management.add_customer(customers, name, email, phone)
        elif choice == "2":
            customer_management.view_customers(customers)
        elif choice == "3":
            break
        else:
            print("Invalid choice.")

def salesManagement(customers, books):
    while True:
        print("Sales Management")
        print("1. Sell Book")
        print("2. View Sales")
        print("3. Go to Home")
        choice = input("Enter your choice: ")

        if choice == "1":
            customer_name = input("Enter name of the customer: ")
            book_title = input("Enter book name: ")
            try:
                quantity = int(input("Enter quantity: "))
                if quantity <= 0:
                    print("Quantity must be a positive integer.")
                else:
                    sales_management.sell_book(customers, books, customer_name, book_title, quantity)
            except ValueError:
                print("Invalid input! Quantity must be an integer.")
        elif choice == "2":
            sales_management.view_sales()
        elif choice == "3":
            break
        else:
            print("Invalid choice.")

# Run the program
if __name__ == "__main__":
    menu()
