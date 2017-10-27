
# test helper
ok() {
    local example=${1}
    echo "OK: ${example}"
}

ng() {
    local example=${1}
    echo "NG: ${example}"
}

assert_equal() {
    local expected=${1}
    local actual=${2}

    diff <(echo -e "${expected}") <(echo -e "${actual}") > /dev/null
    local exit_code=$?
    
    if [ "${exit_code}" == "0" ]; then
        ok "equal. expected=\"${expected}\", actual=\"${actual}\""
    else
        ng "NOT equal. expected=\"${expected}\", actual=\"${actual}\""
        echo -e "\ndiff: "
        diff <(echo -e "${expected}") <(echo -e "${actual}")
        echo ""
    fi 
}

# ---
# テストヘルパーのテスト

assert_equal "orange" "orange"
# 目視確認 -> OK: equal. expected="orange", actual="orange"

assert_equal "orange" "mikan"
# 目視確認 -> NG: NOT equal. expected="orange", actual="mikan"

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


# テストケース
test_main() {
    local actual=`echo -e "menu,price\npie,100" | mdt`

}

