# 🗂️ Obsidian ディレクトリ構造の例

## 🎯 中心的な主張
**Obsidianの効果的なディレクトリ構造は、数字プレフィックス、階層化、indexフォルダの組み合わせにより、プログラマーの情報管理を体系的かつ柔軟に支援し、リンク機能と組み合わせることで厳密なフォルダ構造に縛られない知識管理を実現する。**

## 📖 詳細な説明

動画内で紹介されていた考え方（数字プレフィックス、階層化、indexフォルダ）に基づいた、プログラマー向けの具体的なディレクトリ構造例です。これはあくまで一例であり、ご自身の使い方に合わせて調整することが重要です。

### 📁 プログラマー向けディレクトリ構造例

```
Obsidian Vault/
├── 000 index/             # 新規ノート、一時的なメモ、未整理のノート
│   ├── Meeting Memo 2023-10-27.md
│   └── Quick Idea - New Feature.md
│
├── 100 Projects/          # 進行中・完了したプロジェクト関連
│   ├── Project Alpha/
│   │   ├── Requirements.md
│   │   ├── Architecture Design.md
│   │   └── Meetings/
│   │       └── Kickoff Meeting Notes.md
│   ├── Project Beta/
│   │   └── Setup Guide.md
│   └── Common Library/
│       └── API Documentation.md
│
├── 200 Learning/          # 技術学習、調査、コンセプト
│   ├── Programming Languages/
│   │   ├── Python/
│   │   │   └── Asyncio Basics.md
│   │   └── JavaScript/
│   │       ├── Promises vs Observables.md
│   │       └── Module System Deep Dive.md
│   ├── Cloud Services/
│   │   ├── AWS/
│   │   │   └── EC2 Instance Types Compared.md
│   │   └── Azure/
│   │       └── Functions - Trigger Types.md
│   ├── Concepts/
│   │   ├── SOLID Principles Explained.md
│   │   └── REST API Best Practices.md
│   └── Tutorials/
│       └── Docker Compose Tutorial Notes.md
│
├── 300 Work/              # 業務関連（プロジェクト以外）、チーム情報
│   ├── Team Processes/
│   │   └── Code Review Guidelines.md
│   │   └── Deployment Checklist.md
│   ├── Company Info/
│   │   └── Internal Tools Links.md
│   └── Snippets/            # よく使うコード片など
│       └── Useful SQL Queries.md
│
├── 400 Personal/          # 個人的なメモ、アイデア、読書記録など
│   ├── Blog Drafts/
│   │   └── My Obsidian Workflow.md
│   └── Book Notes/
│       └── Designing Data-Intensive Applications Summary.md
│
└── 999 Archive/           # 古くなった情報、完了して参照頻度が低いもの
    └── Old University Notes/
        └── Operating Systems Lecture 1.md
```

### 🔑 構造のポイント

#### 🔢 数字プレフィックス（000, 100, 200...）
フォルダが意図した順序で表示されやすくなります。また、動画のように`200`と検索すれば、"Learning"関連のフォルダやノートを素早く絞り込めます。数字の間隔を空けておくと、後で新しいカテゴリを追加しやすくなります。

#### 📥 000 indexフォルダ
新しく作成したノートや、まだ分類が決まっていない一時的なメモを置く場所です。動画の発表者は、ここのノートはあまり整理せず、主にノート間のリンクで辿れるようにしていると述べていました。

#### 🏗️ 階層構造
大分類（例: `Projects`, `Learning`）の中に、さらに具体的な中分類・小分類（例: `Programming Languages` -> `Python`）を作成することで、ノートを体系的に整理できます。ただし、階層が深くなりすぎないよう注意が必要です。

#### ⚖️ 具体性と汎用性
フォルダ名は、自分が理解しやすい具体的でありながら、ある程度の汎用性も持たせると良いでしょう（例：「AWS」だけでなく「Cloud Services/AWS」）。

#### 🔄 柔軟性
これはあくまで一例です。動画でも「フォルダ構造も考えすぎないように」と述べられています。自分のノートの取り方や内容に合わせて、自由に変更・追加・削除してください。リンク機能を活用すれば、厳密なフォルダ構造に縛られすぎる必要はありません。

#### 📦 999 Archiveフォルダ
定期的に見直し、古くなったり参照頻度が極端に低くなったノートを移動させる場所です。メインのフォルダ構造をすっきりと保つのに役立ちます。

## 📊 実例・証拠

### 🎯 効率的な情報分類
- **プロジェクト管理**: 100番台でアクティブな開発案件を集約
- **学習管理**: 200番台で技術習得を体系化
- **業務効率化**: 300番台で日常業務の知識を蓄積

### 🔍 検索とナビゲーションの最適化
- **数字検索**: 「200」で学習関連を瞬時に絞り込み
- **階層的アクセス**: 「Cloud Services」→「AWS」の段階的探索
- **リンク活用**: フォルダ構造に依存しない知識の関連付け

### 📈 継続的なメンテナンス
- **indexフォルダ**: 新規作成ノートの一時保管場所
- **Archive機能**: 古い情報の整理によるメイン構造の保持
- **柔軟な拡張**: 新カテゴリ追加時の番号体系の活用

## ❓ 派生する問い
- プロジェクト終了時のノート移動とアーカイブのベストプラクティスは？
- チーム開発でのObsidianディレクトリ構造の共通化は可能か？
- 大量の技術ドキュメント管理でのフォルダ階層の深さの最適解は？

## 🏷️ タグ

- note
- Obsidian
- 情報整理
- ディレクトリ構造
- 知識管理
- PKM
- 生産性向上