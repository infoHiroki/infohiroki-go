# 🚀 AIでNotionTaskerというChrome拡張機能を作ってみた話。

![NotionTasker Chrome Extension](/images/note/n4205dbdd13c4_cb3d85fdb97a82f7df57c5267bc06250.png)

## 💡 はじめに

**Alt+N**でNotionに簡単にページを追加できるChrome拡張機能「NotionTasker」を、AI開発ツールを使って作成しました。Cursor ProとClaude 3.5 Sonnetを活用した開発プロセスを詳しく解説します。

## 🎯 NotionTaskerとは？

NotionのデータベースにAlt+Nのキーボードショートカットで簡単にページを追加できるChrome拡張機能です。どのWebページからでも瞬時にNotionにタスクやメモを保存できます。

## 🛠️ 使用した技術スタック

### 🤖 Cursor Pro
AI統合開発環境として使用

### 🧠 Claude 3.5 Sonnet
AI coding assistantとして活用

### 🎨 Figma
アイコン作成ツール

### ⚡ Chrome Extension
Manifest V3対応

## 📋 開発プロセス

### 1️⃣ プロジェクトのセットアップ

Cursor Proを使用してChrome拡張機能の基本構造を作成。Manifest V3に対応した設定ファイルを生成。

```json
{
  "manifest_version": 3,
  "name": "NotionTasker",
  "version": "1.0",
  "description": "Alt+NでNotionにページを追加",
  "permissions": ["storage", "activeTab"]
}
```

### 2️⃣ キーボードショートカットの実装

Alt+Nのショートカットを設定し、どのWebページからでも拡張機能を呼び出せるようにしました。

![Extension Management](/images/note/n4205dbdd13c4_1728869853-MGD6qQRfijvuF4wLptyJHKTl.png)

### 3️⃣ デベロッパーモードでのテスト

Chrome拡張機能の管理画面でデベロッパーモードを有効にし、ローカルでの動作確認を実施。

![Developer Mode](/images/note/n4205dbdd13c4_1728869406-ETQDgkMaA6O3pCoPFBwYWKJS.png)

### 4️⃣ 拡張機能の読み込み

パッケージ化されていない拡張機能として読み込み、実際のブラウザ環境でテストを実行。

![Extension Loading](/images/note/n4205dbdd13c4_1728869946-ai9fG0CmwEUPVZcxusg4y2XO.png)

### 5️⃣ アイコン作成

Figmaを使用してChrome拡張機能用のアイコンを作成。16px、48px、128pxのサイズに対応。

![Figma Icon Creation](/images/note/n4205dbdd13c4_1728870700-squ7GvSBCjiQc4pMgDthEZwU.png)

### 6️⃣ ファイル構造の整理

最終的なファイル構造を整理し、Chrome Web Storeでの公開準備を完了。

![File Structure](/images/note/n4205dbdd13c4_1728870996-d7yj0SwWOuXLZpa5P1toCBbr.png)

## 💡 AI開発ツールの活用ポイント

Cursor ProとClaude 3.5 Sonnetの組み合わせにより、Chrome拡張機能の開発が大幅に効率化されました。特に、Manifest V3の設定やイベントリスナーの実装で、AI assistantが的確なコード提案を行ってくれました。

## 🎯 学んだこと

### ✨ 開発効率の向上
- **AI開発の効率性**: 従来の開発時間を大幅に短縮
- **Chrome拡張機能の基礎**: Manifest V3の理解が深まった
- **ユーザビリティ設計**: キーボードショートカットの重要性
- **デバッグプロセス**: 段階的なテストの重要性

## 🔧 今後の改善予定

現在はChrome Web Storeでの公開準備中です。今後はNotion APIとの連携強化、複数データベース対応、カスタマイズ機能の追加を予定しています。

### 📈 予定している機能

- Notion APIとの連携強化
- 複数データベース対応
- カスタマイズ機能の追加
- ユーザー設定の保存・管理

## 🎉 まとめ

AI開発ツールを活用することで、従来では時間のかかっていたChrome拡張機能開発を効率的に進めることができました。特にCursor ProとClaude 3.5 Sonnetの組み合わせは、コード生成からデバッグまで一貫してサポートしてくれる強力なツールでした。

NotionTaskerのようなシンプルな機能であっても、実際の開発プロセスを通じてAI開発ツールの可能性と実用性を実感できた貴重な経験となりました。

## 🏷️ タグ

- note
- Chrome拡張機能
- AI開発
- Cursor
- Claude
- Notion
- プログラミング
- 開発ツール