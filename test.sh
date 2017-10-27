
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

# テストケース
test_main() {
    local actual=`echo -e "menu,price\nmelon,500" | mdt`
    local expected=`cat <<EOF
| menu  | price |
| ----- | ----- |
| melon | 500   |
EOF
`
    assert_equal "${expected}" "${actual}"
}

echo "=== Start test ==="
test_result=`test_main`

echo "=== Test result ==="
echo -e "${test_result}"
