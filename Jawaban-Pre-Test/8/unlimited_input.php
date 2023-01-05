<?php
  $i = 1;
  $hasil=0;
    while($i > 0) {
      echo "Masukan Angka : ";
      $value = (int)trim(fgets(STDIN, 1024));
      $hasil += $value;
      if ($value == '=') {
        $i = false;
      }
    }
    echo $hasil;

?>
