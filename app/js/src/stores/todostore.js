var Reflux = require('reflux');
var ToDoActions = require('./../actions/Todoactions');
var request = require('superagent');

var ToDoStore = Reflux.createStore({
    listenables: [ToDoActions],

    onUpdateList: function(payload){
        request
            .post('/todo')
            .send({"Item":payload})
            .end(function(err,res){
                ToDoStore.publish("triggered",payload);
            })
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