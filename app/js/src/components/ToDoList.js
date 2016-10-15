var React = require('react');
var ToDoElement = require("./ToDoElement");

var ToDoList = React.createClass({
    render: function () {
        let toDoElements = this.props.todoelements;
        let elementList = toDoElements.map((element)=>{
            return <ToDoElement todo_id={element.id} text={element.item}/>
            });
        return (
            <div ref="list">
            {elementList}</div>
        );
    }
});

module.exports = ToDoList;
