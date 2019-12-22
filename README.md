# docker-multi-stage-build-101

## Docker マルチステージビルドのサンプルコード

### 1. ビルド環境と実行環境を同じ Docker ベースイメージ上に展開した場合

```bash
make dbuild1
make drun1
```

### 2. Docker のマルチステージビルドを使ってビルド環境と実行環境を分離した場合

```bash
make dbuild2
make drun2
```
