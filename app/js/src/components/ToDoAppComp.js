var React = require('react');
var Reflux = require('reflux');
var TextDisplay= require('./TextDisplay');
var ToDoInput= require('./ToDoInput');
var ToDoStore= require('./../stores/todostore');

const listener = Reflux.ListenerMixin;

var ToDoApp = React.createClass({
    getInitialState: function () {
        return ({text: "empty"});
    },
    componentDidMount: function(){
        listener.listenTo(ToDoStore, this.onUpdateList);
    },
    onUpdateList: function (result) {
        this.setState({
            text: result.text
        });
        console.log(result.action);
    },
    render: function () {
        return (
            <div>
                <TextDisplay text={this.state.text}/>
                <ToDoInput/>
            </div>
        )
    }
});

module.exports = ToDoApp;