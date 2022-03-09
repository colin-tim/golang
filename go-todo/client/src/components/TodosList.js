import React from "react"
import TodoItem from "./TodoItem";

class TodosList extends React.Component {
  render() {
    return (
      <ul>
        {this.props.todos.map(todo => (
          <TodoItem
            key={todo._id}
            todo={todo}
            deleteTodoProps={this.props.deleteTodoProps}
            setUpdate={this.props.setUpdate}
            editPropTask={this.props.editPropTask}
          />
        ))}
      </ul>
    )
  }
}

export default TodosList