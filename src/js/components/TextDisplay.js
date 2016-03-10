var Reflux = require('reflux');
var ToDoStore= require('./../todostore');

const listener = Reflux.ListenerMixin;

var TextDisplay = React.createClass({
    getInitialState: function () {
        return {text: 'empty'};
    },
    componentDidMount(){
        listener.listenTo(ToDoStore, this.onUpdateList);
    },
    onUpdateList: function (result) {
        this.setState({
            text: result.text
        });
        console.log(result.action);
    },
    render: function () {
        return (
            <div>{this.state.text}</div>
        )
    }
});

module.exports = TextDisplay;