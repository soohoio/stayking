#!/bin/bash
BASE_SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# run test files
BATS=${BASE_SCRIPT_DIR}/bats/bats-core/bin/bats
INTEGRATION_TEST_FILE=${BASE_SCRIPT_DIR}/integration_tests.bats 

CHAIN_NAME=GAIA TRANSFER_CHANNEL_NUMBER=0 $BATS $INTEGRATION_TEST_FILE
CHAIN_NAME=JUNO TRANSFER_CHANNEL_NUMBER=1 $BATS $INTEGRATION_TEST_FILE
CHAIN_NAME=OSMO TRANSFER_CHANNEL_NUMBER=2 $BATS $INTEGRATION_TEST_FILE