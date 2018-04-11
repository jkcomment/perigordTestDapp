pragma solidity ^0.4.4;

contract Greeter {
    string public greeting;

    event Result(address from, string stored);

    function Greeter() {
       greeting = "Hello";
    }

    function setGreeting(string _greeting) public {
       greeting = _greeting;
       Result(msg.sender, greeting);
    }

    function greet() constant returns (string) {
       return greeting;
    }
}
