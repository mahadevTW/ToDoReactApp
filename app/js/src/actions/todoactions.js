var Reflux = require('reflux');
var ToDoActions = Reflux.createActions([
   "updateList",
   "fetchList",
   "deleteItem",
   "fetchCSRF",
]);

module.exports = ToDoActions;