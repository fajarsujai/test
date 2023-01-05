<?php


$string_1 = strtolower($_POST['string_1']);
$string_2 = strtolower($_POST['string_2']);

$string1_array = str_split($string_1);
$string2_array = str_split($string_2);

$result = array_intersect($string1_array,$string2_array);

if(count($string1_array) == count($result)){
    echo ''."$string_1".' = '."$string_2".'<br>
    <h1>True</h1> <br>
    <a href="index.php">Kembali</a>
    ';
}else{
    echo ''."$string_1".' = '."$string_2".'<br>
    <h1>False</h1> <br>
    <a href="index.php">Kembali</a>
    ';
}
