# Git & GitHub 完全チートシート

## 🎯 中心的な主張
**Git・GitHubの完全チートシートとして、基本概念から初期設定、ファイル操作、ブランチ管理、リモート操作、GitHub Flow実践、トラブルシューティング、高度テクニックまで7000文字超の包括的リファレンスを提供し、毎日使うTop 10コマンドと実践的ワークフローにより開発者の日常業務を完全サポートする決定版ガイド。**

## 📖 詳細な説明

### 🎯 基本概念の理解

#### Git とは
分散型バージョン管理システム。ファイルの変更履歴を記録し、複数人での開発を効率化します。

#### GitHub とは
Gitリポジトリをクラウドで管理するプラットフォーム。コラボレーション機能が充実。

#### 重要な用語集
- **リポジトリ**: プロジェクトのデータと履歴を保存する場所
- **コミット**: ファイルの変更を記録する操作
- **ブランチ**: 並行開発のための作業用ブランチ
- **マージ**: ブランチを統合する操作
- **プル**: リモートから最新版を取得
- **プッシュ**: ローカルの変更をリモートに送信
- **フォーク**: 他人のリポジトリを自分のアカウントにコピー
- **クローン**: リモートリポジトリをローカルにコピー

### ⚙️ 初期設定

#### ユーザー情報設定
```bash
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```
最初に必ず設定する基本情報

#### 設定確認
```bash
git config --list
git config user.name
```
現在の設定を確認

#### エディタ設定
```bash
git config --global core.editor "code --wait"
```
VSCodeをデフォルトエディタに設定

#### 改行コード設定
```bash
# Windows
git config --global core.autocrlf true
# Mac/Linux
git config --global core.autocrlf input
```
OS間の改行コード差異を自動処理

### 📝 Git 基本コマンド

#### リポジトリの作成・取得
- **新規リポジトリ作成**: `git init` - 現在のディレクトリをGitリポジトリに初期化
- **リポジトリをクローン**: `git clone <URL>` - リモートリポジトリをローカルにコピー

#### ファイルの操作
- **ステージング**: `git add <file>` / `git add .` / `git add -A` - ファイルを次回コミット対象に追加
- **ステータス確認**: `git status` / `git status -s` - ワーキングエリアの状態を確認
- **差分確認**: `git diff` / `git diff --staged` / `git diff <file>` - 変更内容の詳細を確認
- **コミット**: `git commit -m "message"` / `git commit -am "message"` - ステージされた変更をコミット

#### コミット履歴の確認
- **ログ表示**: `git log` / `git log --oneline` / `git log --graph --oneline` - コミット履歴を様々な形式で表示
- **特定のコミットを表示**: `git show <commit-hash>` / `git show HEAD` - コミットの詳細情報を表示

### 🌿 ブランチ操作

#### 基本的なブランチ操作
- **ブランチ一覧**: `git branch` / `git branch -a` / `git branch -r` - ローカル・リモートブランチを確認
- **ブランチ作成・切り替え**: `git branch <branch-name>` / `git checkout <branch-name>` / `git checkout -b <branch-name>` - 新しいブランチの作成と切り替え
- **ブランチ削除**: `git branch -d <branch-name>` / `git branch -D <branch-name>` - マージ済み(-d)・強制(-D)削除
- **マージ**: `git merge <branch-name>` / `git merge --no-ff <branch-name>` - 指定ブランチを現在のブランチにマージ

### 🌐 リモートリポジトリ操作

#### リモートとの連携
- **リモート確認・追加**: `git remote -v` / `git remote add origin <URL>` - リモートリポジトリの確認と追加
- **プッシュ**: `git push origin <branch>` / `git push -u origin <branch>` / `git push --all` - ローカルの変更をリモートに送信
- **プル・フェッチ**: `git pull origin <branch>` / `git fetch origin` / `git pull --rebase` - リモートの最新情報を取得

### 🔄 実践的なワークフロー

#### GitHub Flow（推奨）
1. **mainブランチから新しいブランチを作成**
2. **機能開発・修正を行う**
3. **定期的にコミットする**
4. **GitHubにプッシュ**
5. **Pull Requestを作成**
6. **レビュー後、mainにマージ**
7. **不要なブランチを削除**

### 🚨 トラブルシューティング

#### よくある問題と解決法

**コンフリクト解決**
```bash
# コンフリクト発生時
git status  # 競合ファイルを確認
# ファイルを手動編集後
git add <file>
git commit
```

**コミットメッセージの修正**
```bash
git commit --amend -m "新しいメッセージ"
```

**プッシュできない場合**
```bash
git pull --rebase origin main
git push origin main
```

### 📋 GitHub 特有の機能

#### Pull Request（PR）
コードレビューとマージのための仕組み。機能追加や修正を安全に本体に取り込めます。

#### Issue 管理
バグ報告・機能要求・タスク管理に使用。ラベルやマイルストーンで整理できます。

#### GitHub Actions（CI/CD）
```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Run tests
        run: npm test
```

### 🚀 高度なテクニック

#### 💡 便利な Git エイリアス
```bash
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.unstage 'reset HEAD --'
git config --global alias.last 'log -1 HEAD'
git config --global alias.visual '!gitk'
```

#### Git Bisect（バグ原因特定）
```bash
git bisect start
git bisect bad  # 現在のコミットが悪い
git bisect good <commit>  # 良いことが分かっているコミット
# Gitが自動で二分探索
git bisect reset  # 終了
```

#### クイックリファレンス：毎日使うコマンド Top 10
1. **`git status`** - 現在の状態確認 (★★★★★)
2. **`git add .`** - 全ての変更をステージング (★★★★★)
3. **`git commit -m "message"`** - コミット実行 (★★★★★)
4. **`git push`** - リモートにプッシュ (★★★★☆)
5. **`git pull`** - リモートから取得 (★★★★☆)
6. **`git checkout -b <branch>`** - 新ブランチ作成・切り替え (★★★☆☆)
7. **`git merge <branch>`** - ブランチマージ (★★★☆☆)
8. **`git log --oneline`** - コミット履歴確認 (★★☆☆☆)
9. **`git diff`** - 変更差分確認 (★★☆☆☆)
10. **`git reset HEAD <file>`** - ステージング取り消し (★★☆☆☆)

### 📚 学習リソース
- **公式ドキュメント**: https://git-scm.com/docs
- **インタラクティブ学習**: Learn Git Branching
- **GitHub Learning Lab**: https://lab.github.com/
- **Pro Git（無料書籍）**: https://git-scm.com/book

## 📊 実例・証拠

### 🔧 実践的な開発環境構築
- **ワンコマンドセットアップ**: 初期設定からリポジトリクローンまでの高速環境構築
- **エイリアス活用**: 日常コマンドの90%短縮による開発効率向上
- **GitHub Flow採用**: 7ステップによる安全で効率的な開発サイクル確立

### 💻 チーム開発の最適化
- **ブランチ戦略**: 並行開発とコンフリクト回避の両立実現
- **Pull Request**: コードレビュー文化による品質向上システム
- **Issue管理**: バグ追跡と機能要求の体系的管理

### 🚨 トラブル対応の体系化
- **コンフリクト解決**: 手動編集からマージまでの標準的手順
- **履歴修正**: commit --amendによる柔軟なコミット管理
- **Git Bisect**: 二分探索による効率的なバグ原因特定

## ❓ 派生する問い
- 大規模チーム開発でのGit Flow vs GitHub Flowの選択基準と実装戦略は？
- CI/CDパイプライン統合時のブランチ戦略とGitHub Actions最適化手法は？
- エンタープライズ環境でのGit権限管理とセキュリティポリシー設計は？

## 🏷️ タグ

- note
- Git
- GitHub
- チートシート
- バージョン管理
- コマンド
- リファレンス
- 開発ツール
- ブランチ戦略
- ワークフロー