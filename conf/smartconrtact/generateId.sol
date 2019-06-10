pragma solidity >=0.4.22;

contract GenerateId{


    bytes32 public id = "1";

    constructor () public {
        id = "1";
    }


    function setId(bytes32 _id) payable public {
        id = _id;
        emit LogId(id);
    }

    event LogId(bytes32 _id);



}