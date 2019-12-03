<script>
  import axios from "axios";

  axios.defaults.baseURL = "https://zod4qdelme.execute-api.ap-northeast-2.amazonaws.com/dev";
  axios.defaults.headers = {
    "Content-Type": "application/json"
  };
  let content = "";
  let base_url = "";
  let todos = [];

  function fetchTodos() {
    axios.get("/Test").then(res => {
      if (!res.data) {
        return;
      }

      todos = res.data;
      todos.sort((a, b) => {
        return new Date(b.created_at) - new Date(a.created_at);
      }).map(todo => {
        todo.active = false;
        return todo;
      });
    });
  }

  function handleSubmit(e) {
    if (e.keyCode !== 13) {
      return;
    }

    axios
      .post("/Test", {
        content
      })
      .then(res => {
        console.log(res.data);
        fetchTodos();
        content = "";
      })
      .catch(err => {
        console.log(err);
      });
  }

  function handleCheck(todo) {
    axios
      .put("/Test/" + todo.created_at, {
        checked: !todo.checked
      })
      .then(res => {
        console.log(res.data);
        fetchTodos();
      })
      .catch(err => {
        console.log(err);
      });
  }

  function handleDelete(todo) {
    axios
      .delete("/Test/" + todo.created_at)
      .then(res => {
        console.log(res.data);
        fetchTodos();
      })
      .catch(err => {
        console.log(err);
      });
  }

  function handleUpdate(e, todo) {
    if (e.keyCode !== 13) return;

    axios
      .put("/Test/" + todo.created_at, {
        ...todo,
        content: e.target.value,
      })
      .then(res => {
        console.log(res.data);
        fetchTodos();
      })
      .catch(err => {
        console.log(err);
      });
  }

  fetchTodos();
</script>
<style>
.complete-all {
  display: flex;
}
.complete-all> input {
  margin-right: 10px;
}
.content {
  margin-left: 10px;
}
.item {
  display: flex;
}
</style>

<section class="todoapp">
  <header class="header">
    <h1>todos</h1>
    <input
      bind:value={content}
      on:keyup={handleSubmit}
      class="new-todo"
      placeholder="What needs to be done?" />
  </header>
  <section>
    <span class="todo-count" />
    <ul class="filters">
      <div href="#/" class="selected">All</div>
      <div href="#/active">Active</div>
      <div href="#/completed">Completed</div>
      <button class="clear-completed">Clear completed</button>
    </ul>
  </section>
  <section>
    <ul class="todo-list">
      <div class='complete-all'>
        <input id="toggle-all" type="checkbox"/>
        <label for="toggle-all">Mark all as complete</label>
      </div>
      {#each todos as todo}
        <li>
          <div class="item">
            {#if todo.checked}
              <input
                type="checkbox"
                on:click={() => handleCheck(todo)}
                checked />
            {:else}
              <input
                type="checkbox"
                on:click={() => handleCheck(todo)} />
            {/if}
            <div class="content">
              {#if todo.active}
              <input type="text" value="{todo.content}" on:keyup={e => handleUpdate(e, todo)}/>
              {:else}
              <label on:click={() => todo.active=true}>{todo.content}</label>
              {/if}
            </div>
            <button
              type="button"
              on:click={() => handleDelete(todo)}>delete</button>
          </div>
        </li>
      {/each}
    </ul>
  </section>
</section>
<footer class="info">
  <p>Double-click to edit a todo</p>
  <p>
    Written by
    <a href="https://github.com/novemberde">Novemberde</a>
  </p>
</footer>
