var chai = require('chai');
var React = require('react');
var ReactTestUtils = require("react-addons-test-utils");
var ToDoElement = require("./../../src/components/ToDoElement");
var ToDoStore = require("./../../src/stores/todostore");
var JsDom = require("./../utils/jsdom");
var sinon = require("sinon");
var ToDoActions = require("./../../src/actions/todoactions");

describe('Text Display', function() {
        it('should display Text', function () {
                let component = ReactTestUtils.renderIntoDocument(<ToDoElement text="Text Here" todo_id="item1"/>);

                chai.expect(component.refs.text.props.children).to.eq("Text Here");
                chai.expect(component.refs.closeBtn.props.className).to.eq("arrow");
                chai.expect(component.refs.element.props.id).to.eq("item1")
        });

        it('should call delete action on click of delete icon', function(done){
                sinon.stub(ToDoActions,"deleteItem", function(todoID){
                        chai.expect(todoID).to.equal("item1")
                        done();
                        }
                );
                let component = ReactTestUtils.renderIntoDocument(<ToDoElement text="Text Here" todo_id="item1"/>);
                
                component.deleteItem1();     
        });
});
