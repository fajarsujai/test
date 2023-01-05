<?php
$kalimat = $_POST['kalimat'];

$balik_kata = strrev($kalimat);

echo '<h3>Kalimat Input : '."$kalimat".'</h3>';
echo '<h3>Hasil Balik Kata : '."$balik_kata".'</h3>';
echo '<a href="form-balik-kata.php">Reset</a>';
