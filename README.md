# ConvexHull Drawer

点をクリックして凸包を描画するアプリケーションです。

## 使い方

github pages で公開しています。

[ConvexHull Drawer](https://hamao0820.github.io/convexHull-drawer/)

ローカルで実行するには Go が必要です。

```sh
go run .
```

## 遊び方

1. 画面上でクリックして点を追加します。
2. 点の上で右クリックすると点を削除できます。
3. escape キーを押すとリセットできます。

## 参考

[グラハムスキャン | アルゴリズムビジュアル大事典](https://yutaka-watanobe.github.io/star-aida/1.0/algorithms/graham_scan/print.html)

[https://x.com/AtCoder8/status/1762758925932453901?s=20](https://x.com/AtCoder8/status/1762758925932453901?s=20)

を参考にさせていただきました

## wasm

```
env GOOS=js GOARCH=wasm go build -o main.wasm github.com/hamao0820/convexHull-drawer
```
