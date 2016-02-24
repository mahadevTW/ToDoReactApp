var ToDoApp = React.createClass({
    render: function () {
        return (
            <ToDoInput/>
        )
    }
});

var ToDoInput = React.createClass({
    render: function () {
        return (
            <form className="commentForm">
                <input type="text" placeholder="ToDo"/>
                <input type="submit" value="Post"/>
            </form>
        )
    }
});


React.render(<ToDoApp />, document.getElementById('content'));
