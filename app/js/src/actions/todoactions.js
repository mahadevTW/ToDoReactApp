var Reflux = require('reflux');
var ToDoActions = Reflux.createActions([
   "updateList",
   "fetchList",
   "deleteItem",
]);

module.exports = ToDoActions;