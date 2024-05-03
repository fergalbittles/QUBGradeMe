<?php
require($_SERVER['DOCUMENT_ROOT']."src/functions.inc.php");
use PHPUnit\Framework\TestCase;

final class FunctionsTest extends TestCase { 
    // Test with valid input
    public function testGetMaxMin_successfulCalculation() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules[0], "Concurrent Programming - 87");
        $this->assertEquals($max_min_modules[1], "Programming - 32");
    } 

    // Test with empty array as input
    public function testGetMaxMin_emptyArrayFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array();

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with null as input
    public function testGetMaxMin_nullParamFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = null;

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with array which contains null values
    public function testGetMaxMin_nullArrayValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", null, "Cyber Security");
        $marks = array(32, 74, null, 87, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with array which contains empty string values
    public function testGetMaxMin_emptyStringValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with array which contains a value with only whitespace
    public function testGetMaxMin_whitespaceValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "         ", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with non integer mark value
    public function testGetMaxMin_nonIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, true, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with negative integer mark value
    public function testGetMaxMin_negativeIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, -40, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 

    // Test with large integer mark value
    public function testGetMaxMin_largeIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, 988, 56);

        // Perform calculation
        $max_min_modules=getMaxMin($modules, $marks);

        // Assertions
        $this->assertEquals($max_min_modules, null);
    } 
}