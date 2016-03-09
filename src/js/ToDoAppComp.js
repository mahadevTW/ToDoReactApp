var TextDisplay= require('./TextDisplay');
var ToDoInput= require('./ToDoInput');

var ToDoApp = React.createClass({
    render: function () {
        return (
            <div>
                <TextDisplay/>
                <ToDoInput/>
            </div>
        )
    }
});

module.exports = ToDoApp;