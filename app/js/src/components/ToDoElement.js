var React = require('react');
var ToDoElement = React.createClass({
    render: function () {
        return (
            <div ref='element' id={this.props.todo_id} className='textElementStyle'>
                <text ref='text'>{this.props.text}</text>
                <a ref='closeBtn' href="#" className="arrow">
                </a>
            </div>
        )
    }
});

module.exports = ToDoElement;
