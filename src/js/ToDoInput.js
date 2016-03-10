var ToDoActions = require('./Todoactions');
var ToDoInput = React.createClass({
    getInitialState: function(){
        return {text: ''};

    },
    handleChange: function(event){
        this.setState({
            text: event.target.value,
        });
    },
    handleSubmit: function (e) {
        e.preventDefault();
        ToDoActions.updateList(this.state.text);
    },
    render: function () {
        return (
            <form className="commentForm" onSubmit={this.handleSubmit}>
                <input type="text" value={this.state.value} placeholder="ToDo" onChange={this.handleChange}/>
                <input type="submit" value="Post" />
            </form>
        )
    }
});

module.exports = ToDoInput;