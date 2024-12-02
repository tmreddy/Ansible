### Building a Simple RESTful API in Node.js with Express

Node.js is a powerful, asynchronous, and event-driven runtime for building fast and scalable network applications. In this tutorial, we will create a basic RESTful API using **Node.js** and **Express** (a minimal web framework). Our API will allow basic CRUD (Create, Read, Update, Delete) operations for managing users.

### Prerequisites:
1. **Node.js** installed (You can download Node.js from [https://nodejs.org/](https://nodejs.org/)).
2. Basic knowledge of JavaScript and how to use the command line.

### Project Structure

```
/node-api
├── server.js
├── routes
│   └── userRoutes.js
├── models
│   └── userModel.js
├── package.json
└── package-lock.json
```

- **server.js**: The main server file where the Express app is initialized and the routes are used.
- **routes/userRoutes.js**: The file where we define the API routes.
- **models/userModel.js**: The file representing the user model and in-memory data store.
- **package.json**: Manages dependencies, scripts, and project metadata.

### Step 1: Initialize the Node.js Project

1. Create a new directory for your project and navigate into it:
   ```bash
   mkdir node-api
   cd node-api
   ```

2. Initialize a new Node.js project:
   ```bash
   npm init -y
   ```

3. Install the required dependencies:
   ```bash
   npm install express
   ```

   - **express**: A lightweight web application framework for building RESTful APIs.

### Step 2: Create the User Model

For simplicity, we will use an in-memory array to store users. In a real-world application, you would typically use a database like MongoDB, MySQL, or PostgreSQL.

**`models/userModel.js`**:

```javascript
// In-memory data store for users (mock database)
let users = [
    { id: 1, name: "John Doe", email: "john@example.com" },
    { id: 2, name: "Jane Smith", email: "jane@example.com" }
];

// Helper function to get all users
const getUsers = () => users;

// Helper function to get a user by ID
const getUserById = (id) => users.find(user => user.id === id);

// Function to add a new user
const addUser = (user) => {
    const newId = users.length + 1;
    const newUser = { id: newId, ...user };
    users.push(newUser);
    return newUser;
};

// Function to update an existing user
const updateUser = (id, userData) => {
    let user = getUserById(id);
    if (user) {
        user = { ...user, ...userData };
        users = users.map(u => (u.id === id ? user : u));
        return user;
    }
    return null;
};

// Function to delete a user
const deleteUser = (id) => {
    const userIndex = users.findIndex(u => u.id === id);
    if (userIndex !== -1) {
        const deletedUser = users.splice(userIndex, 1);
        return deletedUser[0];
    }
    return null;
};

module.exports = { getUsers, getUserById, addUser, updateUser, deleteUser };
```

- `getUsers()`: Returns all users.
- `getUserById(id)`: Finds and returns a user by their ID.
- `addUser(user)`: Adds a new user to the array.
- `updateUser(id, userData)`: Updates a user with new data.
- `deleteUser(id)`: Deletes a user by their ID.

### Step 3: Define the API Routes

Now, let's define the API routes for the CRUD operations.

**`routes/userRoutes.js`**:

```javascript
const express = require('express');
const router = express.Router();
const { getUsers, getUserById, addUser, updateUser, deleteUser } = require('../models/userModel');

// GET /users - Fetch all users
router.get('/users', (req, res) => {
    const users = getUsers();
    res.json(users);
});

// GET /users/:id - Fetch a specific user by ID
router.get('/users/:id', (req, res) => {
    const userId = parseInt(req.params.id);
    const user = getUserById(userId);
    if (user) {
        res.json(user);
    } else {
        res.status(404).json({ message: 'User not found' });
    }
});

// POST /users - Create a new user
router.post('/users', (req, res) => {
    const { name, email } = req.body;
    if (!name || !email) {
        return res.status(400).json({ message: 'Name and email are required' });
    }
    const newUser = addUser({ name, email });
    res.status(201).json(newUser);
});

// PUT /users/:id - Update a user by ID
router.put('/users/:id', (req, res) => {
    const userId = parseInt(req.params.id);
    const { name, email } = req.body;
    if (!name || !email) {
        return res.status(400).json({ message: 'Name and email are required' });
    }
    const updatedUser = updateUser(userId, { name, email });
    if (updatedUser) {
        res.json(updatedUser);
    } else {
        res.status(404).json({ message: 'User not found' });
    }
});

// DELETE /users/:id - Delete a user by ID
router.delete('/users/:id', (req, res) => {
    const userId = parseInt(req.params.id);
    const deletedUser = deleteUser(userId);
    if (deletedUser) {
        res.json({ message: `User with ID ${userId} deleted` });
    } else {
        res.status(404).json({ message: 'User not found' });
    }
});

module.exports = router;
```

- **GET `/users`**: Fetches all users.
- **GET `/users/:id`**: Fetches a specific user by ID.
- **POST `/users`**: Creates a new user with `name` and `email`.
- **PUT `/users/:id`**: Updates an existing user by ID.
- **DELETE `/users/:id`**: Deletes a user by ID.

### Step 4: Set Up the Express Server

Now, let’s set up the Express server to use the routes and start the application.

**`server.js`**:

```javascript
const express = require('express');
const bodyParser = require('body-parser');
const userRoutes = require('./routes/userRoutes');

const app = express();
const port = 3000;

// Middleware to parse JSON bodies
app.use(bodyParser.json());

// Use the user routes
app.use('/api', userRoutes);

// Start the server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
```

- **`app.use(bodyParser.json())`**: This middleware parses incoming JSON requests, making it easier to handle request bodies in routes.
- **`app.use('/api', userRoutes)`**: We prefix all user routes with `/api`.

### Step 5: Running the Application

To run the application, open a terminal in the project directory and execute:

```bash
node server.js
```

The server will start, and you should see the message `Server is running on http://localhost:3000`.

### Step 6: Testing the API

You can test the API using tools like **Postman**, **curl**, or any HTTP client. Below are examples of how to interact with the API.

#### 1. **Create a User (POST)**

```bash
curl -X POST http://localhost:3000/api/users -H "Content-Type: application/json" -d '{"name": "Alice", "email": "alice@example.com"}'
```

This will create a new user and return the user object with an `id` assigned.

#### 2. **Get All Users (GET)**

```bash
curl http://localhost:3000/api/users
```

This will return a list of all users in the system.

#### 3. **Get a User by ID (GET)**

```bash
curl http://localhost:3000/api/users/1
```

This will return the user with ID `1`.

#### 4. **Update a User (PUT)**

```bash
curl -X PUT http://localhost:3000/api/users/1 -H "Content-Type: application/json" -d '{"name": "Alice Updated", "email": "alice.updated@example.com"}'
```

This will update the user with ID `1` and return the updated user.

#### 5. **Delete a User (DELETE)**

```bash
curl -X DELETE http://localhost:3000/api/users/1
```

This will delete the user with ID `1` and return a confirmation message.

### Conclusion

In this tutorial, we've:
- Set up a simple Node.js application using **Express** to create a RESTful API.
- Implemented basic **CRUD** functionality to manage users.
- Tested the API using **curl**.

Node.js with Express is an excellent combination for building fast and scalable web APIs. You can easily extend this example by connecting to a real database (e.g., MongoDB or MySQL) and adding more complex features like validation,