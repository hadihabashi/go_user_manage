<?php
session_start();
if ($_SESSION['email'] != "") {
	$url = 'http://ADDRESS:3000/logout';
	$email = $_SESSION['email'];
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
		}else{
			if ($result == "OK"){
				$_SESSION['user_logged_in'] = False;
				$_SESSION['email'] = "" ;
				session_destroy();
				header('Location:index.php');
			}
		}
		
}else{
	echo '<script type="text/javascript">alert("Session Have Problem And Cant not Logout"); </script>';
	header('Location:index.php');
}


?>