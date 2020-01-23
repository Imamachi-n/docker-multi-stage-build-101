# docker-multi-stage-build-101

## Qiitaの記事

Dockerのマルチステージビルドで、ビルド環境と実行環境のセットアップを１つのDockerfileで完結させよう  
<https://qiita.com/Imamachi-n/items/72c0148c454a810cb7fa>

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
