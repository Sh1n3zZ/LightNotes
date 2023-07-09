<!DOCTYPE html>
<html>
<head>
    <title>用户注册</title>
</head>
<body>
    <h2>用户注册</h2>
    <form action="./backend/register_process.php" method="POST">
        <label for="username">用户名：</label>
        <input type="text" id="username" name="username" required><br><br>
        <label for="password">密码：</label>
        <input type="password" id="password" name="password" required><br><br>
        <input type="submit" value="注册">
    </form>
</body>
</html>
