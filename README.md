# text-ch-parser

This repository provides a CLI tool which parses discord channel definitions written with Terraform and writes parsed data to a file.

このリポジトリは、Terraformを用いて書かれたDiscordのチャンネルの定義をパースし、結果を別ファイルに書き出すCLIツールを提供しています。（自分用といって差し支えないです。）

## Usage

```bash
$ text-ch-parser [flags]
```

### Flags

Both absolute path and relative one are OK.

| flag | description | default |
| --- | --- | --- |
| -f | the HCL file name you want to parse | main.tf |
| -w | the Markdown file name you want to write parsed HCL | main.md |
| -d | whether outputs logs for debugging | (false) |

### Examples

* With file name flags

```bash
$ text-ch-parser -f hcl.tf -w result.md
```

* With debug flag

```bash
$ text-ch-parser -d
```
