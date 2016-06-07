var React = require('react');
var ToDoActions = require('./../actions/Todoactions');
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
        this.setState({
            text:'',
        })
    },
    render: function () {
        return (
            <form className="commentForm" onSubmit={this.handleSubmit}>
                <input type="text" value={this.state.text} placeholder="ToDo" onChange={this.handleChange}/>
                <input ref="buttonClick" type="submit" value="Post" hidden/>
            </form>
        )
    }
});

module.exports = ToDoInput;