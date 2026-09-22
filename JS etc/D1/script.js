let task = [];

function addTask() {
    const taskinput = document.getElementById("todo-input");
    const dateinput = document.getElementById("date-input");
    const categoryinput = document.getElementById("category-input");

    if (taskinput.value === "" || dateinput.value === "" || categoryinput.value === "") {
        alert("Please enter a task, date, and category.");
        return;
    }

    task.push({
        task: taskinput.value,
        date: dateinput.value,
        category: categoryinput.value,
        completed: false 
    });

    saveTasksToLocalStorage();

    console.log("Task Added:", taskinput.value, "on", dateinput.value, "Category:", categoryinput.value);
    console.log("All Tasks:", [...task]);

    renderTasks();

    taskinput.value = "";
    dateinput.value = "";
    categoryinput.value = "";
}

function renderTasks() {
    const taskList = document.getElementById("todo-list");
    const container = document.getElementById("todo-container");
    const deleteBtn = document.getElementById("delete-all-btn"); 
    const filterSelect = document.getElementById("filter-category");
    taskList.innerHTML = "";

    if (task.length === 0) {
        deleteBtn.disabled = true;
        deleteBtn.classList.add("opacity-50", "cursor-not-allowed");
        taskList.innerHTML = `
            <div class="empty-message">
                <p>Your task will appear here</p>
            </div>
    `;


        if (filterSelect) {
            filterSelect.innerHTML = `<option value="all">All Categories</option>`;
        }
        return;
    } else {
        deleteBtn.disabled = false;
        deleteBtn.classList.remove("opacity-50", "cursor-not-allowed");
    }

    const selectedCategory = filterSelect?.value || "all";

    const filteredTasks = selectedCategory === "all"
        ? task
        : task.filter(t => t.category === selectedCategory);

    filteredTasks.forEach((t, index) => {
        const isChecked = t.completed ? "checked" : "";
        const taskClass = t.completed ? "line-through text-gray-400" : "";

        taskList.innerHTML += `
            <li class="todo-item flex justify-between items-center p-4 border-b border-gray-300">
                <div class="flex items-center gap-2">
                    <input type="checkbox" onchange="toggleComplete(${index})" ${isChecked} />
                    <div>
                        <p class="${taskClass} font-semibold">${t.task}</p>
                        <p class="${taskClass} text-sm text-gray-500">${t.date} | ${t.category}</p>
                    </div>
                </div>
                <div class="flex gap-2">
                    <button onclick="editTask(${index})" class="px-3 py-1 bg-green-500 text-white rounded">Edit</button>
                    <button onclick="deleteTask(${index})" class="px-3 py-1 bg-red-500 text-white rounded">Delete</button>
                </div>
            </li>
        `;
    });

    if (filterSelect) {
        const uniqueCategories = [...new Set(task.map(t => t.category))];
        filterSelect.innerHTML = `<option value="all">All Categories</option>`;
        uniqueCategories.forEach(cat => {
            const selected = cat === selectedCategory ? "selected" : "";
            filterSelect.innerHTML += `<option value="${cat}" ${selected}>${cat}</option>`;
        });
    }
}

function toggleComplete(index) {
    const filterSelect = document.getElementById("filter-category");
    const selectedCategory = filterSelect?.value || "all";

    const filteredTasks = selectedCategory === "all"
        ? task
        : task.filter(t => t.category === selectedCategory);

    const taskToToggle = filteredTasks[index];

    const originalIndex = task.findIndex(t =>
        t.task === taskToToggle.task &&
        t.date === taskToToggle.date &&
        t.category === taskToToggle.category
    );

    if (originalIndex > -1) {
        task[originalIndex].completed = !task[originalIndex].completed;
        saveTasksToLocalStorage();
        renderTasks();
    }
}

function deleteAllTasks() {
    if (confirm("Are you sure you want to delete all tasks?")) {
        task = [];
        saveTasksToLocalStorage();
        renderTasks();
        console.log("All tasks deleted.");
    }
}

function deleteTask(index) {
    if (confirm("Delete this task?")) {
        const filterSelect = document.getElementById("filter-category");
        const selectedCategory = filterSelect?.value || "all";

        const filteredTasks = selectedCategory === "all"
            ? task
            : task.filter(t => t.category === selectedCategory);

        const taskToDelete = filteredTasks[index];

        const originalIndex = task.findIndex(t => 
            t.task === taskToDelete.task &&
            t.date === taskToDelete.date &&
            t.category === taskToDelete.category
        );

        if (originalIndex > -1) {
            task.splice(originalIndex, 1);
            saveTasksToLocalStorage();
            renderTasks();
        }
    }
}

function editTask(index) {
    const filterSelect = document.getElementById("filter-category");
    const selectedCategory = filterSelect?.value || "all";

    const filteredTasks = selectedCategory === "all"
        ? task
        : task.filter(t => t.category === selectedCategory);

    const taskToEdit = filteredTasks[index];

    const originalIndex = task.findIndex(t =>
        t.task === taskToEdit.task &&
        t.date === taskToEdit.date &&
        t.category === taskToEdit.category
    );

    document.getElementById("todo-input").value = taskToEdit.task;
    document.getElementById("date-input").value = taskToEdit.date;
    document.getElementById("category-input").value = taskToEdit.category;

    const addButton = document.querySelector("form button[type='button']");
    addButton.textContent = "Save";
    addButton.onclick = function () {
        saveTask(originalIndex);
    };
}

function saveTask(index) {
    const taskinput = document.getElementById("todo-input");
    const dateinput = document.getElementById("date-input");
    const categoryinput = document.getElementById("category-input");

    if (taskinput.value === "" || dateinput.value === "" || categoryinput.value === "") {
        alert("Please enter a task, date, and category.");
        return;
    }

    task[index] = {
        task: taskinput.value,
        date: dateinput.value,
        category: categoryinput.value,
        completed: task[index].completed
    };

    saveTasksToLocalStorage();

    const addButton = document.querySelector("form button[type='button']");
    addButton.textContent = "+";
    addButton.onclick = addTask;

    renderTasks();

    taskinput.value = "";
    dateinput.value = "";
    categoryinput.value = "";
}

function saveTasksToLocalStorage() {
    localStorage.setItem("tasks", JSON.stringify(task));
}


document.addEventListener("DOMContentLoaded", function () {
    const storedTasks = localStorage.getItem("tasks");
    if (storedTasks) {
        task = JSON.parse(storedTasks);
    }
    renderTasks();
});

