var Reflux = require('reflux');
var ToDoActions = require('./Todoactions');

const ToDoStore = Reflux.createStore({
    listenables: [ToDoActions],

    onUpdateList: function(payload){
        this.trigger({
            action: "triggered",
            text: payload

        })
    }

});

module.exports = ToDoStore;