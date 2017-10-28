
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
setup() {
    go build
    chmod +x mdt
    PATH="${PATH}:`pwd`"
}

test_main() {
    local input=`cat <<EOF
menu,price
melon,500    
EOF
`
    local actual=`echo -e "${input}" | mdt`
    local expected=`cat <<EOF
| menu  | price |
| ----- | ----- |
| melon | 500   |
EOF
`
    assert_equal "${expected}" "${actual}"
}

clean() {
    rm mdt
}

# Main
echo "=== Setup ==="
setup
echo ""

echo "=== Execute test ==="
test_result=`test_main`
echo ""

echo "=== Show test result ==="
echo -e "${test_result}"
echo ""

echo "=== Clean ==="
clean
echo ""