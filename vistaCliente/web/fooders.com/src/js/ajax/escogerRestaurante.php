<?php

//$consulta = $_GET['nameTable'];
$consulta="restaurantes";
//$apiUrl = 'https://c5u2jk30nd.execute-api.us-west-2.amazonaws.com/nuevo';
$apiUrl = 'http://10.20.0.201:8080';
$apiUrl.="/".$consulta;

//echo $apiUrl."\n";

if(ini_get('allow_url_fopen')){
    $json = file_get_contents($apiUrl);
} else{
    $curl = curl_init($apiUrl);
    curl_setopt($curl, CURLOPT_ENCODING ,"");
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
    $json = curl_exec($curl);
    //echo $json;
    print($json);
    curl_close($curl);
}

//echo gettype($json);
echo $json;


?>