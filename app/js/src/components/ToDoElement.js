var React = require('react');
var ToDoActions = require("./../actions/todoactions");

var ToDoElement = React.createClass({
    render: function () {
        return (
            <div ref='element' id={this.props.todo_id} className='textElementStyle'>
                <text ref='text'>{this.props.text}</text>
                <a ref='closeBtn' href="#" className="arrow" onClick={this.deleteItem1}>
                </a>
            </div>
        );
        
    },
    deleteItem1: function(){
        ToDoActions.deleteItem(this.props.todo_id);
    }
});

module.exports = ToDoElement;
