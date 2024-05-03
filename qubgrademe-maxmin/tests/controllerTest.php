<?php
use PHPUnit\Framework\TestCase;

final class ControllerTest extends TestCase {
    // Test a successful request
    public function testController_successfulRequest() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "Cloud Computing",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_2" => "74",
            "mark_3" => "68",
            "mark_4" => "87",
            "mark_5" => "56",
        );
        echo "\n\n";
        $output = require($_SERVER['DOCUMENT_ROOT']."src/index.php");
        echo "\n\n";

        // Assertions
        $this->assertEquals($output["error"], false);
        $this->assertEquals($output["status"], 200);
        $this->assertEquals($output["max_module"], "Concurrent Programming - 87");
        $this->assertEquals($output["min_module"], "Programming - 32");
    }

    // Test a request with a missing parameter
    public function testController_missingParameterFailure() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "Cloud Computing",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_3" => "68",
            "mark_4" => "87",
            "mark_5" => "56",
        );
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");

        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "Mark 2 value is missing");
    }

    // Test a request with no parameters
    public function testController_noParameterFailure() {
        // Send request
        $_REQUEST = array();
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");

        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "Module 1 value is missing");
    }

    // Test a request with invalid input
    public function testController_invalidParameterFailure() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "Cloud Computing",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_2" => "74",
            "mark_3" => "68",
            "mark_4" => "asdasdas",
            "mark_5" => "56",
        );
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");
    
        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "You must provide a valid integer for Mark 4");
    }

    // Test a request with negative value
    public function testController_negativeParameterFailure() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "Cloud Computing",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_2" => "74",
            "mark_3" => "68",
            "mark_4" => "-34",
            "mark_5" => "56",
        );
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");
        
        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "You must provide a non-negative integer for Mark 4");
    }

    // Test a request with parameter value exceeding 100
    public function testController_largeParameterFailure() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "Cloud Computing",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_2" => "74",
            "mark_3" => "9880",
            "mark_4" => "34",
            "mark_5" => "56",
        );
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");
        
        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "You cannot exceed 100 marks for Mark 3");
    }

    // Test a request with whitespace module name
    public function testController_whitespaceModuleFailure() {
        // Send request
        $_REQUEST = array(
            "module_1" => "Programming",
            "module_2" => "Databases",
            "module_3" => "      ",
            "module_4" => "Concurrent Programming",
            "module_5" => "Cyber Security",
            "mark_1" => "32",
            "mark_2" => "74",
            "mark_3" => "68",
            "mark_4" => "87",
            "mark_5" => "56",
        );
        echo "\n\n";
        $errorResponse = require($_SERVER['DOCUMENT_ROOT']."src/index.php");

        // Assertions
        $this->assertEquals($errorResponse["error"], true);
        $this->assertEquals($errorResponse["status"], 400);
        $this->assertEquals($errorResponse["string"], "Module 3 value is missing");
    }
}