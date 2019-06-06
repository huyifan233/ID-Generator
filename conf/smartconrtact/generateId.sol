pragma solidity >=0.4.22;

contract GenerateId{


    bytes public id;

    uint256 public intId;

    constructor () public {
        id = "1";
        intId = 1;
    }


    function getId(uint step) public {
        uint flag = 0;
        // bytes32 bstep = new bytes(32);

        // for(uint i=id.length-1; i>=0 ;i--){

        // }
        intId += step;
        emit LogId(intId);
    }

    // function toBytes(uint256 x) private returns (bytes b) {
    //     b = new bytes(32);
    //     assembly { mstore(add(b, 32), x) }
    // }
    event LogId(uint currentId);



}