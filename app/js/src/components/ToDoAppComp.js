var React = require('react');
var Reflux = require('reflux');
var ToDoList= require('./ToDoList');
var ToDoInput= require('./ToDoInput');
var ToDoStore= require('./../stores/todostore');

const listener = Reflux.ListenerMixin;
var ToDoApp = React.createClass({
    
    getInitialState: function () {
        return ({todoelements: []});
    },
    componentDidMount: function(){
        this.list=[];
        ToDoStore.onFetchList();
        listener.listenTo(ToDoStore, this.onUpdateList);
        listener.listenTo(ToDoStore, this.onFetchList);
    },
    onUpdateList: function (result) {
        if(result.action == "triggered") {
            this.list.unshift(result.data);
            this.setState({
                todoelements: this.list
            });
        }
    },
    onFetchList: function (result) {
        if(result.action == "fetch") {
            let todosFlattened = result.data.body.map(x=>x.Item)
            this.list = this.list.concat(todosFlattened)
            this.setState({
                todoelements: this.list
            });
        }
    },
    render: function () {
        return (
            <div className="main-container">
                    <ToDoInput/>
                    <ToDoList todoelements={this.state.todoelements}/>
            </div>
        )
    }
});

module.exports = ToDoApp;