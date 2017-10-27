
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

    if [ ${expected} == ${actual} ]; then
        ok "equal. expected=\"${expected}\", actual=\"${actual}\""
    else
        ng "NOT equal. expected=\"${expected}\", actual=\"${actual}\""
    fi 
}

# ---
# テストヘルパーのテスト

assert_equal "orange" "orange"
# 目視確認 -> OK: equal. expected="orange", actual="orange"

assert_equal "orange" "mikan"
# 目視確認 -> NG: NOT equal. expected="apple", actual="apple"
