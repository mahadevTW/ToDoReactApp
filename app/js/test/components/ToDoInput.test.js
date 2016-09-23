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
        var component,newText,event;

        before(function(){
                component = ReactTestUtils.renderIntoDocument(<ToDoInput/>);
                newText= "New ToDo Item";
                event = { preventDefault: () => {}};
        }),
        
        it('should trigger action on Form Submit', function (done) {

        sinon.stub(ToDoStore,"trigger",function(data){
                let expectedTriggerData = {
                        action:"triggered",
                        data:newText
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
        
        component.handleSubmit(event)
        expect(inputBox.props.onSubmit.name).to.equal("bound handleSubmit");      
        });
        
        it('should clear text on enter', function () {
                component.state.text=newText;

                component.handleSubmit(event);

                expect(component.state.text).to.equal('');
        })
});