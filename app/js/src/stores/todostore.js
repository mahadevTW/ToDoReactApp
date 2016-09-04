var Reflux = require('reflux');
var ToDoActions = require('./../actions/Todoactions');
var request = require('superagent');

var ToDoStore = Reflux.createStore({
    listenables: [ToDoActions],

    onUpdateList: function(payload){
        this.trigger({
            action: "triggered",
            text: payload

        })
    },
    onFetchList: function(){
        request
            .get('/todos')
            .end(function(err,res){
                ToDoStore.publish(res)
            });
    },
    publish: function(data){
        this.trigger({
            action:"fetch",
            data:data,
        })
    },

});

module.exports = ToDoStore;