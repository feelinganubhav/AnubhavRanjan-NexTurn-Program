const expenseForm = document.getElementById('expenseForm');
const expenseTableBody = document.querySelector('#expenseTable tbody');
const summaryDiv = document.getElementById('summary');
const toggleDarkModeButton = document.getElementById('toggleDarkMode');
const chartContainer = document.getElementById('chartContainer');
const toggleChartTypeSwitch = document.getElementById('chartToggle');

let expenses = JSON.parse(localStorage.getItem('expenses')) || [];
let isDarkMode = false;
let chartType = 'pie';
let chartInstance = null;

const filterTypeDropdown = document.getElementById('filterType');
const categoryFilterSection = document.getElementById('categoryFilterSection');
const dateRangeFilterSection = document.getElementById('dateRangeFilterSection');
const applyCategoryFilterButton = document.getElementById('applyCategoryFilter');
const applyDateFilterButton = document.getElementById('applyDateFilter');

let categories = JSON.parse(localStorage.getItem('categories')) || {
    Food: '#ff6b6b',
    Travel: '#1e90ff',
    Shopping: '#ffc107',
};

function renderExpenses(filteredExpenses = expenses) {
    expenseTableBody.innerHTML = '';
    const categoryTotals = {};

    for (const category of Object.keys(categories)) {
        categoryTotals[category] = 0;
    }

    filteredExpenses.forEach((expense, index) => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${expense.date}</td>
            <td>$${expense.amount}</td>
            <td>${expense.description}</td>
                <td><span class="category-badge" style="background-color: ${categories[expense.category]};">${expense.category}</span></td>
            <td>
                <button class = 'btn btn-sm btn-outline-primary me-2' onclick="editExpense(${index})"><i class="bi bi-pencil"></i></button>
                <button class = 'btn btn-sm btn-outline-danger me-2' onclick="deleteExpense(${index})"><i class="bi bi-trash"></i></button>
            </td>
        `;
        expenseTableBody.appendChild(row);

        categoryTotals[expense.category] += parseFloat(expense.amount);
    });

    updateSummary(categoryTotals);
    renderChart(categoryTotals);
}

function renderCategoryOptions() {
    const categoryDropdown = document.getElementById('category');
    const filterByCategory = document.getElementById('categoryFilter');

    categoryDropdown.innerHTML = '<option value="">Select Category</option>';
    filterByCategory.innerHTML = '<option value="">All</option>';

    for (const [name, color] of Object.entries(categories)) {
        const option = document.createElement('option');
        option.value = name;
        option.textContent = name;
        categoryDropdown.appendChild(option);

        const filterOption = document.createElement('option');
        filterOption.value = name;
        filterOption.textContent = name;
        filterByCategory.appendChild(filterOption);
    }
}


function updateSummary(totals = {}) {
    // summaryDiv.innerHTML = `
    //     <p>Food: $${totals.Food.toFixed(2)}</p>
    //     <p>Travel: $${totals.Travel.toFixed(2)}</p>
    //     <p>Shopping: $${totals.Shopping.toFixed(2)}</p>
    // `;
    summaryDiv.innerHTML = '';

    for (const [category, color] of Object.entries(categories)) {
        const total = totals[category] || 0;
        const summaryRow = document.createElement('p');
        summaryRow.innerHTML = `<span style="color: ${color}; font-weight: bold;">${category}:</span> $${total.toFixed(2)}`;
        summaryDiv.appendChild(summaryRow);
    }
}

function renderChart(categoryTotals) {
    const ctx = document.getElementById('expenseChart').getContext('2d');
    const data = {
        labels: Object.keys(categories),
        datasets: [
            {
                data: Object.values(categoryTotals),
                backgroundColor: Object.values(categories),
                hoverOffset: 4,
            },
        ],
    };

    if (chartInstance) {
        chartInstance.destroy();
    }

    chartInstance = new Chart(ctx, {
        type: chartType,
        data,
        options: {
            responsive: true,
            plugins: {
                legend: {
                    position: 'bottom'
                },
                title: {
                    display: true,
                    text: 'Expenses by Category'
                }
            }
        }
    });

    
}

function toggleChartType() {
    chartType = chartType === 'pie' ? 'bar' : 'pie';
    renderExpenses();
}


function addNewCategory() {
    const categoryName = document.getElementById('categoryName').value.trim();
    const categoryColor = document.getElementById('categoryColor').value;

    if (!categoryName || categories[categoryName]) {
        alert('Category already exists or name is invalid.');
        return;
    }

    categories[categoryName] = categoryColor;

    // Update relevant sections
    renderCategoryOptions();
    renderExpenses();
    updateSummary({});

    // Reset and close modal
    document.getElementById('addCategoryForm').reset();
    const modal = bootstrap.Modal.getInstance(document.getElementById('addCategoryModal'));
    modal.hide();
}


function addExpense(event) {
    event.preventDefault();
    const amount = document.getElementById('amount').value;
    const description = document.getElementById('description').value;
    const date = document.getElementById('date').value;
    const category = document.getElementById('category').value;
    const index = document.getElementById('expenseIndex').value;

    if (index === '') {
        // Add new expense
        expenses.push({ date, amount, description, category });
    } else {
        // Update existing expense
        expenses[index] = { date, amount, description, category };

        // Reset editing state
        document.getElementById('expenseIndex').value = '';
        document.getElementById('submitButton').textContent = 'Add Expense';
    }
    localStorage.setItem('expenses', JSON.stringify(expenses));
    localStorage.setItem('categories', JSON.stringify(categories));

    renderExpenses();
    expenseForm.reset();
}

function editExpense(index) {
    const expense = expenses[index];

    document.getElementById('date').value = expense.date;
    document.getElementById('amount').value = expense.amount;
    document.getElementById('description').value = expense.description;
    document.getElementById('category').value = expense.category;

    document.getElementById('expenseIndex').value = index;

    document.getElementById('submitButton').textContent = 'Update Expense';
    
    renderExpenses();
}


function deleteExpense(index) {
    if (confirm('Are you sure you want to delete this expense?')) {
        expenses.splice(index, 1);
        localStorage.setItem('expenses', JSON.stringify(expenses));
        renderExpenses();
    }
}

function selectFilterType() {
    const filterType = filterTypeDropdown.value;

    categoryFilterSection.style.display = 'none';
    dateRangeFilterSection.style.display = 'none';

    if (filterType === 'category') {
        categoryFilterSection.style.display = 'block';
    } else if (filterType === 'dateRange') {
        dateRangeFilterSection.style.display = 'block';
    }
}

function filterByCategory() {
    const categoryFilter = document.getElementById('categoryFilter').value;
    const filteredExpenses = expenses.filter(
        (expense) => !categoryFilter || expense.category === categoryFilter
    );
    renderExpenses(filteredExpenses);
}

function filterByDate() {
    const startDate = document.getElementById('startDate').value;
    const endDate = document.getElementById('endDate').value;

    const filteredExpenses = expenses.filter((expense) => {
        const expenseDate = new Date(expense.date);
        return (
            (!startDate || expenseDate >= new Date(startDate)) &&
            (!endDate || expenseDate <= new Date(endDate))
        );
    });

    renderExpenses(filteredExpenses);
}

// Dark Mode Toggle
function toggleDarkMode() {
    document.body.classList.toggle('dark-mode');
    isDarkMode = !isDarkMode;
    toggleDarkModeButton.textContent = isDarkMode ? 'Light Mode' : 'Dark Mode';
}


expenseForm.addEventListener('submit', addExpense);
filterTypeDropdown.addEventListener('change', selectFilterType);
applyCategoryFilterButton.addEventListener('click', filterByCategory);
applyDateFilterButton.addEventListener('click', filterByDate);
toggleDarkModeButton.addEventListener('click', toggleDarkMode);
toggleChartTypeSwitch.addEventListener('change', toggleChartType);


// Initial render
renderExpenses();
document.addEventListener('DOMContentLoaded', () => {
    renderCategoryOptions();
    renderExpenses();
});