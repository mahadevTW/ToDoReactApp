var React = require('react');
var Reflux = require('reflux');
var ToDoList= require('./ToDoList');
var ToDoInput= require('./ToDoInput');
var ToDoStore= require('./../stores/todostore');

const listener = Reflux.ListenerMixin;
let list =[];
var ToDoApp = React.createClass({
    
    getInitialState: function () {
        return ({todoelements: []});
    },
    componentDidMount: function(){
        listener.listenTo(ToDoStore, this.onUpdateList);
    },
    onUpdateList: function (result) {
        list.unshift(result.text);
        this.setState({
            todoelements : list
        });
    },
    render: function () {
        return (
            <div>
                <ToDoInput/>
                <ToDoList todoelements={this.state.todoelements}/>
            </div>
        )
    }
});

module.exports = ToDoApp;