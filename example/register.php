<?php
$url = 'http://ADDRESS:3000/register';
$email = $_POST['email'];
$pass = $_POST['pass'];
if ($_POST['email'] != "" && $_POST['pass'] !="" && $_POST['copass'] !="") {
	if ($_POST['pass'] == $_POST['copass']){
		$data = array('email' => $email , 'pass' => $pass );
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
			echo "ERROR ...";
		}
		//var_dump($result);
		
	}else{
		echo '<script type="text/javascript">alert("Password Do Not Math"); </script>';
	}
}
?>
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Moolin Monitoring System | Registration Page</title>
  
  <?php include 'includes/header.php'; ?>

</head>
<body class="hold-transition register-page">
<div class="register-box">
  <div class="register-logo">
    <a href="index.php"><b>Moolin</b></a>
  </div>

  <div class="register-box-body">
    <p class="login-box-msg">Register a new membership</p>

    <form  action="register1.php" method="post">
      <!--
	  <div class="form-group has-feedback">
        <input type="text" class="form-control" placeholder="Full name">
        <span class="glyphicon glyphicon-user form-control-feedback"></span>
      </div>
	  -->
      <div class="form-group has-feedback">
        <input type="email" name="email" class="form-control" placeholder="Email">
        <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
      </div>
      <div class="form-group has-feedback">
        <input type="password" name="pass" class="form-control" placeholder="Password">
        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
      </div>
      <div class="form-group has-feedback">
        <input type="password" name="copass" class="form-control" placeholder="Retype password">
        <span class="glyphicon glyphicon-log-in form-control-feedback"></span>
      </div>
      <div class="row">
        <div class="col-xs-8">
          <div class="checkbox icheck">
            <label>
              <input type="checkbox"> I agree to the <a href="#">terms</a>
            </label>
          </div>
        </div>
        <!-- /.col -->
        <div class="col-xs-4">
          <button type="submit" class="btn btn-primary btn-block btn-flat">Register</button>
        </div>
        <!-- /.col -->
      </div>
    </form>
	
	<!--
    <div class="social-auth-links text-center">
      <p>- OR -</p>
      <a href="#" class="btn btn-block btn-social btn-facebook btn-flat"><i class="fa fa-facebook"></i> Sign up using
        Facebook</a>
      <a href="#" class="btn btn-block btn-social btn-google btn-flat"><i class="fa fa-google-plus"></i> Sign up using
        Google+</a>
    </div>
	-->
	
    <a href="login.php" class="text-center">I already have a membership</a>
	
	
	<div id="results">		
			<?php echo $result?>
	</div>
	
	
	
  </div>
  <!-- /.form-box -->
</div>
<!-- /.register-box -->


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
