var Reflux = require('reflux');
var ToDoActions = require('./../actions/Todoactions');
var request = require('superagent');

var ToDoStore = Reflux.createStore({
    listenables: [ToDoActions],

    onUpdateList: function(payload){
        ToDoStore.publish("triggered",payload);
    },
    onFetchList: function(){
        request
            .get('/todos')
            .end(function(err,res){
                ToDoStore.publish("fetch",res)
            });
    },
    publish: function(action, data){
        this.trigger({
            action:action,
            data:data,
        })
    },

});

module.exports = ToDoStore;