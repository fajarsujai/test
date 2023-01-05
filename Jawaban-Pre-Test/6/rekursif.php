<?php

$input_rekursif = $_POST['input_rekursif'];

function rekursif($input=1)
{
  $hasil = 1;
  for ($i=$input; $i >= 1 ; $i--) {
    $hasil = $hasil * $i;
  }

  return $hasil;
}

$hasil_rekursif = rekursif($input_rekursif);

echo '<h3>hasil rekursif dari '."$input_rekursif".'! = '."$hasil_rekursif".'</h3>';
echo '<br><a href="form-rekursif.php">Reset</a>';
