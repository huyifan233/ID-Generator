pragma solidity >=0.4.22;

contract GenerateId{


    bytes32 public id;

    constructor () public {
        id = "0";
    }


    function setId(bytes32 _id) payable public {
        id = _id;

        emit LogId(id);
    }

    event LogId(bytes32 _id);



}