<?php 
$config = file_get_contents("./config.json");
?>

<!DOCTYPE html>
<html>
<head>
    <title>QUB GradeMe Metrics</title>
    <script> var CONFIG = <?php echo $config; ?>;</script>
    <link rel="stylesheet" href="./metrics.css">
    <link rel="stylesheet" href="./style.css">
</head>

<body>
    <div id="qga">    
        <div class="inputoutput-wrapper">
        <div id="logo">
            QUB GradeMe Metrics
        </div>
        <p>Execute tests or analyse data from previous tests</p>
        <div>
            <button id="runtests-button" class="qgabutton-active" onclick="runTests();">Run Tests</button>
            <button id="analysedata-button" class="qgabutton-active" onclick="analyseTestData();">Analyse Test Data</button>
        </div>

        <p id="tests_passed_message"></p>
        <p id="server_error">Error: Unable to connect to a server</p>
        <p id="generic_error"></p>
        <p id="tests_are_running">Running tests, this may take a while...</p>
        <p id="logs_are_fetching">Fetching logs, this may take a while...</p>
    </div>

    <div id="chart-wrapper">
        <div id="chart-wrapper-inner">
            <p id="execution_time" style="margin-bottom: 5px;">Test Execution Time (s)</p>
            <p id="average_execution_time" style="margin-bottom: 5px;">Average Execution Time (s)</p>
            <canvas id="myChart"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-1">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-1">
            <p style="margin-bottom: 5px;">Proxy Route - Historic Execution Time</p>
            <canvas id="myGraph-1"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-2">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-2">
            <p style="margin-bottom: 5px;">Max/Min - Historic Execution Time</p>
            <canvas id="myGraph-2"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-3">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-3">
            <p style="margin-bottom: 5px;">Sort Modules - Historic Execution Time</p>
            <canvas id="myGraph-3"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-4">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-4">
            <p style="margin-bottom: 5px;">Total Marks - Historic Execution Time</p>
            <canvas id="myGraph-4"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-5">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-5">
            <p style="margin-bottom: 5px;">Average Mark - Historic Execution Time</p>
            <canvas id="myGraph-5"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-6">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-6">
            <p style="margin-bottom: 5px;">Overall Classify - Historic Execution Time</p>
            <canvas id="myGraph-6"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-7">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-7">
            <p style="margin-bottom: 5px;">Module Classify - Historic Execution Time</p>
            <canvas id="myGraph-7"></canvas>
        </div>
    </div>

    <div class="graph-wrapper" id="graph-wrapper-8">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-8">
            <p style="margin-bottom: 5px;">Database Save - Historic Execution Time</p>
            <canvas id="myGraph-8"></canvas>
        </div>
    </div>

    <div style="margin-bottom: 20px;" class="graph-wrapper" id="graph-wrapper-9">
        <div class="graph-wrapper-inner" id="graph-wrapper-inner-9">
            <p style="margin-bottom: 5px;">Database Load - Historic Execution Time</p>
            <canvas id="myGraph-9"></canvas>
        </div>
    </div>

    <div id="test-1-result" class="test-result-output">
        <div id="test-1-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Proxy - Route to Average Mark
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-1-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-1-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-1-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-1-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-1-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-1-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-1-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-2-result" class="test-result-output">
        <div id="test-2-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Highest & Lowest Marks
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-2-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-2-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-2-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-2-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-2-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-2-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-2-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-3-result" class="test-result-output">
        <div id="test-3-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Sort Modules
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-3-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-3-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-3-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-3-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-3-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-3-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-3-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-4-result" class="test-result-output">
        <div id="test-4-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Total Marks
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-4-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-4-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-4-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-4-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-4-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-4-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-4-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-5-result" class="test-result-output">
        <div id="test-5-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Average Mark
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-5-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-5-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-5-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-5-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-5-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-5-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-5-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-6-result" class="test-result-output">
        <div id="test-6-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Overall Classification
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-6-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-6-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-6-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-6-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-6-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-6-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-6-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-7-result" class="test-result-output">
        <div id="test-7-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Module Classification
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-7-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-7-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-7-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-7-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-7-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-7-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-7-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-8-result" class="test-result-output">
        <div id="test-8-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Database - Save Operation
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-8-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-8-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-8-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-8-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-8-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-8-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-8-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <div id="test-9-result" class="test-result-output">
        <div id="test-9-result-status" class="test-status">
            Fetching test result...
        </div>
        <div class="test-title">
            <b>Test Name:</b> Database - Load Operation
        </div>
        <table class="test-stats">
            <tr>
                <td><b>Date:</b> <span id="test-9-result-date">N/A</span></td>
                <td><b>Result:</b> <span id="test-9-result-result">N/A</span></td>
                <td><b>Time Taken (s):</b> <span id="test-9-result-time">N/A</span></td>
            </tr>
        </table>
        <table class="expected-actual">
            <tr>
                <td style="border-right: 2px solid black; border-bottom: 2px solid black;"><b>Expected Status:</b> <span id="test-9-result-ex-status">N/A</span></td>
                <td style="border-bottom: 2px solid black;"><b>Actual Status:</b> <span id="test-9-result-ac-status">N/A</span></td>
            </tr>
            <tr>
                <td style="border-right: 2px solid black;"><b>Expected Response:</b> <br><br><span id="test-9-result-ex-res">N/A</span></td>
                <td><b>Actual Response:</b> <br><br><span id="test-9-result-ac-res">N/A</span></td>
            </tr>
        </table>
    </div>

    <a href="./index.php">QUB GradeMe App</a>

    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <script type="text/javascript" src="./metrics.js"></script>
</body>
</html>
