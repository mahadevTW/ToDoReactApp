var React = require('react');
var ToDoAppComp = require('./components/ToDoAppComp');


var Main = React.createClass({
    render: function () {
        return (
        	<ToDoAppComp></ToDoAppComp>
            
        )
    }
});

React.render(<Main />, document.getElementById('content'));
