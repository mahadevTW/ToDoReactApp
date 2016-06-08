var React = require('react');
var ToDoElement = React.createClass({
    render: function () {
        return (
            <div className='textElementStyle'>{this.props.text}</div>
        )
    }
});

module.exports = ToDoElement;