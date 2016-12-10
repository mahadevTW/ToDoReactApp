var React = require('react');
var Reflux = require('reflux');
var ToDoList= require('./ToDoList');
var ToDoInput= require('./ToDoInput');
var ToDoActions= require('./../actions/todoactions');
var ToDoStore = require('./../stores/todostore');
var config = require('./../config');

const listener = Reflux.ListenerMixin;
var ToDoApp = React.createClass({
    
    getInitialState: function () {
        return ({todoelements: []});
    },
    componentDidMount: function(){
        this.list=[];
        ToDoStore.onFetchList();
        ToDoActions.fetchCSRF();
        listener.listenTo(ToDoStore, this.onUpdateList);
        listener.listenTo(ToDoStore, this.onFetchList);
        listener.listenTo(ToDoStore, this.onDeleteItem);
        listener.listenTo(ToDoStore, this.onFetchCSRF);
    },
    onUpdateList: function (result) {
        if(result.action == "triggered") {
            let todosFlattened = [{item:result.data.Item, id: result.data.Id}]
            this.list = this.list.concat(todosFlattened);
            this.setState({
                todoelements: this.list
            });
        }
    },
    onFetchList: function (result) {
        if(result.action == "fetch") {
            let todosFlattened = result.data.body.map(x=>({item:x.Item, id:x.Id}))
            this.list = this.list.concat(todosFlattened)
            this.setState({
                todoelements: this.list
            });
        }
    },
    onDeleteItem:function(result){
        if(result.action == "deleteItem"){

            this.list = this.list.filter(x=>x.id != result.data)
            this.setState({
                todoelements: this.list
            })

        }
    },
    onFetchCSRF: function(result){
        if(result.action == "csrfToken"){
            config.csrfToken = result.data
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