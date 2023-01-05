<?php

function cekPalindrom( $kalimat )
{
    // strip out whitespace
    $kalimat = str_replace( ' ', '', $kalimat );
    // return bool
    return $kalimat == strrev( $kalimat );
}

$kalimat = $_POST['kalimat'];

if( cekPalindrom( $kalimat ) == true )
{
    echo '<h2>'."$kalimat".'</h2>';
    echo '<h3>Kalimat diatas <span style="color:orange;">adalah Palindrom</span></h3><br><a href="form-palindrom.php">Reset</a>';
}
else
{
  echo '<h2>'."$kalimat".'</h2>';
  echo '<h3>Kalimat diatas <span style="color:red;">bukan Palindrom</span></h3><br><a href="form-palindrom.php">Reset</a>';
}
