<?php
$servername = "yourserverlocation";
$username = "databaseusername";
$password = "databasepassword";
$dbname = "databasename";

// 创建数据库连接
$conn = new mysqli($servername, $username, $password, $dbname);

// 检查连接是否成功
if ($conn->connect_error) {
    die("数据库连接失败: " . $conn->connect_error);
}
?>
