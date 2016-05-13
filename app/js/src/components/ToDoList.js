var React = require('react');
var ToDoElement = require("./ToDoElement");

var ToDoList = React.createClass({
    render: function () {
        let toDoElements = this.props.todoelements;
        let list = toDoElements.map((element)=>{
            return <ToDoElement text={element}/>
            });
        return (
            <div>
            {list}</div>
        );
    }
});

module.exports = ToDoList;