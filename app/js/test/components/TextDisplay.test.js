var assert = require('assert');
var React = require('react');
var Reflux = require('reflux');
var ReactTestUtils = require("react-addons-test-utils");
var TextDisplay = require("./../../src/components/TextDisplay")
var JsDom = require("./../utils/jsdom");
describe('Text Display', function() {
        it('should display Text', function () {

                var component = ReactTestUtils.renderIntoDocument(<TextDisplay text="eher"/>);
                // var textDisplay = ReactTestUtils.findRenderedDOMComponentWithTag(
                //     component,
                //     'div'
                //         );

                // // var textDisplayComponent = ReactDOM

                // assert(textDisplay.textContent,"Helsdffasdflo");
        });
});