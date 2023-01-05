<?php

$nama_kota = ["Jakarta", "Aceh", "Malang", "Medan", "Bontang", "Jogja", "Jakarta", "Bandung", "Malang", "Solo", "Palembang", "Bandung"];

$unik_array = array_unique($nama_kota);

echo "<h2>Program Menghilangkat Array Duplikat</h2>";
echo "<h3>List Duplikat</h3>";
foreach ($nama_kota as $key => $value) {
  echo '<ul><li>'."$value".'</li></ul>';
}

echo "<h3>List Hasil Filter</h3>";
foreach ($unik_array as $key => $value) {
  echo '<ul><li>'."$value".'</li></ul>';
}
