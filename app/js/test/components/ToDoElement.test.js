var chai = require('chai');
var React = require('react');
var ReactTestUtils = require("react-addons-test-utils");
var ToDoElement = require("./../../src/components/ToDoElement");
var JsDom = require("./../utils/jsdom");
describe('Text Display', function() {
        it('should display Text', function () {

                var component = ReactTestUtils.renderIntoDocument(<ToDoElement text="Text Here"/>);
                var toDoElement = ReactTestUtils.findRenderedDOMComponentWithTag(
                    component,
                    'div'
                        );
                chai.expect(toDoElement.innerHTML).to.eq("Text Here");

        });
});