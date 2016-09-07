var JsDom = require("../utils/jsdom");
var React = require('react');
var expect = require('chai').expect;
var ReactTestUtils = require("react-addons-test-utils");
var ToDoComponent = require("../../src/components/ToDoAppComp");
var ToDoStore = require("../../src/stores/todostore");
var sinon = require("sinon");
var nock = require("nock");

describe('ToDoComp', function() {
        it('should render TextDisplay and ToDoInput', function () {
            let renderer = ReactTestUtils.createRenderer();
            renderer.render(<ToDoComponent />);
            let component = renderer.getRenderOutput();
            expect(component.props.children[0].type.displayName).to.equal('ToDoInput');
            expect(component.props.children[1].type.displayName).to.equal('ToDoList');
        });

        it('should fetch ToDos', function (done) {
            nock('http://localhost/')
                .get('/todos')
                .reply(200,[
                    {Item:'item1'},
                    {Item:'item2'},
                ])

            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);

            nock.restore();


            setTimeout(function(){
                var comps = ReactTestUtils.scryRenderedDOMComponentsWithClass(component,'textElementStyle');

                expect(comps.length).to.be.equal(3)


            }, 200);
                done();
        });

        xit('should update List when notified', function(){
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);
            let data={
                text:"Hello",
                action:"triggered"
            }
            ToDoStore.onUpdateList("Hello");
            expect(component.state.text).to.be.equal("Hello");
                                                             
        })
});