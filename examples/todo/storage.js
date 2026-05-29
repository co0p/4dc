(function () {
  var STORAGE_KEY = "todo-app-data";

  function saveTodos(todos) {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(todos));
  }

  function loadTodos() {
    var data = localStorage.getItem(STORAGE_KEY);
    if (!data) {
      return [];
    }
    try {
      return JSON.parse(data);
    } catch (err) {
      return [];
    }
  }

  window.TodoStorage = {
    saveTodos: saveTodos,
    loadTodos: loadTodos
  };
})();
