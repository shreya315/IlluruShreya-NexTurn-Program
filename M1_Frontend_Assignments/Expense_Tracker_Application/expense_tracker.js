let amountInput = document.getElementById("amount");
let descriptionInput = document.getElementById("description");
let categorySelect = document.getElementById("category");
let addExpenseBtn = document.getElementById("addExpense");
let expenseList = document.getElementById("expenseList");
let categorySummary = document.getElementById("categorySummary");

window.onload = function () {
    loadExpenses();
};

addExpenseBtn.onclick = function () {
    let amount = amountInput.value;  
    let description = descriptionInput.value;  
    let category = categorySelect.value;  

    if (amount === "" || description === "") {
        alert("Please fill both amount and description.");
    } else {
        let expense = {
            amount: amount,
            description: description,
            category: category
        };

        addExpenseToDOM(expense);
        saveExpenseToLocalStorage(expense);
        updateCategorySummary();
        amountInput.value = "";
        descriptionInput.value = "";
    }
};

function addExpenseToDOM(expense) {
    let row = document.createElement("tr");

    row.innerHTML = `
        <td>$${expense.amount}</td>
        <td>${expense.description}</td>
        <td>${expense.category}</td>
        <td><button class="delete-btn">Delete</button></td>
    `;

    row.querySelector(".delete-btn").onclick = function () {
        row.remove();
        deleteExpenseFromLocalStorage(expense);
        updateCategorySummary();
    };
    expenseList.appendChild(row);
}

function saveExpenseToLocalStorage(expense) {
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
    expenses.push(expense);
    localStorage.setItem("expenses", JSON.stringify(expenses));
}

function deleteExpenseFromLocalStorage(expense) {
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
    expenses = expenses.filter(e => e.amount !== expense.amount || e.description !== expense.description || e.category !== expense.category);
    localStorage.setItem("expenses", JSON.stringify(expenses));
}

function loadExpenses() {
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
    expenses.forEach(expense => {
        addExpenseToDOM(expense);
    });
    updateCategorySummary();
}

function updateCategorySummary() {
    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];
    let categoryTotals = {};
    for (let i = 0; i < expenses.length; i++) {
        let category = expenses[i].category;
        if (!categoryTotals[category]) {
            categoryTotals[category] = 0;
        }
        categoryTotals[category] += parseFloat(expenses[i].amount);
    }
    categorySummary.innerHTML = ""; 
    for (let category in categoryTotals) {
        let listItem = document.createElement("li");
        listItem.textContent = `${category}: $${categoryTotals[category].toFixed(2)}`;
        categorySummary.appendChild(listItem);
    }
}
