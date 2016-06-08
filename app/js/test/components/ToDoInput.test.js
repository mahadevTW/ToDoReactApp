var expect = require('chai').expect;
var Reflux = require('reflux');
var React = require('react');
var sinon = require('sinon');
var ReactTestUtils = require("react-addons-test-utils");
var ToDoInput = require("../../src/components/ToDoInput");
var JsDom = require("../utils/jsdom");
var ToDoActions = require("./../../src/actions/todoactions");
var ToDoStore = require("./../../src/stores/todostore");
describe('ToDoInput', function() {
        var component,newText;

        before(function(){
                component = ReactTestUtils.renderIntoDocument(<ToDoInput/>);
                newText= "New ToDo Item";
        }),
        
        it('should trigger action on Form Submit', function (done) {

        sinon.stub(ToDoStore,"trigger",function(data){
                let expectedTriggerData = {
                        action:"triggered",
                        text:newText
                }
                expect(data).to.deep.equal(expectedTriggerData);
                ToDoStore.trigger.restore();
                done();
        });
        var inputBox = ReactTestUtils.findRenderedDOMComponentWithTag(
                component,
                'form'
                );
        component.state.text=newText;
        let event = { preventDefault: () => {}};
        
        component.handleSubmit(event)
        expect(inputBox.props.onSubmit.name).to.equal("bound handleSubmit");      
        });
        
});