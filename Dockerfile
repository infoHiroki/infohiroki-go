# Go 1.21のマルチステージビルド
FROM golang:1.21-alpine AS builder

# 作業ディレクトリ設定
WORKDIR /app

# Go modulesファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 本番用の軽量イメージ
FROM alpine:latest

# セキュリティ更新とCA証明書をインストール
RUN apk --no-cache add ca-certificates

# 作業ディレクトリ設定
WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# 静的ファイル・テンプレート・記事をコピー
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/articles ./articles

# ポート設定（環境変数で上書き可能）
EXPOSE 8080

# アプリケーション実行
CMD ["./main"]