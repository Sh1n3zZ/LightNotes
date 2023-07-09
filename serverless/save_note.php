<?php
session_start();

// 导入数据库连接文件
require_once "./database/db_connection.php";
require_once "filter.php";

// 检查用户是否已登录
if (!isset($_SESSION["user_id"])) {
    // 未登录，跳转回登录界面
    header("Location: login.php");
    exit();
}

// 获取当前登录的用户ID和用户名
$user_id = $_SESSION["user_id"];
$username = $_SESSION["username"];

// 获取用户输入的便签内容
$note_content = $_POST["note_content"];

// 过滤用户输入的便签内容
$filtered_note_content = sanitizeInput($note_content);

// 检查过滤后的内容是否与原内容相同，如果不同则提示用户
if ($note_content !== $filtered_note_content) {
    echo "不能输入以下内容：$filtered_note_content";
    exit();
}

// 执行插入便签数据的SQL语句
$sql = "INSERT INTO notes (user_id, username, content) VALUES ('$user_id', '$username', '$filtered_note_content')";

if ($conn->query($sql) === TRUE) {
    echo "便签保存成功！";
} else {
    echo "便签保存失败: " . $conn->error;
}

// 关闭数据库连接
$conn->close();
?>
