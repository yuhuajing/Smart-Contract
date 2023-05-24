// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract store{

event versionchange(string oldVersion , string newVersion);
string public version;

constructor(string memory _version){
    version = _version;
}
function setVersion(string calldata _version)external {
    emit versionchange(version, _version);
    version = _version;
}

}

