<?php
session_start();

// 检查用户是否已登录
if (isset($_SESSION["user_id"])) {
    // 用户已登录，跳转至其他页面（例如，create_note.php）
    header("Location: create_note.php");
    exit();
} else {
    // 用户未登录，检查是否存在持久登录状态的 Cookie
    if (isset($_COOKIE["user_login"])) {
        // 获取持久登录状态的信息，例如用户ID
        $user_id = $_COOKIE["user_login"];
        
        // TODO: 根据用户ID查询数据库或进行其他验证操作
        
        // 验证通过，将用户信息存储在会话中
        $_SESSION["user_id"] = $user_id;
        
        // 跳转至其他页面（例如，create_note.php）
        header("Location: create_note.php");
        exit();
    } else {
        // 用户未登录且没有持久登录状态的 Cookie，跳转至 register.php
        header("Location: register.php");
        exit();
    }
}
?>
