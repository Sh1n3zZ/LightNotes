<?php
session_start();

// 导入数据库连接文件
require_once "./database/db_connection.php";

// 获取提交的表单数据
$username = $_POST["username"];
$password = $_POST["password"];

// 查询匹配的用户记录
$sql = "SELECT * FROM users WHERE username = '$username'";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    $row = $result->fetch_assoc();
    // 验证密码是否匹配
    if (password_verify($password, $row["password"])) {
        // 验证通过，将用户信息存储在会话中
        $_SESSION["user_id"] = $row["id"];
        $_SESSION["username"] = $username;
        // 跳转到创建便签的标签页
        header("Location: create_note.php");
        exit();
    } else {
        echo "密码不正确！";
    }
} else {
    echo "用户不存在！";
}

// 关闭数据库连接
$conn->close();
?>
