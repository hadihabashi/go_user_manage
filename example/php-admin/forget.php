<?php
session_start();
if (isset($_SESSION['user_logged_in']) && $_SESSION['user_logged_in'] === TRUE) {
    header('Location:index.php');
}


if ($_POST['email'] != "") {
	$url = 'http://ADDRESS:3000/resetpass';
	$email = $_POST['email'];
	
		$data = array('email' => $email);
		$options = array(
			'http' => array(
				'header'  => "Content-type: application/x-www-form-urlencoded\r\n",
				'method'  => 'POST',
				'content' => http_build_query($data)
			)
		);
		$context  = stream_context_create($options);
		$result = file_get_contents($url, false, $context);
		if ($result === FALSE) {
			$result = "ERROR";
		}
		
}

?>

<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Moolin Monitoring System | Log in</title>
  
  <?php include 'includes/header.php'; ?>

</head>
<body class="hold-transition login-page">
<div class="login-box">
  <div class="login-logo">
    <a href="index.php"><b>Moolin</b></a>
  </div>
  <!-- /.login-logo -->
  <div class="login-box-body">
    <p class="login-box-msg">Forget Password</p>

    <form action="forget.php" method="post">
      <div class="form-group has-feedback">
        <input type="email" name="email" class="form-control" placeholder="Email">
        <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
      </div>
      <div class="row">   
        <div class="col-xs-8">
          <button type="submit" class="btn btn-primary btn-block btn-flat">Send New Password</button>
        </div>   
      </div>
    </form>
	
	<div class="row">
	<div id="results" class="loader">		
		<?php echo $result; ?>
	</div>
	</div>
	
  </div>
  <!-- /.login-box-body -->
</div>
<!-- /.login-box -->

<?php include 'includes/footer.php'; ?>
<script>
  $(function () {
    $('input').iCheck({
      checkboxClass: 'icheckbox_square-blue',
      radioClass: 'iradio_square-blue',
      increaseArea: '20%' // optional
    });
  });
</script>
</body>
</html>
