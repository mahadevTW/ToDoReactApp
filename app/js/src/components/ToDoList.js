var React = require('react');
var ToDoElement = require("./ToDoElement");

var ToDoList = React.createClass({
    render: function () {
        let toDoElements = this.props.todoelements;
        let elementList = toDoElements.map((element)=>{
            return <ToDoElement text={element}/>
            });
        return (
            <div>
            {elementList}</div>
        );
    }
});

module.exports = ToDoList;