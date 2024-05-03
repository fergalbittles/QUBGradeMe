<?php

if (!headers_sent()) {
	header("Access-Control-Allow-Origin: *");
	header("Content-type: application/json");
}

require('functions.inc.php');

$output = array(
	"error" => false,
  	"modules" => "",
	"marks" => 0,
	"sorted_modules" => ""
);

$errorResponse = array(
	"error" => true,
	"string" => ""
);

// Parse the query params
$queryParams = array();
$queryParams['Module 1'] = array_key_exists("module_1", $_REQUEST) ? $_REQUEST['module_1'] : "";
$queryParams['Module 2'] = array_key_exists("module_2", $_REQUEST) ? $_REQUEST['module_2'] : "";
$queryParams['Module 3'] = array_key_exists("module_3", $_REQUEST) ? $_REQUEST['module_3'] : "";
$queryParams['Module 4'] = array_key_exists("module_4", $_REQUEST) ? $_REQUEST['module_4'] : "";
$queryParams['Module 5'] = array_key_exists("module_5", $_REQUEST) ? $_REQUEST['module_5'] : "";
$queryParams['Mark 1'] = array_key_exists("mark_1", $_REQUEST) ? $_REQUEST['mark_1'] : "";
$queryParams['Mark 2'] = array_key_exists("mark_2", $_REQUEST) ? $_REQUEST['mark_2'] : "";
$queryParams['Mark 3'] = array_key_exists("mark_3", $_REQUEST) ? $_REQUEST['mark_3'] : "";
$queryParams['Mark 4'] = array_key_exists("mark_4", $_REQUEST) ? $_REQUEST['mark_4'] : "";
$queryParams['Mark 5'] = array_key_exists("mark_5", $_REQUEST) ? $_REQUEST['mark_5'] : "";

// Perform error handling
foreach ($queryParams as $key => $value) {
	// Trim whitespace
	$queryParams[$key] = trim($queryParams[$key]);

	// Missing param error
	if (is_null($queryParams[$key]) || $queryParams[$key] == '') {
		$errorResponse['string'] = $key . " value is missing";
		http_response_code(400);
		echo json_encode($errorResponse);
		
		// Allow tests to recieve the response
		$errorResponse["status"] = 400;
		return $errorResponse;
	} 

	// Skip 'module' values
	if (strstr($key, 'Module ')) {
		continue;
	}

	// Invalid input error
	$isNotInt = !ctype_digit($queryParams[$key]);
	$isNotNegativeInt = !(substr($queryParams[$key], 0, 1) == '-' && ctype_digit(substr($queryParams[$key], 1)));

	if ($isNotInt && $isNotNegativeInt) {
		$errorResponse['string'] = "You must provide a valid integer for " . $key;
		http_response_code(400);
		echo json_encode($errorResponse);
		
		// Allow tests to recieve the response
		$errorResponse["status"] = 400;
		return $errorResponse;
	} 

	// Parse to int
	$queryParams[$key] = (int)$queryParams[$key];
	
	// Negative integer error
	if ($queryParams[$key] < 0) {
		$errorResponse['string'] = "You must provide a non-negative integer for " . $key;
		http_response_code(400);
		echo json_encode($errorResponse);
		
		// Allow tests to recieve the response
		$errorResponse["status"] = 400;
		return $errorResponse;
	}

	// Integer exceeds 100
	if ($queryParams[$key] > 100) {
		$errorResponse['string'] = "You cannot exceed 100 marks for " . $key;
		http_response_code(400);
		echo json_encode($errorResponse);
		
		// Allow tests to recieve the response
		$errorResponse["status"] = 400;
		return $errorResponse;
	}
}

$modules = array($queryParams['Module 1'], $queryParams['Module 2'], $queryParams['Module 3'], $queryParams['Module 4'], $queryParams['Module 5']);
$marks = array($queryParams['Mark 1'], $queryParams['Mark 2'], $queryParams['Mark 3'], $queryParams['Mark 4'], $queryParams['Mark 5']);

// Perform calculation
$sorted_modules=getSortedModules($modules, $marks);

// Ensure calculation was succesful
if ($sorted_modules == null) {
	$errorResponse['string'] = "Error while performing calculation, ensure all input is valid";
	http_response_code(400);
	echo json_encode($errorResponse);
	
	// Allow tests to recieve the response
	$errorResponse["status"] = 400;
	return $errorResponse;
}

// Return response
$output['modules']=$modules;
$output['marks']=$marks;
$output['sorted_modules']=$sorted_modules;

http_response_code(200);
echo json_encode($output);

// Allow tests to recieve the response
$output["status"] = 200;
return $output;
