(function () {
  var ALLOWED_STATES = new Set(["open", "active", "completed"]);

  var STATE_LABELS = {
    open: "⚪ Open",
    active: "🟡 Active",
    completed: "✅ Completed"
  };

  var currentTodos = [];

  function generateTodoId() {
    return Date.now() + "-" + Math.random().toString(36).substr(2, 9);
  }

  function validateTodo(todo) {
    if (!todo || typeof todo !== "object") {
      throw new Error("Invalid todo item: expected object.");
    }

    if (!todo.id || typeof todo.id !== "string") {
      throw new Error("Invalid todo id.");
    }

    if (!todo.title || typeof todo.title !== "string" || todo.title.trim() === "") {
      throw new Error("Invalid todo title.");
    }

    if (todo.title.length > 256) {
      throw new Error("Invalid todo title: exceeds 256 characters.");
    }

    if (!ALLOWED_STATES.has(todo.state)) {
      throw new Error("Invalid state: " + todo.state + ". Allowed states: open, active, completed.");
    }
  }

  function validateTodos(todos) {
    if (!Array.isArray(todos)) {
      throw new Error("Invalid todos input: expected array.");
    }

    for (var i = 0; i < todos.length; i += 1) {
      validateTodo(todos[i]);
    }
  }

  function stateLabel(state) {
    return STATE_LABELS[state];
  }

  function createTodoRow(todo) {
    var li = document.createElement("li");
    li.className = "todo-row";
    li.setAttribute("data-todo-id", todo.id);

    var title = document.createElement("span");
    title.className = "todo-title";
    title.textContent = todo.title;

    var state = document.createElement("span");
    state.className = "todo-state";
    state.textContent = stateLabel(todo.state);

    li.appendChild(title);
    li.appendChild(state);

    return li;
  }

  function logActivity(type, payload) {
    if (!window.__todoLog) {
      window.__todoLog = [];
    }

    window.__todoLog.push({
      type: type,
      timestamp: new Date().toISOString(),
      payload: payload || {}
    });
  }

  function addTodo(title, config) {
    var todo = {
      id: generateTodoId(),
      title: title,
      state: "open"
    };

    validateTodo(todo);

    currentTodos.push(todo);
    window.TodoStorage.saveTodos(currentTodos);

    renderTodos({
      todos: currentTodos,
      listElement: config.listElement,
      emptyElement: config.emptyElement,
      errorElement: config.errorElement
    });

    logActivity("todo-add", { id: todo.id, title: todo.title });

    return todo;
  }

  function renderTodos(options) {
    var todos = options.todos || [];
    var listElement = options.listElement;
    var emptyElement = options.emptyElement;
    var errorElement = options.errorElement;

    validateTodos(todos);

    listElement.innerHTML = "";
    if (errorElement) {
      errorElement.hidden = true;
      errorElement.textContent = "";
    }

    if (todos.length === 0) {
      if (emptyElement) {
        emptyElement.hidden = false;
      }
      logActivity("list-load", { count: 0 });
      return;
    }

    if (emptyElement) {
      emptyElement.hidden = true;
    }

    for (var i = 0; i < todos.length; i += 1) {
      listElement.appendChild(createTodoRow(todos[i]));
    }

    logActivity("list-load", { count: todos.length });
  }

  function start(config) {
    var listElement = document.getElementById(config.listElementId);
    var emptyElement = document.getElementById(config.emptyElementId);
    var errorElement = document.getElementById(config.errorElementId);

    if (!listElement) {
      throw new Error("Missing list container.");
    }

    var todos;
    if (config.hasOwnProperty("todos")) {
      todos = config.todos;
    } else {
      todos = window.TodoStorage.loadTodos();
      if (todos.length === 0 && window.TODO_SAMPLE_DATA) {
        todos = window.TODO_SAMPLE_DATA;
      }
    }

    currentTodos = todos;

    try {
      renderTodos({
        todos: currentTodos,
        listElement: listElement,
        emptyElement: emptyElement,
        errorElement: errorElement
      });
    } catch (err) {
      if (errorElement) {
        errorElement.textContent = err.message;
        errorElement.hidden = false;
      }
      throw err;
    }
  }

  window.TodoApp = {
    start: start,
    addTodo: addTodo,
    renderTodos: renderTodos,
    validateTodos: validateTodos,
    stateLabel: stateLabel,
    currentTodos: currentTodos,
    ALLOWED_STATES: ALLOWED_STATES
  };
})();
