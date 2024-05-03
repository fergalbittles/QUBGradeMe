<?php 
$config = file_get_contents("./config.json");
?>

<!DOCTYPE html>
<html>
<head>
    <title>QUB GradeMe Metrics</title>
    <script> var CONFIG = <?php echo $config; ?>;</script>
    <link rel="stylesheet" href="./metrics.css">
    <link rel="stylesheet" href="./style.css">
</head>

<body>
    <div id="qga">    
        <div class="inputoutput-wrapper">
        <div id="logo">
            QUB GradeMe Metrics
        </div>
        <p>Please enter admin credentials to continue</p>
        <div>
            <input class="display-module" type="text" id="username" placeholder="Username">
            <br>
            <input style="margin-top: 15px;" class="display-module" type="password" id="password" placeholder="Password">
            <br>
            <button style="margin-top: 15px;" id="login-button" class="qgabutton-active" onclick="if (validateLogin()) { signIn(); }">Sign In</button>
        </div>

        <p style="color: red;" id="login_error"></p>
    </div>

    <a href="./index.php">QUB GradeMe App</a>

    <script type="text/javascript" src="./metrics.js"></script>
</body>
</html>