# 「Gemini CLI」がきた。Claude CodeをMaxプランで毎日使ってる人の感想

![Gemini CLI 記事のヘッダー画像](/images/note/n7d93c00246fd_d8835b6fc420448474d9eec26b0ab974.png)

## 🎯 中心的な主張
**GoogleのGemini CLIとClaude Codeを実際に比較したClaude Maxプランユーザーの視点から、Gemini CLIは速度・Windows対応・マルチモーダル機能で優位性を示すものの現状はデモレベルであり、一方Claude Codeは成熟したワークフロー統合で日常業務に不可欠であり、最大の差別化要因はWindows環境でのネイティブ動作にある。**

## 📖 詳細な説明

### 🔗 一次情報源
主要な公式リソース：
- **GitHub**: https://github.com/google-gemini/gemini-cli
- **Google公式ブログ**: AI Agent導入に関する包括的な説明

### 🤖 前提：Claude Maxプランでの運用実態

#### 使用状況
Claude Maxプランに加入し、**ほぼすべての作業にClaudeCodeを使用**：
- **メインタスク**: コーディングと文書作成
- **リサーチ作業**: 基本的にClaudeCodeで実行
- **補完ツール**: Gensparkブラウザ版（節約時・詳細リサーチ時のみ）

### 🆚 GeminiCLI vs Claude Code 比較分析

#### GeminiCLIの現状評価
**まだお試し版・デモレベル**の段階で、出たばかりのため当然の状況。Gemini 2.5 Proの使用感から評価を実施。

#### 主要な違い（実測値）

**🚀 GeminiCLIの優位性**
- **処理速度**: Claude Codeより高速
- **透明性**: 裏側の処理がリアルタイムで表示（Claude CodeはCtrl+Rが必要）
- **プラットフォーム**: **Windowsネイティブ対応**

**💻 Windowsネイティブ対応の重要性**
これが**最も大きな差別化要因**。Claude CodeはWindows環境ではWSL経由が必要で、ターミナルツールとしてのアクセシビリティが制限される。クロスプラットフォーム開発の技術的困難は理解できるが、ユーザー体験に直結する課題。

### 🔧 GeminiCLI 技術仕様（Claude Codeによるまとめ）

#### 概要
Googleが提供する**コマンドラインAIワークフローツール**で、コードの理解、編集、ワークフローの自動化を支援し、開発者の生産性向上を目的とする。

#### 主要特徴

**1. 大規模コードベースのサポート**
- **1Mトークンコンテキストウィンドウ**を活用
- 大規模プロジェクトでのクエリと編集が可能

**2. マルチモーダル機能**
- **PDFやスケッチから新しいアプリケーションを生成**
- 画像、ドキュメント、コードを同時に理解

**3. ツール統合**
- **MCPサーバー**を使用した新機能接続
- **Google検索ツール**組み込みによる最新情報アクセス
- **Imagen、Veo、Lyria**を使用したメディア生成

**4. 自動化機能**
- プルリクエストのクエリ
- 複雑なリベース処理
- ワークフローの自動化

#### 実用的な使用例

![Gemini CLI 使用例のスクリーンショット](/images/note/n7d93c00246fd_1750900373-vto41IGXPLc8kJ2BZY3gdpDl.png)

**新規プロジェクトの開始**
```bash
cd new-project/
gemini
> Write me a Gemini Discord bot that answers questions using a FAQ.md file I will provide
```

**既存プロジェクトでの作業**
```bash
git clone https://github.com/google-gemini/gemini-cli
cd gemini-cli
gemini
> Give me a summary of all of the changes that went in yesterday
```

#### 主要ユースケース

**1. コードベースの探索**
```
> Describe the main pieces of this system's architecture.
> What security mechanisms are in place?
```

**2. 既存コードの操作**
```
> Implement a first draft for GitHub issue #123.
> Help me migrate this codebase to the latest version of Java. Start with a plan.
```

**3. ワークフローの自動化**
```
> Make me a slide deck showing the git history from the last 7 days, grouped by feature and team member.
> Make a full-screen web app for a wall display to show our most interacted-with GitHub issues.
```

**4. システムとのインタラクション**
```
> Convert all the images in this directory to png, and rename them to use dates from the exif data.
> Organise my PDF invoices by month of expenditure.
```

#### スラッシュコマンド一覧

![Gemini CLI スラッシュコマンド一覧](/images/note/n7d93c00246fd_1750900973-d8arbSNGM4xLA7HQfDgsV9Ej.png)

GeminiCLIでは特定のタスクを実行するための専用コマンドが提供：

- `/glob` - ファイルパターンマッチング
- `/google_web_search` - Google検索の実行
- `/list_directory` - ディレクトリ内容の一覧表示
- `/read_file` - ファイルの読み取り
- `/read_many_files` - 複数ファイルの一括読み取り
- `/replace` - ファイル内容の置換
- `/run_shell_command` - シェルコマンドの実行
- `/save_memory` - メモリへの情報保存
- `/search_file_content` - ファイル内容の検索
- `/web_fetch` - Webページの取得
- `/write_file` - ファイルの書き込み

### 💡 追記情報
**Gemini Code Assist登録によりGemini 2.5 Proがより活用可能**になる機能が存在し、これは将来的な可能性を示唆している。

## 📊 実例・証拠

### 🔍 実際の使用体験
- **速度比較**: Gemini CLIがClaude Codeより明らかに高速
- **透明性**: Gemini CLIは処理過程がリアルタイム表示、Claude Codeは手動表示切替が必要
- **アクセシビリティ**: Windows環境でのネイティブ動作は開発者体験に直結する重要差分

### 💻 プラットフォーム対応の現実
- **Claude Code**: WSL経由のため追加セットアップが必要
- **Gemini CLI**: Windows環境で即座に利用可能
- **影響**: ターミナルツールとしてのアクセシビリティに大きな差

### 🚀 技術的優位性の検証
- **1Mトークンコンテキスト**: 大規模コードベース対応の実証
- **マルチモーダル処理**: PDF・画像からアプリ生成の実現
- **Google統合**: 検索・メディア生成ツールの標準装備

## ❓ 派生する問い
- Gemini CLIの本格リリース後のエンタープライズ環境導入戦略は？
- Claude CodeのWindows対応改善とクロスプラットフォーム最適化の展望は？
- マルチモーダルAI開発ツールの普及による開発ワークフロー変革の影響は？

## 🏷️ タグ

- note
- Gemini CLI
- Claude Code
- AI開発ツール
- 比較レビュー
- Google
- Anthropic
- 開発環境
- Windows
- マルチモーダル
- コマンドライン