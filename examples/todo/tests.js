(function () {
  var resultsEl = document.getElementById("results");
  var summaryEl = document.getElementById("summary");

  function pushResult(name, passed, detail) {
    var li = document.createElement("li");
    li.className = passed ? "ok" : "fail";
    li.textContent = (passed ? "PASS" : "FAIL") + " - " + name + (detail ? " :: " + detail : "");
    resultsEl.appendChild(li);
    return passed;
  }

  function assert(condition, name, detail) {
    if (!condition) {
      throw new Error(name + (detail ? " :: " + detail : ""));
    }
  }

  function resetFixture() {
    document.getElementById("todo-list").innerHTML = "";
    var empty = document.getElementById("empty-state");
    var error = document.getElementById("error-state");
    empty.hidden = true;
    error.hidden = true;
    error.textContent = "";
    window.__todoLog = [];
    localStorage.clear();
  }

  function runAll() {
    var passCount = 0;
    var totalCount = 0;

    function run(name, fn) {
      totalCount += 1;
      try {
        resetFixture();
        fn();
        passCount += 1;
        pushResult(name, true);
      } catch (err) {
        pushResult(name, false, err.message);
      }
    }

    run("renders all sample todos", function () {
      window.TodoApp.start({
        todos: window.TODO_SAMPLE_DATA,
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var count = document.querySelectorAll("#todo-list .todo-row").length;
      assert(count === window.TODO_SAMPLE_DATA.length, "row count mismatch", "expected " + window.TODO_SAMPLE_DATA.length + ", got " + count);
    });

    run("renders correct emoji+text labels", function () {
      window.TodoApp.start({
        todos: window.TODO_SAMPLE_DATA,
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var text = document.getElementById("todo-list").textContent;
      assert(text.indexOf("⚪ Open") >= 0, "open label missing");
      assert(text.indexOf("🟡 Active") >= 0, "active label missing");
      assert(text.indexOf("✅ Completed") >= 0, "completed label missing");
    });

    run("shows empty state when dataset is empty", function () {
      window.TodoApp.start({
        todos: [],
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var empty = document.getElementById("empty-state");
      assert(empty.hidden === false, "empty state not visible");
    });

    run("fails fast on invalid state", function () {
      var threw = false;
      try {
        window.TodoApp.start({
          todos: [{ id: "bad", title: "Bad state", state: "invalid" }],
          listElementId: "todo-list",
          emptyElementId: "empty-state",
          errorElementId: "error-state"
        });
      } catch (err) {
        threw = true;
        assert(err.message.indexOf("Invalid state") >= 0, "wrong error", err.message);
      }
      assert(threw, "expected fail-fast throw");
    });

    run("writes list-load log entry", function () {
      window.TodoApp.start({
        todos: window.TODO_SAMPLE_DATA,
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      assert(Array.isArray(window.__todoLog), "missing log array");
      assert(window.__todoLog.length > 0, "no log entries");
      assert(window.__todoLog[0].type === "list-load", "unexpected first log type");
      assert(window.__todoLog[0].payload.count === window.TODO_SAMPLE_DATA.length, "log count mismatch");
    });

    run("storage saves and loads todos", function () {
      var todos = [{ id: "1", title: "Test", state: "open" }];
      window.TodoStorage.saveTodos(todos);
      var loaded = window.TodoStorage.loadTodos();
      assert(loaded.length === 1, "loaded count");
      assert(loaded[0].id === "1", "loaded id");
    });

    run("addTodo rejects empty title", function () {
      window.TodoApp.start({
        todos: [],
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var threw = false;
      try {
        window.TodoApp.addTodo("", {
          listElement: document.getElementById("todo-list"),
          emptyElement: document.getElementById("empty-state"),
          errorElement: document.getElementById("error-state")
        });
      } catch (err) {
        threw = true;
      }
      assert(threw, "expected throw on empty");
    });

    run("addTodo rejects title over 256 chars", function () {
      window.TodoApp.start({
        todos: [],
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var longTitle = new Array(258).join("x");
      var threw = false;
      try {
        window.TodoApp.addTodo(longTitle, {
          listElement: document.getElementById("todo-list"),
          emptyElement: document.getElementById("empty-state"),
          errorElement: document.getElementById("error-state")
        });
      } catch (err) {
        threw = true;
      }
      assert(threw, "expected throw on too long");
    });

    run("addTodo creates todo with open state", function () {
      window.TodoApp.start({
        todos: [],
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      var todo = window.TodoApp.addTodo("New task", {
        listElement: document.getElementById("todo-list"),
        emptyElement: document.getElementById("empty-state"),
        errorElement: document.getElementById("error-state")
      });

      assert(todo.state === "open", "state is open");
      assert(todo.title === "New task", "title matches");
      assert(todo.id, "id exists");
    });

    run("added todo persists across reload", function () {
      window.TodoApp.start({
        todos: [],
        listElementId: "todo-list",
        emptyElementId: "empty-state",
        errorElementId: "error-state"
      });

      window.TodoApp.addTodo("Persisted task", {
        listElement: document.getElementById("todo-list"),
        emptyElement: document.getElementById("empty-state"),
        errorElement: document.getElementById("error-state")
      });

      var loaded = window.TodoStorage.loadTodos();
      var found = loaded.some(function(t) { 
        return t.title === "Persisted task"; 
      });
      assert(found, "todo found in storage");
    });

    summaryEl.textContent = "Completed: " + passCount + "/" + totalCount + " passed";
    summaryEl.className = "summary " + (passCount === totalCount ? "ok" : "fail");
  }

  runAll();
})();
