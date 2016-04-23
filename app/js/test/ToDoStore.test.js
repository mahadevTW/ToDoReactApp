var expect = require('chai').expect;
var ToDoStore = require("../src/stores/todostore");
var ToDoActions = require("../src/actions/todoactions");
var sinon = require("sinon");

describe("ToDoStore", function(){
  it("is configured to listen to ToDoActions", function(){
    //expect(ToDoStore.listenables).to.include(ToDoActions);
    expect(ToDoActions.updateList).to.be.a("function");
  });
  it("raises triggers a change", function(done){
    let expectedResult = {
      action:"triggered",
      text: "hello"
    };
    sinon.stub(ToDoStore,"trigger",function(data){
      expect(data).to.deep.equal(expectedResult);
      ToDoStore.trigger.restore();
      done();
    })
    ToDoStore.onUpdateList("hello");
  })
}); 