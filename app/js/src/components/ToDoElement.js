var React = require('react');
var ToDoElement = React.createClass({
    render: function () {
        return (
            <div>{this.props.text}</div>
        )
    }
});

module.exports = ToDoElement;