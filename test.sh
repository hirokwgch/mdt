
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