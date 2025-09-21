# 📋 イシューリスト - infoHiroki Go移植プロジェクト

**生成日**: 2025-09-21
**対象**: Go言語ウェブサイト移植プロジェクト
**評価**: B+ (75/100) - 改善必要

---

## 🔴 高優先度イシュー（すぐ対応）

### Issue #1: XSS脆弱性対策
**優先度**: 🔴 **CRITICAL**
**ファイル**: `src/models/blog_post.go:80,84`

**問題**:
```go
// 危険：HTMLエスケープを回避
return template.HTML(html)      // Line 80
return template.HTML(b.Content) // Line 84
```

**影響**:
- 任意スクリプト実行可能
- Cookie盗取、セッションハイジャック
- フィッシング攻撃のリスク

**解決策**:
```go
import "github.com/microcosm-cc/bluemonday"

func (b *BlogPost) RenderContent() template.HTML {
    policy := bluemonday.UGCPolicy()

    if b.ContentType == "markdown" {
        html := blackfriday.Run([]byte(b.Content), ...)
        safeHTML := policy.Sanitize(string(html))
        return template.HTML(safeHTML)
    }

    safeHTML := policy.Sanitize(b.Content)
    return template.HTML(safeHTML)
}
```

---

### Issue #2: 設定のハードコーディング
**優先度**: 🔴 **HIGH**
**ファイル**: `main.go:42,85`

**問題**:
```go
db, err = gorm.Open(sqlite.Open("database/infohiroki.db"), &gorm.Config{})
r.Run(":8080")
```

**影響**:
- 環境に応じた設定変更不可
- デプロイ時の柔軟性欠如
- 開発・本番環境の切り替え困難

**解決策**:
```go
// 環境変数対応
dbPath := os.Getenv("DATABASE_PATH")
if dbPath == "" {
    dbPath = "database/infohiroki.db"
}

port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
r.Run(":" + port)
```

---

### Issue #3: 危険なエラーハンドリング
**優先度**: 🔴 **HIGH**
**ファイル**: `main.go:43-45`

**問題**:
```go
if err != nil {
    panic("データベース接続に失敗しました: " + err.Error())
}
```

**影響**:
- アプリケーション強制終了
- 本番環境でのサービス停止
- エラー情報の不適切な露出

**解決策**:
```go
if err != nil {
    log.Fatal("データベース接続に失敗しました: ", err)
    // または適切なエラーレスポンス
}
```

---

## 🟡 中優先度イシュー（次フェーズ）

### Issue #4: モノリシック設計
**優先度**: 🟡 **MEDIUM**
**ファイル**: `main.go` (449行)

**問題**:
- 全機能が1ファイルに集約
- ハンドラー、ビジネスロジック、データ処理が混在
- 保守性・テスト容易性の低下

**解決策**:
```
src/
├── handlers/     # HTTP ハンドラー分離
│   ├── home.go
│   ├── blog.go
│   └── pages.go
├── services/     # ビジネスロジック
│   ├── blog_service.go
│   └── content_service.go
├── repositories/ # データアクセス層
│   └── blog_repository.go
└── middleware/   # ミドルウェア
    └── auth.go
```

---

### Issue #5: データベース設計の問題
**優先度**: 🟡 **MEDIUM**
**ファイル**: `src/models/blog_post.go:18`

**問題**:
```go
Tags string `json:"tags"` // JSON array as string
```

**影響**:
- タグ検索の非効率性
- データ正規化不足
- 将来的な拡張困難

**解決策**:
```go
// 新しいテーブル設計
type Tag struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"uniqueIndex"`
}

type BlogPostTag struct {
    BlogPostID uint `gorm:"primaryKey"`
    TagID      uint `gorm:"primaryKey"`
}

type BlogPost struct {
    // ...
    Tags []Tag `gorm:"many2many:blog_post_tags;"`
}
```

---

### Issue #6: 検索パフォーマンス問題
**優先度**: 🟡 **MEDIUM**
**ファイル**: `main.go:263-268`

**問題**:
```go
dbQuery.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%")
```

**影響**:
- LIKE検索でフルテーブルスキャン
- 大量データ時のパフォーマンス劣化
- FTS機能未活用

**解決策**:
```sql
-- SQLiteでFTS5テーブル作成
CREATE VIRTUAL TABLE blog_posts_fts USING fts5(
    title, content, description, tags,
    content='blog_posts',
    content_rowid='id'
);

-- Go側でFTS検索実装
db.Raw("SELECT * FROM blog_posts WHERE id IN (SELECT rowid FROM blog_posts_fts WHERE blog_posts_fts MATCH ?)", query)
```

---

### Issue #7: ロギング不備
**優先度**: 🟡 **MEDIUM**
**ファイル**: 全体

**問題**:
- 構造化ログ未実装
- デバッグ情報不足
- 運用監視困難

**解決策**:
```go
import "github.com/sirupsen/logrus"

// 構造化ログ実装
log := logrus.WithFields(logrus.Fields{
    "action": "blog_search",
    "query":  query,
    "user_ip": c.ClientIP(),
})
log.Info("検索実行")
```

---

## 🟢 低優先度イシュー（将来対応）

### Issue #8: テストカバレッジ不足
**優先度**: 🟢 **LOW**
**影響**: 品質保証・リグレッション防止

**解決策**:
```bash
# テスト実装
go test ./...
go test -cover ./...

# 目標カバレッジ: 80%以上
```

---

### Issue #9: キャッシュ機能未実装
**優先度**: 🟢 **LOW**
**影響**: レスポンス速度最適化

**解決策**:
```go
// Redisまたはインメモリキャッシュ
import "github.com/go-redis/redis/v8"

// ブログ記事キャッシュ実装
```

---

### Issue #10: API仕様書不足
**優先度**: 🟢 **LOW**
**影響**: API利用時の混乱

**解決策**:
```yaml
# OpenAPI 3.0仕様書作成
swagger: "3.0"
info:
  title: "infoHiroki API"
  version: "1.0.0"
```

---

## 📊 対応優先度マトリックス

| イシュー | 優先度 | 影響度 | 工数 | 期限 |
|---------|--------|--------|------|------|
| XSS対策 | 🔴 CRITICAL | 高 | 2日 | 即座 |
| 設定外部化 | 🔴 HIGH | 中 | 1日 | 1週間 |
| エラーハンドリング | 🔴 HIGH | 中 | 1日 | 1週間 |
| アーキテクチャ分離 | 🟡 MEDIUM | 高 | 5日 | 1ヶ月 |
| DB設計改善 | 🟡 MEDIUM | 中 | 3日 | 1ヶ月 |
| 検索最適化 | 🟡 MEDIUM | 中 | 2日 | 1ヶ月 |
| ロギング実装 | 🟡 MEDIUM | 低 | 2日 | 2ヶ月 |
| テスト追加 | 🟢 LOW | 中 | 5日 | 3ヶ月 |
| キャッシュ実装 | 🟢 LOW | 低 | 3日 | 6ヶ月 |
| API仕様書 | 🟢 LOW | 低 | 2日 | 6ヶ月 |

---

## 🎯 推奨実装順序

### Phase 1: セキュリティ・安定性 (1-2週間)
1. ✅ XSS対策実装 (#1)
2. ✅ 設定外部化 (#2)
3. ✅ エラーハンドリング改善 (#3)

### Phase 2: アーキテクチャ改善 (1ヶ月)
4. ✅ コード分離・モジュール化 (#4)
5. ✅ データベース設計最適化 (#5)
6. ✅ 検索パフォーマンス向上 (#6)

### Phase 3: 運用品質向上 (2-3ヶ月)
7. ✅ ログ実装 (#7)
8. ✅ テスト実装 (#8)

### Phase 4: 拡張機能 (6ヶ月+)
9. ✅ キャッシュ実装 (#9)
10. ✅ API仕様書整備 (#10)

---

## 📝 注意事項

- **セキュリティイシュー (#1)** は最優先で対応必須
- **Phase 1** 完了までは本番デプロイ非推奨
- 各イシューは独立して対応可能
- 実装時は必ずテストケース追加を推奨

---

**次のアクション**: Issue #1 (XSS対策) から着手することを強く推奨します。