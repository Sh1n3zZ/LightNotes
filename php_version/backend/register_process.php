<?php
// 导入数据库连接文件
require_once "../database/db_connection.php";

// 获取提交的表单数据
$username = $_POST["username"];
$password = $_POST["password"];

// 对密码进行加密处理，以增加安全性
$hashed_password = password_hash($password, PASSWORD_DEFAULT);

// 执行插入用户数据的SQL语句
$sql = "INSERT INTO users (username, password) VALUES ('$username', '$hashed_password')";

if ($conn->query($sql) === TRUE) {
    // 注册成功
    echo "注册成功！";

    // 将当前账号的 Cookie 写入
    $cookie_name = "user_login";
    $cookie_value = $conn->insert_id; // 使用插入语句返回的自增ID作为用户标识
    $cookie_expire = time() + (30 * 24 * 60 * 60); // 当前时间 + 30 天
    setcookie($cookie_name, $cookie_value, $cookie_expire, "/");

    // 跳转至 index.php
    header("Location: ../index.php");
    exit();
} else {
    // 注册失败
    echo "注册失败: " . $conn->error;
}

// 关闭数据库连接
$conn->close();
?>
