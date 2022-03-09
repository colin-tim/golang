import React, { Component } from "react"

class InputTodo extends Component {
  state = {
    task: "",
  }
  onChange = e => {
    this.setState({
      [e.target.name]: e.target.value,    
    })
  }

  handleSubmit = e => {
    e.preventDefault();
    if(this.state.task.trim()) {
      this.props.addTodoProps(this.state.task);
      this.setState({    
        task: ""  
      });
    } else {
      alert("Please write item")
    } 
  };

  render() {
    return (
      <form onSubmit={this.handleSubmit} className="form-container">
        <input
          type="text"
          className="input-text"
          placeholder="Add todo..."
          value={this.state.task}
          name="task"          
          onChange={this.onChange}
        />
        <button className="input-submit">Submit</button>
      </form>
    )
  }
}
export default InputTodo


