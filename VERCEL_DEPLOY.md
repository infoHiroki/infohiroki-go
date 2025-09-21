# 🚀 Vercel デプロイガイド

Go + Vercel対応が完了しました。以下の手順でVercel無料枠にデプロイできます。

## ✅ 完了済み項目

### Phase 1: データ層修正
- ✅ SQLite削除完了
- ✅ ファイルベース管理に移行
- ✅ 94件のブログ記事対応

### Phase 2: Vercel Functions化
- ✅ main.goを複数のapi/*.go関数に分割
- ✅ 全エンドポイントをVercel Functions対応
- ✅ 静的ファイルをpublic/に配置

### Phase 3: デプロイ設定
- ✅ vercel.json作成完了
- ✅ ルーティング設定完了
- ✅ 環境変数対応

## 📁 新しいプロジェクト構造

```
infohiroki-go/
├── api/                    # Vercel Functions
│   ├── common.go          # 共通処理（データ読み込み）
│   ├── index.go           # ホームページ
│   ├── blog.go            # ブログ一覧
│   ├── post.go            # ブログ記事詳細
│   ├── search.go          # 検索API
│   └── pages.go           # 固定ページ
├── public/                # 静的ファイル
│   ├── css/
│   ├── js/
│   └── images/
├── content/               # コンテンツ
│   ├── metadata.json     # ブログメタデータ
│   └── articles/         # Markdown/HTML記事
├── src/models/           # データモデル
├── vercel.json           # Vercel設定
└── go.mod               # Go依存関係
```

## 🚀 デプロイ手順

### 1. Vercel CLIインストール
```bash
npm i -g vercel
```

### 2. ログイン
```bash
vercel login
```

### 3. プロジェクトデプロイ
```bash
# プロジェクトルートで実行
vercel

# または直接デプロイ
vercel --prod
```

### 4. 環境変数設定（オプション）
```bash
# 必要に応じて環境変数を設定
vercel env add CONTENT_DIR
# production
# content
```

## 🔗 エンドポイント一覧

| エンドポイント | 機能 | Vercel Function |
|---------------|------|----------------|
| `/` | ホームページ | api/index.go |
| `/blog` | ブログ一覧 | api/blog.go |
| `/blog/:slug` | ブログ記事詳細 | api/post.go |
| `/services` | サービス | api/pages.go |
| `/products` | 開発製品 | api/pages.go |
| `/results` | 実績 | api/pages.go |
| `/about` | スキルスタック | api/pages.go |
| `/faq` | FAQ | api/pages.go |
| `/contact` | お問い合わせ | api/pages.go |
| `/api/search` | 検索API | api/search.go |

## 💰 コスト

- **Vercel無料枠**: $0/月
- **機能制限**:
  - Functions実行時間: 10秒/リクエスト
  - 月間実行回数: 100GB-hour
  - 帯域幅: 100GB/月

## 🔧 既存サイトとの比較

| 項目 | 既存サイト | Go版 |
|------|----------|------|
| **フレームワーク** | Vanilla HTML/CSS/JS | Go + Vercel Functions |
| **データ管理** | 静的ファイル | ファイルベース管理 |
| **記事数** | 94記事 | 94記事（同一） |
| **検索機能** | JavaScript | サーバーサイド |
| **SEO** | 静的HTML | 動的メタデータ |
| **パフォーマンス** | CDN配信 | Functions + CDN |

## ✨ 改善点

1. **サーバーサイド検索**: 高速で正確な検索
2. **動的メタデータ**: ブログ記事ごとの最適化されたSEO
3. **型安全性**: Goによる堅牢なコード
4. **拡張性**: 新機能追加が容易

## 🎯 次のステップ（オプション）

1. **カスタムドメイン**: 既存ドメインの移行
2. **Analytics**: Vercel Analyticsの有効化
3. **CDN最適化**: 画像最適化の追加
4. **キャッシュ戦略**: Edge Cacheの活用

---

**これで完全にVercel無料枠での運用が可能になりました！** 🎉