#!/bin/bash

# 执行文件 (进程也会包含此字符串)
BIN="./checker"

# kill 进程
pgrep -f "${BIN}" | xargs -i kill -9 {}