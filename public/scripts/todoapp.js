var ToDoApp = React.createClass({
    render: function () {
        return (
            <div>
                <ToDoInput></ToDoInput>
            </div>
        )
    }
});

var ToDoInput = React.createClass({
    render: function () {
        <form onsubmit={this.handleSubmit}>
            <input type="text" placeholder="ToDo" onChange={this.handleChange}/>
            <input type="submit" value="Post"/>
        </form>
    }
});




ReactDOM.render(<ToDoApp />, document.getElementById('content'));
