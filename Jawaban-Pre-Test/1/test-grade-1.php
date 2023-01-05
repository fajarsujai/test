<?php

$point = $_POST['grade'];
$grade = NULL;
if ($point >= 90) {
  $grade = 'A';
}elseif($point >= 80 && $point <= 89 ){
  $grade = 'B';
}elseif ($point >= 70 && $point <= 79) {
  $grade = 'C';
}elseif ($point >= 60 && $point <= 69) {
  $grade = 'D';
}elseif ($point <= 59) {
  $grade = 'E';
}

echo '<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title></title>
  </head>
  <body>
    <h3>Your grade is <span style="color:green;">'."$grade".'</span></h3>
    <a href="form-grade.php">Reset</a>
  </body>
</html>';
