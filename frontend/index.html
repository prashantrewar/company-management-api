document.getElementById('loginForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const response = await fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password })
    });

    const data = await response.json();
    if (response.ok) {
        localStorage.setItem('token', data.token);
        localStorage.setItem('role', data.role);
        alert('Login successful');
        showMenu(data.role);
    } else {
        alert('Login failed: ' + data.error);
    }
});

function showMenu(role) {
    document.getElementById('login').style.display = 'none';
    document.getElementById('adminMenu').style.display = role === 'Admin' ? 'block' : 'none';
    document.getElementById('hrMenu').style.display = role === 'HR' ? 'block' : 'none';
    document.getElementById('salesMenu').style.display = role === 'Sales' ? 'block' : 'none';
    document.getElementById('accountantMenu').style.display = role === 'Accountant' ? 'block' : 'none';
}




function showSection(section) {
    document.querySelectorAll('.section').forEach(el => el.style.display = 'none');
    document.getElementById(section).style.display = 'block';
}

async function fetchUsers() {
    const response = await fetch('/users', {
        headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    });
    const users = await response.json();
    const usersList = document.getElementById('usersList');
    usersList.innerHTML = users.map(user => `<p>${user.username} (${user.role})</p>`).join('');
}

async function fetchPayrolls() {
    const response = await fetch('/payroll', {
        headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    });

    if (response.ok) {
        const payrolls = await response.json();
        console.log('Fetched payrolls:', payrolls); // Log to verify the response
        const payrollsList = document.getElementById('payrollsList');
        payrollsList.innerHTML = payrolls.map(payroll => `<p>${payroll.employee_name} - ${payroll.amount} - ${payroll.status}</p>`).join('');
    } else {
        alert('Failed to fetch payrolls');
    }
}

async function fetchCustomers() {
    const response = await fetch('/customers', {
        headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    });
    const customers = await response.json();
    const customersList = document.getElementById('customersList');
    customersList.innerHTML = customers.map(customer => `<p>${customer.name} (${customer.email})</p>`).join('');
}

async function fetchBillings() {
    const response = await fetch('/billings', {
        headers: { 'Authorization': 'Bearer ' + localStorage.getItem('token') }
    });
    const billings = await response.json();
    const billingsList = document.getElementById('billingsList');
    billingsList.innerHTML = billings.map(billing => `<p>${billing.customer} - ${billing.amount}</p>`).join('');
}

document.getElementById('createUserForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const username = document.getElementById('createUsername').value;
    const password = document.getElementById('createPassword').value;
    const role = document.getElementById('createRole').value;

    const response = await fetch('/users', {
        method: 'POST',
        headers: { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({ username, password, role })
    });

    if (response.ok) {
        alert('User created successfully');
        showSection('login');
    } else {
        const data = await response.json();
        alert('User creation failed: ' + data.error);
    }
});

document.getElementById('createPayrollForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const employee_name = document.getElementById('payrollEmployee').value;
    const amount = parseFloat(document.getElementById('payrollAmount').value);
    const status = document.getElementById('payrollStatus').value;

    const response = await fetch('/payroll', {
        method: 'POST',
        headers: { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({ employee_name, amount, status })
    });

    if (response.ok) {
        alert('Payroll created successfully');
    } else {
        const data = await response.json();
        alert('Payroll creation failed: ' + data.error);
    }
});

document.getElementById('createCustomerForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const name = document.getElementById('customerName').value;
    const address = document.getElementById('customerAddress').value;
    const email = document.getElementById('customerEmail').value;

    const response = await fetch('/customers', {
        method: 'POST',
        headers: { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({ name, address, email })
    });

    if (response.ok) {
        alert('Customer created successfully');
    } else {
        const data = await response.json();
        alert('Customer creation failed: ' + data.error);
    }
});

document.getElementById('createBillingForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const customer_name = document.getElementById('billingCustomer').value;
    const amount = parseFloat(document.getElementById('billingAmount').value);
    const status = document.getElementById('billingStatus').value;
    

    const response = await fetch('/billings', {
        method: 'POST',
        headers: { 
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({ customer_name, amount, status })
    });

    if (response.ok) {
        alert('Billing created successfully');
    } else {
        const data = await response.json();
        alert('Billing creation failed: ' + data.error);
    }
});
