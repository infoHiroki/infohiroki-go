# 🚀 楽しいGAS用語集

Google Apps Scriptをマスターしよう！

Google Apps Script（GAS）の世界で使われる用語を楽しく覚えましょう！Googleサービスの自動化に役立つ用語をまとめています。

## 📚 基本用語

### 🚀 GAS (Google Apps Script)
Googleが提供するクラウドベースのスクリプトプラットフォーム。JavaScriptでGoogleサービスを自動化できる。

### ⚙️ スクリプトエディタ
GASのコードを書くためのブラウザ上の開発環境。リアルタイムで編集・実行できる。

### 📁 プロジェクト
GASのコードとリソースをまとめた単位。複数のファイルやライブラリを含むことができる。

### ⏰ トリガー
スクリプトを自動実行するための仕組み。時間ベースやイベントベースで実行できる。

## 🔗 Google サービス連携

### 📊 SpreadsheetApp
Google スプレッドシートを操作するためのAPIクラス。セルの読み書きや書式設定ができる。
```javascript
SpreadsheetApp.getActiveSheet()
```

### 📧 GmailApp
Gmailを操作するためのAPIクラス。メールの送信や受信ボックスの管理ができる。
```javascript
GmailApp.sendEmail()
```

### 💾 DriveApp
Google ドライブのファイルやフォルダを操作するためのAPIクラス。
```javascript
DriveApp.getFileById()
```

### 📅 CalendarApp
Google カレンダーのイベントやカレンダーを操作するためのAPIクラス。
```javascript
CalendarApp.createEvent()
```

## ⚡ 実行・デプロイ

### 🔧 関数
GASで実行できる処理の単位。function キーワードで定義し、スクリプトエディタから直接実行できる。

### ⚡ 実行
書いたスクリプトを動かすこと。手動実行とトリガーによる自動実行がある。

### 🚀 デプロイ
スクリプトをWebアプリやAPIとして公開すること。外部からアクセス可能になる。

### 📋 ログ
スクリプトの実行結果やエラー情報を記録したもの。console.log()で出力できる。

## 🎨 高度な機能

### 🌐 UrlFetchApp
外部のAPIやWebサイトにHTTPリクエストを送信するためのクラス。
```javascript
UrlFetchApp.fetch(url)
```

### 🎨 HtmlService
HTMLベースのユーザーインターフェースを作成するためのサービス。

### 🔐 認証
Googleサービスへのアクセス許可。初回実行時に権限の確認が必要。

### 📚 ライブラリ
他のGASプロジェクトで作成された再利用可能なコード集。

## ⚠️ エラー・トラブル

### ⚠️ 実行時間制限
GASには6分（360秒）の実行時間制限がある。長時間処理は分割が必要。

### 🚫 API制限
Googleサービスへのアクセス回数に制限がある。過度のアクセスは避ける。

### 🐛 デバッグ
スクリプトの問題を見つけて修正すること。ブレークポイントやログを活用。

### 🔄 バージョン管理
スクリプトの変更履歴を管理する機能。以前のバージョンに戻すことができる。

## 🏷️ タグ

- GAS
- Google Apps Script
- 用語集
- 自動化