let input_task = document.getElementById("input");
let addtaskbtn = document.getElementById("addtask");
let list = document.getElementById("toDo");

addtaskbtn.onclick = function () {
    if (input_task.value === "") {
        alert("Add a task to keep your day on track.");
    } else {
        addTaskToDOM(input_task.value, false);
        saveTasksToLocalStorage();
        updatePendingTaskCount();
    }
    input_task.value = ""; 
};

window.onload = function () {
    loadTasksFromLocalStorage();
    updatePendingTaskCount();
};

function addTaskToDOM(taskText, isCompleted) {
    let task = document.createElement("li");
    task.innerHTML = taskText;

    if (isCompleted) {
        task.classList.add("completed");
    }

    let del = document.createElement("span");
    del.innerHTML = "\u00D7"; 
    task.appendChild(del);

    del.onclick = function (event) {
        event.stopPropagation(); 
        task.remove();
        saveTasksToLocalStorage();
        updatePendingTaskCount();
    };

    task.onclick = function () {
        task.classList.toggle("completed");
        saveTasksToLocalStorage();
        updatePendingTaskCount();
    };

    task.ondblclick = function () {
        if (!task.classList.contains("completed")) {
            let newTaskText = prompt("Edit your task:", taskText);
            if (newTaskText) {
                task.innerHTML = newTaskText; 
                task.appendChild(del); 
                saveTasksToLocalStorage();
            }
        } else {
            alert("Completed tasks cannot be edited.");
        }
    };
t
    list.appendChild(task);
}

function saveTasksToLocalStorage() {
    let tasks = [];
    let taskItems = document.querySelectorAll("#toDo li");
    taskItems.forEach((task) => {
        tasks.push({
            text: task.textContent.replace("×", "").trim(),
            completed: task.classList.contains("completed"),
        });
    });
    localStorage.setItem("tasks", JSON.stringify(tasks));
}

function loadTasksFromLocalStorage() {
    let tasks = JSON.parse(localStorage.getItem("tasks")) || [];
    tasks.forEach((task) => {
        addTaskToDOM(task.text, task.completed);
    });
}

function updatePendingTaskCount() {
    let taskItems = document.querySelectorAll("#toDo li");
    let pendingCount = 0;
    taskItems.forEach((task) => {
        if (!task.classList.contains("completed")) {
            pendingCount++;
        }
    });

    let taskCountElement = document.getElementById("taskCount");
    if (!taskCountElement) {
        taskCountElement = document.createElement("div");
        taskCountElement.id = "taskCount";
        document.querySelector(".app").appendChild(taskCountElement);
    }
    taskCountElement.innerHTML = `Pending tasks: ${pendingCount}`;
}
