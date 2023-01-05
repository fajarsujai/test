<?php

$angka_1 = $_POST['angka_1'];
$angka_2 = $_POST['angka_2'];
$angka_3 = $_POST['angka_3'];
$angka_4 = $_POST['angka_4'];
$angka_5 = $_POST['angka_5'];

$angka_array = [$angka_1,$angka_2,$angka_3,$angka_4,$angka_5];

$max = max($angka_array);
$min = min($angka_array);
$average = array_sum($angka_array)/count($angka_array);

echo '<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title></title>
  </head>
  <body>
      <h4>Angka Input : '."$angka_1, $angka_2, $angka_3, $angka_4, $angka_5".'</h4>
      <h3>Nilai Min : '."$max".'</h3>
      <h3>Nilai Min : '."$min".'</h3>
      <h3>Nilai Min : '."$average".'</h3>
      <a href="form-min-max-average.php">Reset</a>
  </body>
</html>';
