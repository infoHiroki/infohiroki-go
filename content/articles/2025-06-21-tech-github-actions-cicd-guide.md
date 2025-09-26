# GitHub Actions 完全解説

## 🎯 中心的な主張
**GitHub ActionsはGitHubが提供する強力なCI/CDプラットフォームで、YAMLファイルによる簡潔な設定でプルリクエスト作成時の自動テスト実行からGoogle Cloud Runへの自動デプロイまで開発ワークフロー全体を自動化でき、特にPythonプロジェクトにおいて環境変数・シークレット管理・Workload Identity活用によりモダンな開発手法を容易に実装できる優れたツール。**

## 📖 詳細な説明

### 🔧 GitHub Actions とは

#### 概要
GitHubが提供するワークフロー自動化の仕組み
- ビルド、テスト、デプロイの自動化が可能
- イベントトリガーによる自動実行
- GitHubリポジトリとの統合

#### 料金体系
- **パブリックリポジトリ**: 基本無料
- **プライベートリポジトリ**: 一定の無料枠あり、超過分は有料
- 実行時間とOS種類により変動

### 🏗️ ワークフローの基本構造

**フローの構成**: ワークフロー（全体の処理） → トリガー（実行のきっかけ） → ジョブ（処理の単位） → ステップ（具体的なコマンド）

#### ワークフロー定義
- `.github/workflows`ディレクトリにYAMLファイルを配置
- 一つのYAMLファイルが一つのワークフロー
- 複数のジョブを並列・直列で実行可能

#### 主要なトリガー
- **pull_request**: プルリクエスト作成時
- **push**: 特定ブランチへのプッシュ時
- **schedule**: 定期実行（cron形式）

### 🧪 CI/CD の実装

#### CI（継続的インテグレーション）
- コード変更のたびに自動ビルド・テスト実行
- PyTestによる自動テスト
- 品質の向上とバグの早期発見

#### CD（継続的デリバリー）
- Google Cloud Run への自動デプロイ
- Dockerイメージの自動ビルド・プッシュ
- Workload Identity による安全な認証

#### 実装例：PyTestワークフロー
- `actions/checkout@v4` でリポジトリチェックアウト
- `actions/setup-python@v5` でPython環境構築
- requirements.txt によるライブラリインストール
- pytest コマンドでテスト実行

### ⚙️ 高度な機能と構文

#### 環境変数とシークレット
- `env:` による環境変数設定
- GitHub Secrets による機密情報管理
- `${{secrets.SECRET_NAME}}` での参照
- Repository Secrets と Environment Secrets

#### 条件分岐とステップ間連携
- `if: failure()` による失敗時処理
- ステップID設定とアウトプット受け渡し
- GITHUB_OUTPUT による値の共有

#### 実行環境
- **ubuntu-latest**（Linux）
- **windows-latest**（Windows）
- **macos-latest**（macOS）
- セルフホステッドランナーの選択肢

### ☁️ Google Cloud との連携

#### Workload Identity 設定
- 外部からGoogle Cloudへの安全なアクセス
- プール・プロバイダーの作成
- OIDC認証プロトコルの使用
- 属性マッピングと条件設定

#### サービスアカウントと権限
- Cloud Run管理者権限
- Artifact Registry書き込み権限
- サービスアカウントユーザー権限
- Workload Identityユーザー権限

#### デプロイフロー
- `google-github-actions/auth@v2` による認証
- Artifact Registry への Docker認証
- Dockerイメージのビルド・プッシュ
- `google-github-actions/deploy-cloudrun@v2` によるデプロイ

### 💡 ベストプラクティス

- 定期実行の信頼性に注意（GitHubのリソース不足で実行されない場合あり）
- 重要な定期ジョブは他のサービスとの併用を検討
- Gitハッシュ値をDockerタグに使用（`${{github.sha}}`）
- mainブランチへの直接プッシュ制限を設定

## 📊 実例・証拠

### 🔄 CI/CD自動化の効果
- **品質向上**: PyTestによる自動テストでバグの早期発見
- **デプロイ効率**: Google Cloud Runへの自動デプロイで手作業排除
- **コスト削減**: パブリックリポジトリでの基本無料利用

### 🛡️ セキュリティ強化
- **Workload Identity**: OIDC認証による安全なクラウド連携
- **GitHub Secrets**: 機密情報の安全な管理
- **権限管理**: 細かい粒度でのサービスアカウント権限設定

### ⚡ 開発効率化
- **YAMLベース**: 簡潔な設定ファイルによる高い可読性
- **GitHubエコシステム**: プルリクエストとの密接な統合
- **豊富なアクション**: actions/setup-python等の公式アクション活用

## ❓ 派生する問い
- 大規模チーム開発における複雑なワークフロー設計とマトリックスビルド活用方法は？
- セキュリティ要件の厳しい企業環境でのGitHub Actions運用ベストプラクティスは？
- マイクロサービスアーキテクチャでの効率的なCI/CDパイプライン構築戦略は？

## 🏷️ タグ

- note
- GitHub Actions
- CI/CD
- 自動化
- PyTest
- Google Cloud
- Workload Identity
- Docker
- DevOps
- YAML