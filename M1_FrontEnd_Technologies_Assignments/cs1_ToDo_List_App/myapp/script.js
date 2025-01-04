const taskList = document.getElementById('taskList');
const pendingCount = document.getElementById('pendingCount');

const taskN = document.getElementById('taskName');
taskN.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        addTask();
    }
});

function loadTasks() {
    const tasks = JSON.parse(localStorage.getItem('tasks')) || [];
    tasks.forEach(task => createTaskElement(task.name, task.completed));
    updatePendingCount();
}

function saveTasks() {
    const tasks = Array.from(taskList.children).map(task => ({
        name: task.querySelector('span').textContent,
        completed: task.classList.contains('completed')
    }));
    localStorage.setItem('tasks', JSON.stringify(tasks));
}

function createTaskElement(name, completed = false) {
    const li = document.createElement('li');
    li.className = `task${completed ? ' completed' : ''}`;

    const span = document.createElement('span');
    span.textContent = name;
    span.addEventListener('click', () => toggleTask(li));
    span.addEventListener('contextmenu', () => deleteTask(li));


    const buttonsDiv = document.createElement('div');
    buttonsDiv.className = 'task-buttons';

    const upBtn = document.createElement('button');
    // upBtn.textContent = '⬆';
    // upBtn.className = 'up';
    upBtn.innerHTML = '<i class="bi bi-arrow-up"></i>';
    upBtn.className = 'btn btn-sm btn-outline-secondary me-2';

    upBtn.addEventListener('click', () => moveTaskUp(li));

    const downBtn = document.createElement('button');
    // downBtn.textContent = '⬇';
    // downBtn.className = 'down';
    downBtn.innerHTML = '<i class="bi bi-arrow-down"></i>';
    downBtn.className = 'btn btn-sm btn-outline-secondary me-2';

    downBtn.addEventListener('click', () => moveTaskDown(li));

    const editBtn = document.createElement('button');
    // editBtn.textContent = '✎';
    // editBtn.className = 'edit';
    editBtn.innerHTML = '<i class="bi bi-pencil"></i>';
    editBtn.className = 'btn btn-sm btn-outline-primary me-2';

    editBtn.addEventListener('click', () => editTask(li));

    const deleteBtn = document.createElement('button');
    // deleteBtn.textContent = 'Del';
    deleteBtn.innerHTML = '<i class="bi bi-trash"></i>';
    deleteBtn.className = 'btn btn-sm btn-outline-danger me-2';

    // deleteBtn.className = 'delete';
    deleteBtn.addEventListener('click', () => deleteTask(li));

    buttonsDiv.appendChild(upBtn);
    buttonsDiv.appendChild(downBtn);
    buttonsDiv.appendChild(editBtn);
    buttonsDiv.appendChild(deleteBtn);

    li.appendChild(span);
    li.appendChild(buttonsDiv);
    taskList.appendChild(li);

    saveTasks();
}

function addTask() {
    const taskName = document.getElementById('taskName').value.trim();
    if (taskName) {
        createTaskElement(taskName);
        document.getElementById('taskName').value = '';
        updatePendingCount();
    }
}

function toggleTask(taskElement) {
    taskElement.classList.toggle('completed');
    saveTasks();
    updatePendingCount();
}

function editTask(taskElement) {
    const span = taskElement.querySelector('span');
    const newTaskName = prompt('Edit task:', span.textContent);
    if (newTaskName !== null && newTaskName.trim() !== '') {
        span.textContent = newTaskName.trim();
        saveTasks();
    }
}

function deleteTask(taskElement) {
    taskList.removeChild(taskElement);
    saveTasks();
    updatePendingCount();
}

function moveTaskUp(taskElement) {
    const prevSibling = taskElement.previousElementSibling;
    if (prevSibling) {
        taskList.insertBefore(taskElement, prevSibling);
        saveTasks();
    }
}

function moveTaskDown(taskElement) {
    const nextSibling = taskElement.nextElementSibling;
    if (nextSibling) {
        taskList.insertBefore(nextSibling, taskElement);
        saveTasks();
    }
}

function clearAllTasks() {
    if (confirm('Are you sure you want to clear all tasks?')) {
        taskList.innerHTML = ''; 
        saveTasks(); 
        updatePendingCount(); 
    }
}


function updatePendingCount() {
    const pendingTasks = Array.from(taskList.children).filter(task => !task.classList.contains('completed')).length;
    pendingCount.textContent = `${pendingTasks} task${pendingTasks !== 1 ? 's' : ''} pending`;
}

document.addEventListener('DOMContentLoaded', () => {
    loadTasks();

    // Initialize Sortable.js
    new Sortable(taskList, {
        animation: 150,
        onEnd: saveTasks
    });
});