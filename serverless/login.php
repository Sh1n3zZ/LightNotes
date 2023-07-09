<!DOCTYPE html>
<html>
<head>
    <title>用户登录</title>
</head>
<body>
    <h2>用户登录</h2>
    <form action="./backend/login_process.php" method="POST">
        <label for="username">用户名：</label>
        <input type="text" id="username" name="username" required><br><br>
        <label for="password">密码：</label>
        <input type="password" id="password" name="password" required><br><br>
        <input type="submit" value="登录">
    </form>
</body>
</html>
