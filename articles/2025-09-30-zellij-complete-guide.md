# 🚀 Zellij: 次世代ターミナルマルチプレクサー完全ガイド

## 🎯 中心的な主張

Zellijは、Rust製の次世代ターミナルマルチプレクサーで、使いやすさと高機能のバランスに優れています。最大の特徴は、画面下部に常時表示されるキーバインドヘルプにより、初心者でも学習曲線を容易に乗り越えられる点です。tmuxの「学習コストの高さ」という課題を解決しつつ、フローティングペインや自動最適配置などの独自機能で、さらなる生産性向上を実現します。

---

## 📊 Zellijの概要

### 🔍 基本情報

```
プロジェクト名: Zellij
言語: Rust
種類: ターミナルマルチプレクサー
ライセンス: MIT
対応OS:
  ├─ Linux
  ├─ macOS
  └─ Windows (WSL)

コンセプト:
  「初期状態でも使いやすく、
   パワーユーザーにも十分な機能」
```

### 💡 ターミナルマルチプレクサーとは？

```
定義:
  └─ 1つの端末画面内で複数のターミナルセッションを
     管理できるツール

できること:
  ├─ 画面分割（複数のプロセスを同時表示）
  ├─ タブ管理（作業コンテキストの切り替え）
  ├─ セッション保存（作業の中断・再開）
  └─ リモート作業の継続性（SSH切断後も復帰可能）

代表的なツール:
  ├─ tmux（伝統的な定番）
  ├─ GNU Screen（最古参）
  └─ Zellij（次世代）
```

---

## 🌟 Zellijの5大特徴

### 1️⃣ 分かりやすいUI

```
✨ 主要機能:
  └─ キーバインドが画面下部に常時表示

メリット:
  ├─ 📚 コマンドを覚える必要が大幅に減少
  ├─ 🎯 必要な操作が一目瞭然
  ├─ 🚀 学習曲線が極めて緩やか
  └─ 💡 ヘルプを探す時間を削減

対比:
  tmux → マニュアル必須、暗記が前提
  Zellij → 画面を見れば分かる
```

### 2️⃣ カスタマイズ可能なレイアウト

```
✨ 主要機能:
  └─ プロジェクトごとの作業環境を保存・再利用可能

メリット:
  ├─ 🏗️ プロジェクトの構造に合わせた最適配置
  ├─ ⚡ 環境セットアップ時間ゼロ
  ├─ 🔄 チーム間での設定共有
  └─ 📊 複雑な構成も一発で再現

活用例:
  ├─ Web開発: フロント・バック・DB・ログの4分割
  ├─ インフラ: 複数サーバーの監視画面
  └─ データ分析: コード・実行・可視化の3分割
```

### 3️⃣ 特殊なペイン配置

```
✨ 独自機能:

📌 フローティングペイン
  ├─ 他のペインの上に重なって表示
  ├─ マウスでドラッグして自由に移動
  ├─ 一時的な作業に最適
  └─ メイン作業を邪魔しない

📚 スタックペイン
  ├─ 複数のペインを同じ領域に重ねる
  ├─ タブのように切り替え可能
  ├─ 画面スペースを有効活用
  └─ 関連する作業のグループ化
```

### 4️⃣ マルチプレイヤー機能

```
✨ 主要機能:
  └─ 同じセッションに複数人が同時接続

メリット:
  ├─ 👥 ペアプログラミング
  ├─ 🎓 技術指導・トレーニング
  ├─ 🐛 リアルタイムデバッグ
  └─ 🤝 リモートトラブルシューティング

安全性:
  └─ 明示的な共有のみ（誤って公開される心配なし）
```

### 5️⃣ プラグインシステム

```
✨ 主要機能:
  └─ WebAssembly (WASM) を使用した拡張性

メリット:
  ├─ 🔧 任意の言語でプラグイン開発可能
  ├─ 🛡️ サンドボックス環境で安全
  ├─ ⚡ ネイティブに近いパフォーマンス
  └─ 🌐 クロスプラットフォーム対応

人気プラグイン例:
  ├─ ステータスバーカスタマイズ
  ├─ プロジェクトランチャー
  └─ Git統合
```

---

## 🆚 tmuxとの徹底比較

### 📊 基本比較表

| 項目 | tmux | Zellij |
|------|------|--------|
| **開発言語** | C | Rust |
| **初版リリース** | 2007年 | 2020年 |
| **安定性** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **学習難易度** | 高い | 低い |
| **UIアプローチ** | ミニマリスト | ヘルプバー常時表示 |
| **デフォルト設定** | ユーザーが構築 | バッテリー同梱 |
| **ペイン管理** | 手動設定 | 自動最適配置 |
| **プラグイン** | シェルスクリプト | WebAssembly |
| **マルチプレイヤー** | ❌ | ✅ |
| **フローティングペイン** | ❌ | ✅ |

### 🎯 詳細比較

#### 1. インターフェースの分かりやすさ

```
tmux:
  ❌ キーバインドは全て暗記必須
  ❌ man tmux を読む必要あり
  ❌ プレフィックスキー（Ctrl+b）の概念が独特
  ✅ ミニマリストには好評

Zellij:
  ✅ 画面下部に常にヘルプ表示
  ✅ モード名が明確
  ✅ 直感的な操作感
  ❌ ヘルプバーが邪魔という声も
```

#### 2. デフォルト設定の考え方

```
tmux:
  ├─ 最小限の機能のみ提供
  ├─ .tmux.confでカスタマイズ前提
  ├─ プレフィックスキーも変更推奨
  └─ 自分好みに育てる文化

Zellij:
  ├─ 初期状態で快適に使える
  ├─ 設定ファイル不要で起動可能
  ├─ 合理的なデフォルト値
  └─ 「すぐ使える」を重視
```

#### 3. ペイン管理の直感性

```
tmux:
  ├─ 手動でサイズ調整が必要
  ├─ レイアウトプリセットは固定的
  ├─ リサイズモードが別途必要
  └─ きめ細かい制御が可能

Zellij:
  ├─ 自動で最適サイズに調整
  ├─ ドラッグ&ドロップ対応
  ├─ フローティングペインで柔軟性向上
  └─ 直感的な操作
```

#### 4. 便利機能の充実度

```
tmux:
  ✅ 長年の実績による安定性
  ✅ 豊富なプラグインエコシステム
  ✅ あらゆる環境で動作
  ❌ 独自の便利機能は少ない

Zellij:
  ✅ フローティングペイン
  ✅ スタックペイン
  ✅ レイアウトシステム
  ✅ コマンドペイン
  ✅ マルチプレイヤー
  ❌ エコシステムは発展途上
```

### 🎯 選択の指針

```
tmuxを選ぶべき人:
  ✅ 長年tmuxを使っている
  ✅ 安定性を最優先
  ✅ ミニマリストな環境を好む
  ✅ 既存の.tmux.confがある
  ✅ 古いサーバーでも動作させたい

Zellijを選ぶべき人:
  ✅ これからマルチプレクサーを学ぶ
  ✅ モダンな機能を使いたい
  ✅ 設定ファイルを書きたくない
  ✅ 直感的な操作を重視
  ✅ ペアプログラミングをする
```

---

## 📥 インストール方法

### 🍎 macOS

```bash
# Homebrew
brew install zellij

# MacPorts
sudo port install zellij
```

### 🐧 Linux

```bash
# Ubuntu/Debian
# バイナリを直接ダウンロード
wget https://github.com/zellij-org/zellij/releases/latest/download/zellij-x86_64-unknown-linux-musl.tar.gz
tar -xzf zellij-x86_64-unknown-linux-musl.tar.gz
sudo mv zellij /usr/local/bin/

# Arch Linux
sudo pacman -S zellij

# Fedora
sudo dnf install zellij
```

### 🦀 Rust (cargo)

```bash
# 最新版をソースからビルド
cargo install zellij

# 開発版
cargo install --git https://github.com/zellij-org/zellij.git
```

### 🪟 Windows (WSL)

```bash
# WSL2環境でLinuxの手順に従う
```

### ✅ インストール確認

```bash
# バージョン確認
zellij --version

# 初回起動
zellij

# ヘルプ表示
zellij --help
```

### 💡 インストール不要で試す

```bash
# 一時的に実行（ダウンロード＆実行）
bash <(curl -L zellij.dev/launch)
```

---

## 🔄 モードシステム

Zellijは**モード**という概念で機能を整理しています。各モードに入ることで、特定の操作に特化したキーバインドが有効になります。

### 📊 モード構造の視覚化

```
🏠 ノーマルモード（起動時のデフォルト）
 │
 ├─ Ctrl+P → 🖼️ ペインモード
 ├─ Ctrl+T → 📑 タブモード
 ├─ Ctrl+N → 📏 リサイズモード
 ├─ Ctrl+O → 💾 セッションモード
 ├─ Ctrl+S → 🔍 スクロールモード
 ├─ Ctrl+G → 🔒 ロックモード
 └─ Ctrl+B → 🔧 tmuxモード
```

### 1️⃣ ノーマルモード

```
概要:
  └─ 基本的なナビゲーション・モード切替を行う

主要操作:
  ├─ Alt+矢印キー → ペイン間移動
  ├─ Alt+数字 → タブ切り替え
  ├─ Ctrl+各種 → モード切替
  └─ ESC → 他のモードから戻る

特徴:
  └─ 最もよく使う状態
```

### 2️⃣ ペインモード (Ctrl+P)

```
概要:
  └─ ペインの作成・削除・配置を管理

主要操作:
  ├─ N → 新規ペイン作成
  ├─ ↓ → 下に分割
  ├─ → → 右に分割
  ├─ Z → 最大化/復元
  ├─ X → ペイン削除
  ├─ W → フローティングペイン切替
  └─ E → ペイン名変更

使用頻度: ⭐⭐⭐⭐⭐
```

### 3️⃣ タブモード (Ctrl+T)

```
概要:
  └─ タブの管理（作業コンテキストの切り替え）

主要操作:
  ├─ N → 新規タブ作成
  ├─ ← / → → タブ間移動
  ├─ R → タブ名変更
  ├─ X → タブ削除
  └─ S → タブの並び替え

使用頻度: ⭐⭐⭐⭐
```

### 4️⃣ リサイズモード (Ctrl+N)

```
概要:
  └─ ペインのサイズを細かく調整

主要操作:
  ├─ ↑↓←→ → サイズ変更
  ├─ = → 均等サイズに
  └─ ESC → ノーマルモードに戻る

使用頻度: ⭐⭐⭐
```

### 5️⃣ セッションモード (Ctrl+O)

```
概要:
  └─ セッション管理・デタッチ

主要操作:
  ├─ D → デタッチ
  ├─ W → セッション一覧
  └─ C → 新規セッション作成

使用頻度: ⭐⭐⭐⭐
```

### 6️⃣ スクロールモード (Ctrl+S)

```
概要:
  └─ 出力履歴のスクロール・検索

主要操作:
  ├─ ↑↓ → スクロール
  ├─ PageUp/PageDown → 大きくスクロール
  ├─ / → 検索
  └─ ESC → ノーマルモードに戻る

使用頻度: ⭐⭐⭐
```

### 7️⃣ ロックモード (Ctrl+G)

```
概要:
  └─ ほとんどのキーバインドを無効化

主要操作:
  └─ Ctrl+G → 解除

使用シーン:
  ├─ Neovim使用時の競合回避
  ├─ 誤操作防止
  └─ アプリケーション側のキーを優先

使用頻度: ⭐⭐⭐⭐
```

### 8️⃣ tmuxモード (Ctrl+B)

```
概要:
  └─ tmux互換の操作感

主要操作:
  └─ tmuxのキーバインドが使用可能

対象ユーザー:
  └─ tmux経験者の移行を支援

使用頻度: ⭐⭐
```

---

## 🖼️ ペイン操作

### 📊 ペイン作成・分割

#### 基本操作

```
ペインモード起動: Ctrl+P

分割パターン:
  ├─ N → 新規ペイン（自動配置）
  ├─ ↓ → 現在のペインを下に分割
  ├─ → → 現在のペインを右に分割
  ├─ ← → 現在のペインを左に分割
  └─ ↑ → 現在のペインを上に分割
```

#### 視覚的な分割イメージ

```
初期状態:
┌────────────┐
│     A      │
└────────────┘

↓を押した場合:
┌────────────┐
│     A      │
├────────────┤
│     B      │
└────────────┘

→を押した場合（Bが選択状態）:
┌────────────┐
│     A      │
├──────┬─────┤
│  B   │  C  │
└──────┴─────┘
```

### 🎯 ペイン移動・選択

```
基本移動:
  Alt+矢印キー → ペイン間を移動

ペインモード内での移動:
  矢印キー → ペイン選択
  Enter → 選択して編集開始

サイクル移動:
  Tab → 次のペインへ
  Shift+Tab → 前のペインへ
```

### ⚡ ペイン操作

| 操作 | キーバインド | 説明 |
|------|-------------|------|
| **最大化/復元** | `Ctrl+P` → `Z` | トグル動作（再度押すと元に戻る） |
| **全画面化** | `Ctrl+P` → `F` | 他のペインを完全に隠す |
| **名前変更** | `Ctrl+P` → `E` | ペインに識別用の名前を付ける |
| **閉じる** | `Ctrl+P` → `X` | 確認なしで即座に閉じる |
| **移動** | `Ctrl+P` → `H/J/K/L` | Vim風の移動 |

### 🌟 フローティングペイン

```
概要:
  └─ 他のペインの上に重ねて表示される特殊ペイン

起動方法:
  Ctrl+P → W

特徴:
  ├─ マウスでドラッグ可能
  ├─ サイズ変更可能
  ├─ 通常ペインとトグル可能
  └─ 複数のフローティングペイン作成可能

活用シーン:
  ✅ メイン作業しながら別コマンドの出力確認
  ✅ ログをフローティング表示しながら開発
  ✅ ドキュメント参照しながらコーディング
  ✅ テスト結果確認しながらコード修正
  ✅ 一時的なメモ・スクラッチ領域
```

---

## 📑 タブ操作

### 🎯 タブの概念

```
タブとは:
  └─ 異なる作業コンテキストを切り替える仕組み

ブラウザのタブとの類似性:
  ├─ 複数のタブを同時に開ける
  ├─ タブ間を素早く切り替え
  ├─ 各タブに名前を付けて識別
  └─ タブごとに独立した作業環境

タブとペインの違い:
  タブ: 異なる作業の切り替え
  ペイン: 同じ作業内での画面分割
```

### 📊 基本操作

| 操作 | キーバインド | 説明 |
|------|-------------|------|
| **タブモード起動** | `Ctrl+T` | タブ操作モードに入る |
| **新規タブ作成** | タブモード → `N` | 新しいタブを作成 |
| **タブ間移動** | `Alt+数字(1-9)` | 指定番号のタブに移動 |
| **次のタブ** | タブモード → `→` | 右のタブに移動 |
| **前のタブ** | タブモード → `←` | 左のタブに移動 |
| **名前変更** | タブモード → `R` | タブ名を変更 |
| **並び替え** | タブモード → `S` | タブの順序を変更 |
| **閉じる** | タブモード → `X` | 現在のタブを削除 |
| **同期モード** | タブモード → `Y` | 全タブに同じ入力を送信 |

### 💡 タブ活用例

```
🌐 Web開発プロジェクト:
  ├─ Tab 0: "frontend" → React開発サーバー
  ├─ Tab 1: "backend"  → Go APIサーバー
  ├─ Tab 2: "database" → PostgreSQL接続
  ├─ Tab 3: "git"      → バージョン管理
  └─ Tab 4: "logs"     → 各種ログ監視

🏗️ インフラ管理:
  ├─ Tab 0: "web-01"   → Webサーバー1
  ├─ Tab 1: "web-02"   → Webサーバー2
  ├─ Tab 2: "db-01"    → DBサーバー
  └─ Tab 3: "monitor"  → 監視ダッシュボード

📊 データ分析:
  ├─ Tab 0: "notebook" → Jupyter Notebook
  ├─ Tab 1: "code"     → Pythonスクリプト
  ├─ Tab 2: "data"     → データ確認・加工
  └─ Tab 3: "viz"      → 可視化ツール
```

---

## 💾 セッション管理

### 🎯 セッションの概念

```
セッションとは:
  └─ Zellijの実行単位（tmuxと同じ概念）

重要性:
  ✅ SSH接続が切れても作業状態を保持
  ✅ 複数のプロジェクトを並行管理
  ✅ 作業の中断・再開が自由
  ✅ 長時間実行するタスクのバックグラウンド化
```

### 📊 基本操作（ターミナルで実行）

```bash
# 新規セッション（自動命名）
zellij

# 名前付きセッション作成
zellij --session my-project
# または
zellij -s my-project

# セッション一覧表示
zellij list-sessions
# または
zellij ls

# 特定セッションに接続
zellij attach my-project
# または
zellij a my-project

# 最後のセッションに再接続
zellij attach

# レイアウト指定で起動
zellij --layout ./layouts/web-dev.kdl

# セッション削除
zellij delete-session my-project
# または
zellij kill-session my-project

# 全セッション削除
zellij delete-all-sessions
```

### 📊 セッション内操作

```
セッションモード起動: Ctrl+O

主要操作:
  ├─ D → デタッチ（バックグラウンドに移行）
  ├─ W → セッション一覧表示・切り替え
  └─ C → 新規セッション作成

デタッチの動作:
  1. Ctrl+O → D でデタッチ
  2. Zellijが終了するが、セッションは残る
  3. zellij a で再接続
  4. 全ての状態が復元される
```

### 💡 セッション活用パターン

```
🎯 プロジェクトベース:
  ├─ zellij -s project-alpha
  ├─ zellij -s project-beta
  └─ zellij -s project-gamma

🎯 役割ベース:
  ├─ zellij -s development
  ├─ zellij -s testing
  ├─ zellij -s monitoring
  └─ zellij -s operations

🎯 顧客ベース:
  ├─ zellij -s client-a
  ├─ zellij -s client-b
  └─ zellij -s client-c

🎯 環境ベース:
  ├─ zellij -s local-dev
  ├─ zellij -s staging
  └─ zellij -s production
```

---

## 🔒 Neovimとの統合

### ⚠️ キーバインド競合問題

Zellijとエディタ（Neovim等）のキーバインドが競合する場合があります。

#### 主な競合例

```
Ctrl+O:
  ├─ Zellij: セッションモード
  └─ Neovim: ジャンプリスト（戻る）

Ctrl+N:
  ├─ Zellij: リサイズモード
  └─ Neovim: 補完候補（次へ）

Ctrl+P:
  ├─ Zellij: ペインモード
  └─ Neovim: 補完候補（前へ）
```

### ✅ 解決策

#### 方法1: ロックモードの活用（推奨）

```
手順:
  1. Neovim起動前に Ctrl+G でロックモード有効化
  2. Neovimを使用（Zellijのキーバインド無効）
  3. Neovim終了後に Ctrl+G で解除

メリット:
  ✅ 一時的な無効化が可能
  ✅ 設定ファイル不要
  ✅ 他のアプリでも使える

デメリット:
  ❌ 毎回手動で切り替えが必要
```

#### 方法2: Unlock-Firstプリセット

```
設定: ~/.config/zellij/config.kdl

keybinds {
    normal {
        // 最初にロック解除が必要
        bind "Ctrl g" { SwitchToMode "locked"; }
    }
}

メリット:
  ✅ デフォルトがロック状態
  ✅ Neovimとの競合を最小化

デメリット:
  ❌ Zellij操作時に毎回解除が必要
```

#### 方法3: tmuxモードの使用

```
使い方:
  └─ Ctrl+B でtmux互換モード

メリット:
  ✅ tmux経験者には馴染みやすい
  ✅ Ctrl+Oなどの競合を回避

デメリット:
  ❌ Zellijの独自機能が使いにくい
```

#### 方法4: カスタムキーバインド

```
設定: ~/.config/zellij/config.kdl

keybinds {
    normal {
        // ZellijのプレフィックスをAltに変更
        bind "Alt p" { SwitchToMode "pane"; }
        bind "Alt t" { SwitchToMode "tab"; }
        bind "Alt n" { SwitchToMode "resize"; }
        bind "Alt o" { SwitchToMode "session"; }
    }
}

メリット:
  ✅ 完全に競合を回避
  ✅ 自分好みにカスタマイズ

デメリット:
  ❌ 設定ファイルの記述が必要
  ❌ デフォルトとの乖離
```

### 🔌 Neovim統合プラグイン

```
人気プラグイン:
  ├─ zellij.nvim
  │   └─ Neovim内からZellijを制御
  ├─ smart-splits.nvim
  │   └─ ペイン移動をシームレスに
  └─ Navigator.nvim
      └─ ペインとNeovimバッファの統一ナビゲーション

インストール例 (lazy.nvim):
{
  "https://git.sr.ht/~swaits/zellij-nav.nvim",
  lazy = true,
  event = "VeryLazy",
  keys = {
    { "<c-h>", "<cmd>ZellijNavigateLeft<cr>",  { silent = true, desc = "navigate left"  } },
    { "<c-j>", "<cmd>ZellijNavigateDown<cr>",  { silent = true, desc = "navigate down"  } },
    { "<c-k>", "<cmd>ZellijNavigateUp<cr>",    { silent = true, desc = "navigate up"    } },
    { "<c-l>", "<cmd>ZellijNavigateRight<cr>", { silent = true, desc = "navigate right" } },
  },
  opts = {},
}
```

---

## ⚙️ 設定ファイル

### 📂 設定ファイルの場所

```
デフォルトパス:
  Linux/macOS: ~/.config/zellij/config.kdl
  Windows: %APPDATA%\zellij\config.kdl

設定ディレクトリ構造:
~/.config/zellij/
  ├─ config.kdl        # メイン設定
  ├─ layouts/          # レイアウトファイル
  │   ├─ default.kdl
  │   ├─ compact.kdl
  │   └─ custom.kdl
  ├─ themes/           # テーマファイル
  │   ├─ dracula.kdl
  │   └─ nord.kdl
  └─ plugins/          # プラグイン
      └─ *.wasm
```

### 🔧 KDL形式の特徴

```
KDL (KDL Document Language):
  ├─ Rust製の設定言語
  ├─ YAMLよりも読みやすい
  ├─ インデントエラーが発生しにくい
  ├─ 型安全
  └─ コメント対応

比較例:

YAML:
  keybinds:
    normal:
      - bind: "Ctrl p"
        action: "SwitchToMode"
        mode: "pane"

KDL:
  keybinds {
      normal {
          bind "Ctrl p" { SwitchToMode "pane"; }
      }
  }
```

### 📝 基本設定例

```kdl
// ~/.config/zellij/config.kdl

// デフォルトシェル
default_shell "zsh"

// マウス操作
mouse_mode true

// コピーモード
copy_command "pbcopy"  // macOS
// copy_command "xclip -selection clipboard"  // Linux

// スクロールバッファ
scroll_buffer_size 10000

// ペインフレーム
pane_frames false

// テーマ
theme "default"

// デフォルトレイアウト
default_layout "compact"

// 自動起動
// on_force_close "quit"

// セッションシリアライゼーション
session_serialization true

// スクロールバーの表示
scrollback_editor "nvim"
```

### 🎨 テーマカスタマイズ

```kdl
themes {
    my-theme {
        fg "#D8DEE9"
        bg "#2E3440"
        black "#3B4252"
        red "#BF616A"
        green "#A3BE8C"
        yellow "#EBCB8B"
        blue "#81A1C1"
        magenta "#B48EAD"
        cyan "#88C0D0"
        white "#E5E9F0"
        orange "#D08770"
    }
}
```

### ⌨️ キーバインドカスタマイズ

```kdl
keybinds {
    normal {
        // Altキーでペイン移動
        bind "Alt h" { MoveFocus "Left"; }
        bind "Alt j" { MoveFocus "Down"; }
        bind "Alt k" { MoveFocus "Up"; }
        bind "Alt l" { MoveFocus "Right"; }

        // Ctrl+Spaceでペインモード
        bind "Ctrl Space" { SwitchToMode "pane"; }
    }

    pane {
        // Vimライクな分割
        bind "v" { NewPane "Right"; SwitchToMode "normal"; }
        bind "s" { NewPane "Down"; SwitchToMode "normal"; }
    }
}
```

### 🔄 設定の反映

```bash
# 設定ファイル編集後、Zellijを再起動
# セッション内で:
Ctrl+O → W → 選択 → Enter

# またはデタッチして再接続
Ctrl+O → D
zellij a
```

---

## 🌟 便利機能

### 1️⃣ フローティングペイン（詳細版）

```
起動: Ctrl+P → W

特徴:
  ├─ 通常ペインの上に重ねて表示
  ├─ マウスでドラッグ＆リサイズ可能
  ├─ 複数のフローティングペインを同時表示
  ├─ 通常ペインとトグル可能
  └─ 独立したzインデックス管理

操作:
  ├─ マウスドラッグ → 移動
  ├─ マウスドラッグ（端） → リサイズ
  ├─ Ctrl+P → W → トグル（通常⇔フローティング）
  └─ Ctrl+P → X → 閉じる

活用シーン:
  ✅ メイン作業しながらログを常時監視
  ✅ ドキュメントを参照しながらコーディング
  ✅ テスト結果を確認しながらデバッグ
  ✅ 複数ファイルの内容を比較
  ✅ 一時的な計算やメモ
```

#### 実践例

```
シーン: Webアプリのデバッグ

レイアウト:
┌─────────────────────────────────┐
│  メインペイン: コードエディタ   │
│                                 │
│  ┌──────────────────┐          │
│  │ フローティング   │          │
│  │ サーバーログ     │          │
│  │                  │          │
│  └──────────────────┘          │
│                                 │
└─────────────────────────────────┘

メリット:
  ✅ コードを編集しながらログをリアルタイム確認
  ✅ 必要に応じてログペインを移動・拡大
  ✅ 不要な時は最小化
```

---

### 2️⃣ コマンドペイン

```
概要:
  └─ 特定のコマンドを繰り返し実行するための特殊ペイン

作成方法:

# ターミナルから
zellij run -- npm test

# セッション内から
zellij action new-pane --command "npm test"

特徴:
  ├─ Enterキーで同じコマンドを再実行
  ├─ 出力履歴が保持される
  ├─ コマンドが終了してもペインは残る
  └─ 通常のペインとしても使用可能
```

#### 活用シーン

```
✅ テストコマンドをEnterで簡単再実行
  └─ コード修正 → Enter → テスト結果確認

✅ ビルドコマンドの頻繁な実行
  └─ zellij run -- npm run build

✅ 開発サーバーの再起動
  └─ zellij run -- go run main.go

✅ システム監視コマンド
  └─ zellij run -- htop

✅ ログ監視
  └─ zellij run -- tail -f /var/log/app.log
```

#### 実践例

```
シーン: TDD（テスト駆動開発）

ワークフロー:
  1. コードを修正
  2. Enterキーでテスト再実行
  3. 結果を確認
  4. 1に戻る

設定:
# テスト用ペインを作成
zellij run -- npm test

# レイアウト:
┌─────────────┬─────────────┐
│    Code     │   Tests     │
│   Editor    │  (Enter to  │
│             │   rerun)    │
└─────────────┴─────────────┘

メリット:
  ✅ テスト実行が超高速
  ✅ ターミナルの切り替え不要
  ✅ 集中力が途切れない
```

---

### 3️⃣ レイアウトシステム

```
概要:
  └─ プロジェクトごとの作業環境を保存・再利用

保存場所:
  ~/.config/zellij/layouts/

使用方法:
  # 起動時に指定
  zellij --layout ~/.config/zellij/layouts/web-dev.kdl

  # 短縮形（layoutsディレクトリ内なら名前のみ）
  zellij --layout web-dev
```

#### レイアウトファイル例

```kdl
// ~/.config/zellij/layouts/web-dev.kdl

layout {
    pane size=1 borderless=true {
        plugin location="zellij:tab-bar"
    }

    pane split_direction="vertical" {
        // 左: エディタ
        pane size="60%" {
            name "editor"
        }

        // 右: サーバー＆ログ
        pane split_direction="horizontal" size="40%" {
            pane {
                name "server"
                command "npm"
                args "run" "dev"
            }
            pane {
                name "logs"
                command "tail"
                args "-f" "/var/log/app.log"
            }
        }
    }

    pane size=2 borderless=true {
        plugin location="zellij:status-bar"
    }
}
```

#### よく使うレイアウト例

```kdl
// シンプルな2分割
layout {
    pane split_direction="vertical" {
        pane size="50%"
        pane size="50%"
    }
}

// 3x3グリッド
layout {
    pane split_direction="horizontal" {
        pane split_direction="vertical" {
            pane
            pane
            pane
        }
        pane split_direction="vertical" {
            pane
            pane
            pane
        }
        pane split_direction="vertical" {
            pane
            pane
            pane
        }
    }
}

// IDE風レイアウト
layout {
    pane split_direction="vertical" {
        // メインエディタ領域
        pane size="70%" {
            name "main"
        }
        // サイドバー
        pane split_direction="horizontal" size="30%" {
            pane name="sidebar"
            pane name="terminal"
        }
    }
}
```

#### レイアウト活用戦略

```
🎯 プロジェクトタイプ別:
  ├─ web-frontend.kdl  → React/Vue開発
  ├─ web-backend.kdl   → API開発
  ├─ full-stack.kdl    → フルスタック
  ├─ data-science.kdl  → データ分析
  └─ devops.kdl        → インフラ管理

🎯 役割別:
  ├─ development.kdl   → 開発作業
  ├─ debugging.kdl     → デバッグ
  ├─ monitoring.kdl    → 監視
  └─ presentation.kdl  → デモ・プレゼン

🎯 チーム共有:
  └─ レイアウトファイルをGit管理し、
     チーム全体で統一された環境を共有
```

---

## 🎯 実践的なワークフロー例

### 🌐 ケース1: フルスタックWeb開発

```bash
# セッション作成
zellij --session web-project --layout web-full-stack

# レイアウト内容（web-full-stack.kdl）:
```

```kdl
layout {
    tab name="Frontend" {
        pane split_direction="vertical" {
            pane name="React Dev" {
                command "npm"
                args "run" "dev"
                cwd "~/projects/frontend"
            }
            pane size="40%" split_direction="horizontal" {
                pane name="Editor" {
                    cwd "~/projects/frontend"
                }
                pane name="Tests" {
                    command "npm"
                    args "test" "--" "--watch"
                }
            }
        }
    }

    tab name="Backend" {
        pane split_direction="vertical" {
            pane name="API Server" {
                command "go"
                args "run" "main.go"
                cwd "~/projects/backend"
            }
            pane size="40%" {
                name "Database"
                command "psql"
                args "-U" "postgres"
            }
        }
    }

    tab name="Logs" {
        pane split_direction="horizontal" {
            pane name="App Logs" {
                command "tail"
                args "-f" "~/projects/logs/app.log"
            }
            pane name="Error Logs" {
                command "tail"
                args "-f" "~/projects/logs/error.log"
            }
        }
    }

    tab name="Git" {
        pane cwd="~/projects"
    }
}
```

```
メリット:
  ✅ 一発で完全な開発環境が立ち上がる
  ✅ Alt+1, Alt+2 でフロント・バック切り替え
  ✅ 全ての情報が一画面に集約
```

---

### 🔍 ケース2: システム監視ダッシュボード

```bash
zellij --session monitoring --layout monitoring

# レイアウト内容:
```

```kdl
layout {
    pane split_direction="horizontal" {
        // 左列: リソース監視
        pane split_direction="vertical" size="33%" {
            pane name="CPU/Memory" {
                command "htop"
            }
            pane name="Disk" {
                command "watch"
                args "-n" "1" "df" "-h"
            }
        }

        // 中央列: ネットワーク＆プロセス
        pane split_direction="vertical" size="33%" {
            pane name="Network" {
                command "watch"
                args "-n" "1" "ss" "-tunapl"
            }
            pane name="Docker" {
                command "watch"
                args "-n" "2" "docker" "ps"
            }
        }

        // 右列: ログ監視
        pane split_direction="vertical" size="34%" {
            pane name="System Log" {
                command "tail"
                args "-f" "/var/log/syslog"
            }
            pane name="Application Log" {
                command "tail"
                args "-f" "/var/log/app/production.log"
            }
        }
    }
}
```

---

### 📊 ケース3: データサイエンスプロジェクト

```bash
zellij --session data-analysis --layout data-science
```

```kdl
layout {
    tab name="Notebook" {
        pane {
            command "jupyter"
            args "lab"
        }
    }

    tab name="Code" {
        pane split_direction="vertical" {
            pane size="70%" name="Editor"
            pane split_direction="horizontal" size="30%" {
                pane name="Python REPL" {
                    command "ipython"
                }
                pane name="Tests" {
                    command "pytest"
                    args "--watch"
                }
            }
        }
    }

    tab name="Data" {
        pane split_direction="vertical" {
            pane name="Data Browser" {
                command "vd"  # VisiData
            }
            pane name="SQL" {
                command "usql"
                args "postgres://localhost/mydb"
            }
        }
    }

    tab name="Viz" {
        pane {
            command "python"
            args "-m" "http.server" "8000"
        }
    }
}
```

---

## 💡 ベストプラクティス

### 🎯 命名規則

```
セッション名:
  ✅ プロジェクト名をそのまま使用
     例: zellij -s my-awesome-app

  ✅ 目的を明確に
     例: zellij -s prod-deploy

  ❌ 避けるべき名前
     - test, tmp, temp（削除されやすい）
     - session1, session2（識別しにくい）

タブ名:
  ✅ 短く明確に
     例: "editor", "server", "logs"

  ❌ 長すぎる名前
     例: "This is my development server tab"

ペイン名:
  ✅ 実行内容を反映
     例: "npm dev", "go run", "psql"
```

### ⚡ 効率的なワークフロー

```
1. プロジェクト開始時:
   └─ 専用レイアウトで起動
      zellij --layout my-project

2. 作業中断時:
   └─ デタッチして保存
      Ctrl+O → D

3. 作業再開時:
   └─ セッションに再接続
      zellij a my-project

4. 複数プロジェクト並行時:
   └─ セッション一覧で切り替え
      Ctrl+O → W
```

### 🔧 カスタマイズの順序

```
推奨順:
  1️⃣ デフォルト設定で使い始める
  2️⃣ 不便に感じた点をメモ
  3️⃣ レイアウトを作成
  4️⃣ キーバインドを調整
  5️⃣ テーマをカスタマイズ
  6️⃣ プラグインを導入

❌ 避けるべき:
  └─ 最初から全てカスタマイズ
     → デフォルトの良さに気付けない
```

### 📚 学習の進め方

```
Week 1: 基本操作
  ├─ ペイン分割・移動
  ├─ タブ作成・切り替え
  └─ デタッチ・アタッチ

Week 2: 便利機能
  ├─ フローティングペイン
  ├─ コマンドペイン
  └─ スクロールモード

Week 3: カスタマイズ
  ├─ 簡単なレイアウト作成
  ├─ キーバインド調整
  └─ テーマ変更

Week 4: 応用
  ├─ 複雑なレイアウト
  ├─ プラグイン導入
  └─ チーム共有設定
```

---

## 🐛 トラブルシューティング

### ❌ よくある問題と解決法

#### 問題1: Zellijが起動しない

```
症状:
  └─ zellij コマンドが見つからない

解決策:
  # パスの確認
  which zellij

  # パスが通っていない場合
  export PATH="$HOME/.cargo/bin:$PATH"

  # .bashrc / .zshrc に追加
  echo 'export PATH="$HOME/.cargo/bin:$PATH"' >> ~/.zshrc
```

#### 問題2: セッションが残り続ける

```
症状:
  └─ zellij ls で大量のセッションが表示

解決策:
  # 全セッション削除
  zellij delete-all-sessions

  # または個別削除
  zellij kill-session [セッション名]

  # 自動クリーンアップ設定
  # config.kdl に追加:
  on_force_close "quit"
```

#### 問題3: キーバインドが効かない

```
症状:
  └─ Ctrl+P などが反応しない

確認事項:
  1. ロックモードになっていないか
     → Ctrl+G で解除

  2. ターミナルがキーを横取りしていないか
     → iTerm2等の設定を確認

  3. Neovim等のアプリ内で使っていないか
     → ロックモードを活用
```

#### 問題4: 日本語表示が崩れる

```
症状:
  └─ 日本語が文字化けする

解決策:
  # ロケール設定確認
  locale

  # UTF-8に設定
  export LANG=ja_JP.UTF-8
  export LC_ALL=ja_JP.UTF-8

  # .bashrc / .zshrc に追加
  echo 'export LANG=ja_JP.UTF-8' >> ~/.zshrc
```

#### 問題5: マウスが効かない

```
症状:
  └─ マウスでペインをクリックしても反応しない

解決策:
  # config.kdl に追加
  mouse_mode true

  # ターミナルのマウスモード確認
  # iTerm2: Preferences → Profiles → Terminal
  #   → "Report mouse clicks & drags" を有効化
```

#### 問題6: パフォーマンスが遅い

```
症状:
  └─ 画面描画が遅い、重い

解決策:
  # スクロールバッファを減らす
  scroll_buffer_size 1000

  # アニメーションを無効化
  simplified_ui true

  # ペインフレームを無効化
  pane_frames false
```

---

## 🔗 参考リソース

### 📖 公式ドキュメント

```
公式サイト:
  └─ https://zellij.dev

GitHubリポジトリ:
  └─ https://github.com/zellij-org/zellij

ドキュメント:
  └─ https://zellij.dev/documentation/

コミュニティ:
  ├─ Discord: https://discord.gg/CrUAFH3
  └─ Reddit: r/zellij
```

### 🎨 テーマ・レイアウト集

```
公式テーマコレクション:
  └─ https://github.com/zellij-org/zellij/tree/main/zellij-utils/assets/themes

コミュニティレイアウト:
  └─ https://github.com/topics/zellij-layout

設定サンプル:
  └─ https://github.com/zellij-org/zellij/tree/main/example
```

### 🔌 プラグイン

```
公式プラグイン:
  ├─ zellij:tab-bar
  ├─ zellij:status-bar
  └─ zellij:strider (ファイルブラウザ)

サードパーティ:
  └─ https://github.com/topics/zellij-plugin
```

---

## 🆚 他ツールとの比較まとめ

| 項目 | tmux | GNU Screen | Zellij |
|------|------|------------|--------|
| **学習曲線** | 急 | 急 | 緩やか |
| **デフォルトUX** | ⭐⭐ | ⭐ | ⭐⭐⭐⭐⭐ |
| **安定性** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **カスタマイズ性** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| **モダンな機能** | ⭐⭐ | ⭐ | ⭐⭐⭐⭐⭐ |
| **エコシステム** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| **パフォーマンス** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **マルチプレイヤー** | ❌ | ✅ | ✅ |

---

## ✨ 300字要約

Zellijは、Rust製の次世代ターミナルマルチプレクサーで、使いやすさと高機能のバランスに優れています。最大の特徴は、画面下部に常時表示されるキーバインドヘルプにより、初心者でも学習曲線を容易に乗り越えられる点です。tmuxと比較して、「バッテリー同梱」のコンセプトで初期状態から快適に使用でき、フローティングペインや自動最適配置などの独自機能を備えています。モードシステムにより限られたキーで多機能にアクセス可能で、ペイン・タブ・セッション管理が直感的に行えます。Neovimとの統合時にはロックモード等でキーバインド競合を回避でき、KDL形式の設定ファイルで柔軟なカスタマイズが可能です。コマンドペインやレイアウトシステムにより、プロジェクトごとの作業環境を効率的に再現でき、開発生産性を大幅に向上させる強力なツールとなっています。