var Reflux = require('reflux');
var ToDoActions = require('./../actions/Todoactions');

var ToDoStore = Reflux.createStore({
    listenables: [ToDoActions],

    onUpdateList: function(payload){
        this.trigger({
            action: "triggered",
            text: payload

        })
    }
});

module.exports = ToDoStore;