var assert = require('assert');
var React = require('react');
var Reflux = require('reflux');
var ToDoStore= require('../../../src/js/stores/todostore');
var ReactTestUtils = require("react-addons-test-utils");
var TextDisplay = require("../../../src/js/components/TextDisplay");

describe('TextDisplay', function() {
        it('should display Text', function () {
                var component = ReactTestUtils.renderIntoDocument(<TextDisplay></TextDisplay>);
                var textDisplay = ReactTestUtils.findRenderedDOMComponentWithTag(
                    component,
                    'div'
                        );
                assert(textDisplay.getDOMNode().getAttribute('innerHTML'),"Hello");
        });
});