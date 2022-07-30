# notify

デスクトップの通知センタにメッセージを表示するコマンド

## Usage

ポモドーロ:

```sh
$ notify -duration 2700 45分作業しました 休憩してください &
```

コマンドの終了通知:

```sh
$ long-running-command && notify コマンド実行が終了しました
```

## Requirements

go +1.13

## Installation

```sh
$ go install github.com/micheam/notify
```

## TODO

- [ ] 通知メッセージのアイコンを指定する
- [ ] 通知メッセージのアイコンを変更可能にする
- [ ] duration の単位を指定可能にする (eg: `2s`, `3hours`)
- [ ] 標準入力からメッセージ本文を指定できるようにする
- [ ] 通知する'時刻'を指定できるようにする

## License
MIT

## Author
Michito Maeda <https://github.com/micheam>
