var expect = require('chai').expect;
var Reflux = require('reflux');
var React = require('react');
var ReactTestUtils = require("react-addons-test-utils");
var JsDom = require("../utils/jsdom");
var ToDoList = require("./../../src/components/ToDoList");
var ToDoInput = require("./../../src/components/ToDoInput");

describe('ToDoList', function() {
        it('should render a list of todo elements', function () {
          let todo1 = {id:"1", item:"item1"};
          let todo2 = {id:"2", item:"item2"};
          let todoElements =[todo1, todo2]
          let component = ReactTestUtils.renderIntoDocument(<ToDoList todoelements={todoElements}/>);
          let comp1 = component.refs.list.props.children[0]
          let comp2 = component.refs.list.props.children[1]
          expect(comp1.props.todo_id).to.equal('1')
          expect(comp1.props.text).to.equal('item1')
          expect(comp2.props.todo_id).to.equal('2')
          expect(comp2.props.text).to.equal('item2')
        });
      });
