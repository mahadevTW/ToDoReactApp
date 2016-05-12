var React = require('react');
var JsDom = require("../utils/jsdom");
var expect = require('chai').expect;
var ReactTestUtils = require("react-addons-test-utils");
var ToDoComponent = require("../../src/components/ToDoAppComp");
var ToDoStore = require("../../src/stores/todostore");
var sinon = require("sinon");

describe('ToDoComp', function() {
        it('should render TextDisplay and ToDoInput', function () {
            let renderer = ReactTestUtils.createRenderer();
            renderer.render(<ToDoComponent />);
            let component = renderer.getRenderOutput();
            expect(component.props.children[0].type.displayName).to.equal('TextDisplay');
            expect(component.props.children[1].type.displayName).to.equal('ToDoInput');
        });   
        it.only('should update List when notified', function(){
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);
            let data={
                text:"Hello"
            }
            ToDoStore.onUpdateList("Hello");
            expect(component.state.text).to.be.equal("Hello");
                                                             
        })   
});