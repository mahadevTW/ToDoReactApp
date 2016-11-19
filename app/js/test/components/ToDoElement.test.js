var chai = require('chai');
var React = require('react');
var ReactTestUtils = require("react-addons-test-utils");
var ToDoElement = require("./../../src/components/ToDoElement");
var JsDom = require("./../utils/jsdom");
describe('Text Display', function() {
        it('should display Text', function () {
                let component = ReactTestUtils.renderIntoDocument(<ToDoElement text="Text Here" todo_id="item1"/>);

                chai.expect(component.refs.text.props.children).to.eq("Text Here");
                chai.expect(component.refs.closeBtn.props.className).to.eq("arrow");
                chai.expect(component.refs.element.props.id).to.eq("item1")
        });
});
