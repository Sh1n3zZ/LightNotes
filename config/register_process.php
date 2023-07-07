<?php
// 导入数据库连接文件
require_once "./database/db_connection.php";

// 获取提交的表单数据
$username = $_POST["username"];
$password = $_POST["password"];

// 对密码进行加密处理，以增加安全性
$hashed_password = password_hash($password, PASSWORD_DEFAULT);

// 执行插入用户数据的SQL语句
$sql = "INSERT INTO users (username, password) VALUES ('$username', '$hashed_password')";

if ($conn->query($sql) === TRUE) {
    echo "注册成功！";
} else {
    echo "注册失败: " . $conn->error;
}

// 关闭数据库连接
$conn->close();
?>
