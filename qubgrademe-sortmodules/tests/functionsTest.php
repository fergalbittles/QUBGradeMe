<?php
require($_SERVER['DOCUMENT_ROOT']."src/functions.inc.php");
use PHPUnit\Framework\TestCase;

final class FunctionsTest extends TestCase { 
    // Test with valid input
    public function testGetSortedModules_successfulCalculation() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules[0]["module"], "Concurrent Programming");
        $this->assertEquals($sorted_modules[0]["marks"], "87");

        $this->assertEquals($sorted_modules[1]["module"], "Databases");
        $this->assertEquals($sorted_modules[1]["marks"], "74");

        $this->assertEquals($sorted_modules[2]["module"], "Cloud Computing");
        $this->assertEquals($sorted_modules[2]["marks"], "68");

        $this->assertEquals($sorted_modules[3]["module"], "Cyber Security");
        $this->assertEquals($sorted_modules[3]["marks"], "56");

        $this->assertEquals($sorted_modules[4]["module"], "Programming");
        $this->assertEquals($sorted_modules[4]["marks"], "32");
    } 

    // Test with empty array as input
    public function testGetSortedModules_emptyArrayFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array();

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with null as input
    public function testGetSortedModules_nullParamFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = null;

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with array which contains null values
    public function testGetSortedModules_nullArrayValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", null, "Cyber Security");
        $marks = array(32, 74, null, 87, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with array which contains empty string values
    public function testGetSortedModules_emptyStringValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with array which contains a value with only whitespace
    public function testGetSortedModules_whitespaceValueFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "         ", "Cyber Security");
        $marks = array(32, 74, 68, 87, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with non integer mark value
    public function testGetSortedModules_nonIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, true, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with negative integer mark value
    public function testGetSortedModules_negativeIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, -40, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 

    // Test with large integer mark value
    public function testGetSortedModules_largeIntegerFailure() {
        $modules = array("Programming", "Databases", "Cloud Computing", "Concurrent Programming", "Cyber Security");
        $marks = array(32, 74, 68, 988, 56);

        // Perform calculation
        $sorted_modules=getSortedModules($modules, $marks);

        // Assertions
        $this->assertEquals($sorted_modules, null);
    } 
}