var JsDom = require("../utils/jsdom");
var React = require('react');
var expect = require('chai').expect;
var ReactTestUtils = require("react-addons-test-utils");
var ToDoComponent = require("../../src/components/ToDoAppComp");
var ToDoStore = require("../../src/stores/todostore");
var sinon = require("sinon");
var nock = require("nock");
var config = require("../../src/config");

describe('ToDoComp', function() {
    let csrfNock, fetchTodosNock;

    beforeEach(function(){

        csrfNock = nock('http://localhost/')
            .get('/csrfToken')
            .reply(200, {CSRFToken:'XXXXXYYYYY'});

        fetchTodosNock = nock('http://localhost/')
            .get('/todos')
            .reply(200,[
                {Item:'item1', Id:1},
                {Item:'item2', Id:2},
            ])


    });

    afterEach(function(){
        nock.cleanAll();
    })

    it('should render TextDisplay and ToDoInput', function () {
            let renderer = ReactTestUtils.createRenderer();
            renderer.render(<ToDoComponent />);
            let component = renderer.getRenderOutput();
            expect(component.props.children[0].type.displayName).to.equal('ToDoInput');
            expect(component.props.children[1].type.displayName).to.equal('ToDoList');
        });

    it('should fetch ToDos', function (done) {


            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);

            setTimeout(function(){
                var comps = ReactTestUtils.scryRenderedDOMComponentsWithClass(component,'textElementStyle');
                expect(comps.length).to.be.equal(3)
                expect(comps[1].props.children[0].props.children).to.equal("item1")
                expect(comps[2].props.children[0].props.children).to.equal("item2")
                expect(comps[1].props.id).to.equal(1)
                expect(comps[2].props.id).to.equal(2)

                done();
            }, 200);
        });

    it('should update List when item is added', function(done){

            let scope=nock('http://localhost/')
                .post('/todo',{
                    "Item": "Hello"
                })
                .reply(200,3);
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);
            let data={
                data:"Hello",
                action:"triggered"
            }
            ToDoStore.onUpdateList("Hello");
            setTimeout(function(){
                var comps = ReactTestUtils.scryRenderedDOMComponentsWithClass(component,'textElementStyle');
                expect(comps.length).to.be.equal(4)

                expect(comps[1].props.children[0].props.children).to.equal("item1")
                expect(comps[2].props.children[0].props.children).to.equal("item2")
                expect(comps[3].props.children[0].props.children).to.equal("Hello")
                expect(comps[1].props.id).to.equal(1)
                expect(comps[2].props.id).to.equal(2)
                expect(comps[3].props.id).to.equal(3)

                done();
            }, 200);

        });

    it('should update List when item is removed', function(done){

            let scope1=nock('http://localhost/')
                .delete('/todo',{
                    "Id": 1
                })
                .reply(200,"Success");
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);
            
            ToDoStore.onDeleteItem(1);
            
            setTimeout(function(){
                var comps = ReactTestUtils.scryRenderedDOMComponentsWithClass(component,'textElementStyle');
                expect(comps.length).to.be.equal(2);

                expect(comps[1].props.children[0].props.children).to.equal("item2");
                expect(comps[1].props.id).to.equal(2);

                done();
            }, 200);

        });

    it('should make csrf Fetch on load', function(done){
            var component = ReactTestUtils.renderIntoDocument(<ToDoComponent/>);

            setTimeout(function() {
                expect(config.csrfToken).to.be.equal("XXXXXYYYYY");
                done();
            }, 200);
        })

});
