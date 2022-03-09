import React from "react"
import axios from "axios";
import Header from "./Header"
import InputTodo from "./InputTodo"
import TodosList from "./TodosList";

class TodoContainer extends React.Component {
  endpoint = "http://localhost:8080";

  state = {
    todos: []
  };
  
  componentDidMount() {
    this.getTasks()
  };

  getTasks = () => {
    axios.get(this.endpoint + "/api/task").then((res) => {
      if (res.data) {
        this.setState({
          todos: res.data
          })
        }
       else {
        this.setState({
          todos: [],
        });
      }
    });
  };


  delTodo = id => {
    axios
    .delete(this.endpoint + "/api/deleteTask/" + id, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    })
    .then((res) => {
      console.log(res);
      this.getTasks();
    });
  };

  addTodoItem = (task) => {
    if (task) {
      axios
        .post(
          this.endpoint + "/api/task",
          {
            task,
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded",
            },
          }
        )
        .then((res) => {
          this.getTasks();
          console.log(res);
        });
    }
  };

  
  setUpdate = (updatedTitle, id) => {
    this.setState({
      todos: this.state.todos.map((todo) => {
        if(todo._id === id) {
          todo.task = updatedTitle
        }
        return todo
      })
    })
  }
 
  editTask = (id, task, status) => {
    console.log(id, task, status);
    axios
    .post(
        this.endpoint + "/api/task/" + id,
        {
            task,
            status,
        },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )    
      .then((res) => {
        console.log(res);
        this.getTasks();
      });
  };
  
  render() {
    return (
      <div className="container">
        <div className="inner">
          <Header />
          <InputTodo addTodoProps={this.addTodoItem} />
          <TodosList 
            todos={this.state.todos} 
            deleteTodoProps={this.delTodo}
            setUpdate ={this.setUpdate}
            editPropTask={this.editTask}
          />
        </div>
      </div>
    );
  }
}
export default TodoContainer