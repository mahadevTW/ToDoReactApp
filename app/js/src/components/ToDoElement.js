var React = require('react');
var ToDoElement = React.createClass({
    render: function () {
        return (
            <div id={this.props.todo_id} className='textElementStyle'>{this.props.text}</div>
        )
    }
});

module.exports = ToDoElement;
