var chai = require('chai');
var React = require('react');
var Reflux = require('reflux');
var ReactTestUtils = require("react-addons-test-utils");
var TextDisplay = require("./../../src/components/TextDisplay")
var JsDom = require("./../utils/jsdom");
describe('Text Display', function() {
        it('should display Text', function () {

                var component = ReactTestUtils.renderIntoDocument(<TextDisplay text="Text Here"/>);
                var textDisplay = ReactTestUtils.findRenderedDOMComponentWithTag(
                    component,
                    'div'
                        );

                chai.expect(textDisplay.innerHTML).to.eq("Text Here");

        });
});