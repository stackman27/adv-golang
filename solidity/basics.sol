pragma solidity >=0.8.2 <0.9.0;

contract MyContract {
    string public myString = "My string"; // public global variables
    string name = "Example 1"; // local variable

    // GLOBAL VARIABLES
    address public origin;
    address public payer;
    uint public balance;

    // Constructor: Runs only once when SC is initialized
    constructor() {}

    // VISIBILITY
    string name1 = "name 1"; // no visibility outside of the contract
    string private name2 = "name 2"; // private: only accessible inside the smart contract
    string internal name3 = "name 3"; // internal: only accessible inside SC but can be inherited
    string public name4 = "name 4"; // public: can be called inside/outside + can be inherited

    // CUSTOM Modifier
    address private owner;
    modifier onlyOwner() {
        require(msg.sender == owner, "Error: caller must be owner");
        _;
    }

    // MAPPINGS: Key Value Pair
    mapping(uint => string) public names;

    // STRUCT
    struct Person {
        string name;
        string homeAddr;
        uint age;
    }

    // Array of Person
    Person[] public people;

    // EVENT
    event PersonCreated(address indexed _user, string _message);

    // Note: no "memory" key needed for uint because it has fixed size allocation whereas string has dynamic size allocation.
    function AddNewPerson(
        string memory _name,
        string memory _homeAddr,
        uint _age
    ) public {
        people.push(Person(_name, _homeAddr, _age));

        // "emit" will write logs to the blockchain
        emit PersonCreated(msg.sender, "new person has been added!");
    }

    function GetNewPerson(
        uint _index
    ) public view returns (string memory, string memory, uint) {
        Person storage person = people[_index];
        return (person.name, person.homeAddr, person.age);
    }

    // Set writes to blockchain so you have to pay gas
    function setName(string memory _name) public {
        name = _name;
    }

    // "onlyOwner" Custom Modifier User case
    function setNameOnlyowner(string memory _name) public onlyOwner {
        name = _name;
    }

    // Get reads blockchain so no need to pay gas.
    // "view" modifier: Cannot modify the state of the blockchain but can read the state.
    function getName() public view returns (string memory) {
        return name;
    }

    // "pure" modifier: doesnot read state and cannot modify state.
    function addLastName(
        string memory firstName,
        string memory lastName
    ) public pure returns (string memory) {
        return string.concat(firstName, lastName);
    }

    // "payable" modifier: allowed to recieve ether whenever a tx is submitted
    // initialize global variables here
    function pay() public payable {
        payer = msg.sender;
        origin = tx.origin;
        balance = msg.value;
    }

    // "public" accessible outside of smart contract
    function resetName() public {
        name = "Sishir Public";
    }

    // "Internal" only accessible in the smart contract
    function resetNameInternal() internal {
        name = "Sishir Internal";
    }

    // "External" only accessible outside of the smart contract
    function resetNameExternal() external {
        name = "Sishir External";
    }

    // GetBlockInfo returns some global variables, that is available throughout the SC.
    function getBlockInfo() public view returns (uint, uint, uint) {
        return (block.number, block.timestamp, block.chainid);
    }
}
