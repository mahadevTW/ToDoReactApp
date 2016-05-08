var expect = require('chai').expect;
var React = require('react');
var sinon = require('sinon');
var ReactTestUtils = require("react-addons-test-utils");
var ToDoInput = require("../../src/components/ToDoInput")
var JsDom = require("../utils/jsdom");
var ToDoActions = require("../../src/actions/todoactions")
var ToDoStore = require("../../src/stores/todostore")
describe('ToDoInput', function() {
        it('should trigger action on Form Submit', function () {
        // sinon.stub(ToDoActions, "updateList",function(){
        //         console.log("Stub Caled")
        // });
        
        // let enteredText = "ToDoItem1";
        var component = ReactTestUtils.renderIntoDocument(<ToDoInput/>);
        var inputBox = ReactTestUtils.findRenderedDOMComponentWithTag(
                component,
                'form'
                );
        
        expect(inputBox.props.onSubmit.name).to.equal("bound handleSubmit");      
        
        // let event = { preventDefault: () => {}};
        // inputBox.props.onSubmit(event);
        // component.state.text=enteredText;
        // component.handleSubmit(event);                        
                
        })
});