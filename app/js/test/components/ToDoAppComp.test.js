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
            let scope = nock('http://localhost/')
                .get('/todos')
                .reply(200,[
                    {Item:'item1', Id:1},
                    {Item:'item2', Id:2},
                ])
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);

            setTimeout(function(){
                var comps = ReactTestUtils.scryRenderedDOMComponentsWithClass(component,'textElementStyle');
                expect(comps.length).to.be.equal(3)
                expect(comps[1].props.children).to.equal("item1")
                expect(comps[2].props.children).to.equal("item2")
                expect(comps[1].props.id).to.equal(1)
                expect(comps[2].props.id).to.equal(2)

                done();
            }, 200);
        });

        xit('should update List when notified', function(){
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);
            let data={
                data:"Hello",
                action:"triggered"
            }
            ToDoStore.onUpdateList("Hello");
            expect(component.state.text).to.be.equal("Hello");

        })
});
