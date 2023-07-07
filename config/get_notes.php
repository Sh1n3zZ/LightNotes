<?php
session_start();

// 导入数据库连接文件
require_once "./database/db_connection.php";

// 检查用户是否已登录
if (!isset($_SESSION["user_id"])) {
    // 未登录，返回空数组
    echo json_encode([]);
    exit();
}

// 获取当前登录的用户ID
$user_id = $_SESSION["user_id"];

// 获取页码和每页便签数量
$page = isset($_GET["page"]) ? intval($_GET["page"]) : 1;
$perPage = isset($_GET["perPage"]) ? intval($_GET["perPage"]) : 5;
$offset = ($page - 1) * $perPage;

// 查询当前用户的便签数量
$countSql = "SELECT COUNT(*) as total FROM notes WHERE user_id = $user_id";
$countResult = $conn->query($countSql);
$totalCount = $countResult->fetch_assoc()["total"];

// 计算总页数
$totalPages = ceil($totalCount / $perPage);

// 查询当前页的便签数据
$sql = "SELECT content FROM notes WHERE user_id = $user_id ORDER BY id DESC LIMIT $offset, $perPage";
$result = $conn->query($sql);

$notes = [];
if ($result->num_rows > 0) {
    while ($row = $result->fetch_assoc()) {
        $notes[] = $row;
    }
}

// 返回便签数据和总页数
$response = [
    "notes" => $notes,
    "totalPages" => $totalPages
];

echo json_encode($response);

// 关闭数据库连接
$conn->close();
?>
