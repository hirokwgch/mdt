#!/bin/bash

source test.sh

# ---
# テストヘルパーのテスト

assert_equal "orange" "orange"
# 目視確認 -> OK: equal. expected="orange", actual="orange"

assert_equal "orange" "mikan"
# 目視確認 -> NG: NOT equal. expected="orange", actual="mikan"
# b
# c", actual="fizz
# buzz

expected=`cat <<EOF
a
b
c
EOF
` 
actual=`cat <<EOF
fizz
buzz
EOF
`
assert_equal "${expected}" "${actual}"
# 目視確認
# NG: NOT equal. expected="a
# b
# c", actual="fizz
# buzz"
#
# diff:
# 1,3c1,2
# < a
# < b
# < c
# ---
# > fizz
# > buzz
