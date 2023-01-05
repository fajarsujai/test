<?php
$start_date = $_POST['input_tahun_1'];
$end_date = $_POST['input_tahun_2'];

$period = new DatePeriod(
     new DateTime($start_date),
     new DateInterval('P1Y'),
     new DateTime($end_date)
);

echo '<h3>List Tahun Kabisat</h3>';
$list_tahun;
foreach ($period as $key => $value) {
  $list_tahun = $value->format('Y');
  if ($list_tahun%4==0) {
    echo '<ul>
    <li>'."$list_tahun".'</li>
    </ul>';
  }
}

echo "<a href='form-kabisat.php'>Reset</a>";
