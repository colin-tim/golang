import React from "react"
import styles from "./TodoItem.module.css"

class TodoItem extends React.Component {

  state = {
    editing: false
  }

  handleEditing = () => {
    this.setState({
      editing: true
    })
  } 

  handleUpdatedDone = (event, id, status) => {
    if (event.key === "Enter" ) {
      console.log("updating the task", event.target.value, id)
      this.props.editPropTask(id, event.target.value, status);
      this.setState({ editing: false });
    }
  }

  componentWillUnmount() {
    console.log("Cleaning up...")
  }
    
  render() {
    const completedStyle = {
      fontStyle: "italic",
      color: "#595959",
      opacity: 0.4,
      textDecoration: "line-through",
    }

    const {_id, task, status , createddatetime} = this.props.todo
    
    let viewMode = {};
    let editMode = {};

    if (this.state.editing) {
      viewMode.display = 'none';
    } else {
      editMode.display = 'none';
    }

    return (
      <li className={styles.item}>
        <div onDoubleClick={this.handleEditing} style={viewMode}>
          <input 
            type="checkbox" 
            className={styles.checkbox}
            checked={status}  
            onChange={() => this.props.editPropTask(_id, task, !status)}
          />
          <button onClick={() => this.props.deleteTodoProps(_id)}>Delete</button>
          <span style={status ? completedStyle : null} >{task}</span>
          <div className={styles.dateText}>Created on: {new Date(createddatetime).toLocaleString()}</div>
        </div>
        <input 
          type="text" 
          style={editMode} 
          className={styles.textInput} 
          value={task} 
          onChange={(e)=> {this.props.setUpdate(e.target.value, _id)}}
          onKeyDown={(e)=> {this.handleUpdatedDone(e, _id, task, status)}}
        />
      </li>
    )
  }
}

export default TodoItem