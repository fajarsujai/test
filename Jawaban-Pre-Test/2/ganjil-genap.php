<?php

$angka = $_POST['ganjil_genap'];
$hasil = NULL;
function ganjilGenap($angka)
{
  if ($angka % 2 == 0) {
    $hasil = 'Genap';
  }else{
    $hasil = 'Ganjil';
  }
  return $hasil;
}

$ganjilGenap = ganjilGenap($angka);

echo '<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title></title>
  </head>
  <body>
    <h3>'."$angka".' Adalah bilangan '."$ganjilGenap".'</h3>
    <a href="form-ganjil-genap.php">Reset</a>
  </body>
</html>';
