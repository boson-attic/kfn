// Add dependencies using comments, like:
// kfn:dependency fibonacci latest

const fibonacci = require('fibonacci');

module.exports = context => {
  return "Hello world! A cool fibonacci number: " + fibonacci.iterate(10).number ;
};