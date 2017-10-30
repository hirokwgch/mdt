mdt: mark down table

#概要
csv形式のデータを、Markdownの表形式に変換して出力します。

#例
```bash
$ cat <<EOF | mdt
> menu,price
> melon,1000
> apple,100
> mix juice,10000
> EOF
| menu      | price |
| :--       | :--   |
| melon     | 1000  |
| apple     | 100   |
| mix juice | 10000 |
```