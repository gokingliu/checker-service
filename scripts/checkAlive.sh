#!/bin/bash

# 当前目录
CURRENT_DIR=$(cd ../; pwd)
cd "${CURRENT_DIR}" || return

# 资源控制
ulimit -n 100000
ulimit -c unlimited
ulimit -s 65536

# 执行文件 (进程也会包含此字符串)
BIN="./checker"
# 配置文件
CONFIG="conf/trpc_go.yaml"
# 修改执行文件权限
chmod a+x ${BIN}

# 查找进程
TOTAL_PID=$(pgrep -f "${BIN}" | wc -l)
# 判断进程数量
if [[ $TOTAL_PID -gt 0 ]]; then
    echo "${BIN} -conf ${CONFIG} service is alive"
else
    echo "restart service "
    nohup ${BIN} -conf ${CONFIG} > /dev/null 2>&1 &
fi

# 等待 1 秒
sleep 1

# 再次查找进程
TOTAL_PID=$(pgrep -f "${BIN}" | wc -l)
# 判断进程数量
if [[ TOTAL_PID -gt 0 ]]; then
    echo "${BIN} -conf ${CONFIG} service is alive"
else
    echo "${BIN} -conf ${CONFIG} service is bad"
fi