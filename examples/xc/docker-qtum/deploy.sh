#! /bin/bash

qcli generate 600

solar deploy contracts/INK.sol

solar deploy contracts/XCPlugin.sol '["7174756d00000000000000000000000000000000000000000000000000000000"]'

solar deploy contracts/XC.sol '["7174756d00000000000000000000000000000000000000000000000000000000"]'

