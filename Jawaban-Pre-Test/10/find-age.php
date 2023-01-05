<?php
$file = 'data.json'; // path to your JSON file
$data = file_get_contents($file); // put the contents of the file into a variable
$peoples = json_decode($data); // decode the JSON feed


echo "<h2>Program mencari orang-orang dibarawah 21 tahun </h2>";
echo "<h3>Daftar orang dibawah 21 tahun</h3>";

foreach ($peoples as $people) {
  if ($people->age < 21) {
    echo '<ul><li>ID : '."$people->id".' <br>Nama : '."$people->name".'<br>Age : '."$people->age".'</li></ul>';
  }
}
