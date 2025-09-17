document.addEventListener('DOMContentLoaded', function() {
    // メールアドレスコピー機能
    document.querySelectorAll('a[href^="mailto:"]').forEach(function(link) {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const email = this.href.replace('mailto:', '');
            
            // クリップボードAPIのエラーハンドリング
            navigator.clipboard.writeText(email).then(() => {
                // 成功時の通知
                showNotification('コピーしました');
            }).catch(() => {
                // フォールバック：テキスト選択でコピー
                const textArea = document.createElement('textarea');
                textArea.value = email;
                textArea.style.position = 'fixed';
                textArea.style.opacity = '0';
                document.body.appendChild(textArea);
                textArea.select();
                document.execCommand('copy');
                document.body.removeChild(textArea);
                showNotification('コピーしました');
            });
        });
    });
    
    // 通知表示関数（アクセシビリティ対応）
    function showNotification(message) {
        const notification = document.createElement('div');
        notification.textContent = message;
        notification.className = 'copy-notification';
        notification.setAttribute('aria-live', 'polite');
        notification.setAttribute('role', 'status');
        document.body.appendChild(notification);
        
        setTimeout(() => notification.remove(), 2000);
    }

    // ハンバーガーメニュー機能
    const hamburgerButton = document.querySelector('.hamburger-button');
    const sidebar = document.querySelector('.sidebar');
    const mobileOverlay = document.querySelector('.mobile-overlay');

    if (hamburgerButton && sidebar && mobileOverlay) {
        function toggleMenu() {
            hamburgerButton.classList.toggle('active');
            sidebar.classList.toggle('active');
            mobileOverlay.classList.toggle('active');
            
            // body のスクロールを制御
            if (sidebar.classList.contains('active')) {
                document.body.style.overflow = 'hidden';
            } else {
                document.body.style.overflow = '';
            }
        }

        function closeMenu() {
            hamburgerButton.classList.remove('active');
            sidebar.classList.remove('active');
            mobileOverlay.classList.remove('active');
            document.body.style.overflow = '';
        }

        // ハンバーガーボタンクリック
        hamburgerButton.addEventListener('click', toggleMenu);

        // オーバーレイクリックでメニューを閉じる
        mobileOverlay.addEventListener('click', closeMenu);

        // メニューリンククリックでメニューを閉じる
        const navLinks = sidebar.querySelectorAll('.nav-link');
        navLinks.forEach(link => {
            link.addEventListener('click', closeMenu);
        });

        // ESCキーでメニューを閉じる
        document.addEventListener('keydown', function(e) {
            if (e.key === 'Escape' && sidebar.classList.contains('active')) {
                closeMenu();
            }
        });

        // ウィンドウリサイズ時にモバイルメニューを閉じる
        window.addEventListener('resize', function() {
            if (window.innerWidth > 768 && sidebar.classList.contains('active')) {
                closeMenu();
            }
        });
    }
});