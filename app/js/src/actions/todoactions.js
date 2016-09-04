var Reflux = require('reflux');
var ToDoActions = Reflux.createActions([
   "updateList",
   "fetchList"
]);

module.exports = ToDoActions;