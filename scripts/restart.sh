#!/bin/bash

CURRENT_DIR=$(cd ../; pwd)
cd "${CURRENT_DIR}" || return

sh killAll.sh && sleep 1 && sh checkAlive.sh