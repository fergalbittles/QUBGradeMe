<?php 
$config = file_get_contents("./config.json");
?>

<!DOCTYPE html>
<html>
<head>
    <title>QUB GradeMe</title>
    <script> var CONFIG = <?php echo $config; ?>;</script>
    <script type="text/javascript" src="./script.js"></script>
    <link rel="stylesheet" href="./style.css">
</head>

<body>
<div id="qga">    
    <div class="inputoutput-wrapper">
    <div id="logo">
        QUB GradeMe
    </div>
        <div>
            <input class="display-module" type="text" id="module_1" name="module_1" placeholder="Module 1">
            <input class="display-mark"  type="number" id="mark_1" name="mark_1" placeholder="Mark 1"></br>
            <input class="display-module" type="text" id="module_2" name="module_2" placeholder="Module 2">
            <input class="display-mark"  type="number" id="mark_2" name="mark_2" placeholder="Mark 2"></br>
            <input class="display-module" type="text" id="module_3" name="module_3" placeholder="Module 3">
            <input class="display-mark"  type="number" id="mark_3" name="mark_3" placeholder="Mark 3"></br>
            <input class="display-module" type="text" id="module_4" name="module_4" placeholder="Module 4">
            <input class="display-mark"  type="number" id="mark_4" name="mark_4" placeholder="Mark 4"></br>
            <input class="display-module" type="text" id="module_5" name="module_5" placeholder="Module 5">
            <input class="display-mark"  type="number" id="mark_5" name="mark_5" placeholder="Mark 5">
        </div>
        <div>
            <textarea class="display-output" id="output-text" rows="7" cols="35" readonly=1 placeholder="Results here..." value=""></textarea>
        </div>
    </div>
    
    <div class="button-wrapper">
        <div>
            <button id="maxmin-button" class="qgabutton-active" onclick="if (validateInput()) { getMaxMin(); }">Highest & Lowest Marks</button>
            <button id="averagemark-button" class="qgabutton-active" onclick="if (validateInput()) { getAverageMark(); }">Average Mark</button>
            <button id="classify-button" class="qgabutton-active" onclick="if (validateInput()) { getClassification(); }">Overall Classification</button>
        </div>
        <div>
            <button id="sortmodules-button" class="qgabutton-active" onclick="if (validateInput()) { getSortedModules(); }">Sort Modules</button>
            <button id="totalmarks-button" class="qgabutton-active" onclick="if (validateInput()) { getTotalMark(); }">Total Marks</button>
            <button id="classifymodules-button" class="qgabutton-active" onclick="if (validateInput()) { getModuleClassifications(); }">Module Classification</button>
        </div>
        <div>
            <button id="clear-button" class="qgabutton-clear" onclick="clearText();">Clear Input</button>
        </div>
    </div>

    <div class="saveload-wrapper">
        <p class="info-paragraph">Save or load:</p>
        <p id="save-input-error">Please enter valid input before saving</p>
        <p id="load-input-error">Please enter your unique identifier</p>
        <input class="display-module" type="text" id="unique_id" name="unique_id" placeholder="Unique Identifier">
        <div>
            <button id="save-button" class="qgabutton-active saveload-button" onclick="if (validateSave()) { saveData(); }">Save</button>
            <button id="load-button" class="qgabutton-active saveload-button" onclick="if (validateLoad()) { loadData(); }">Load</button>
        </div>
    </div>
    <a href="./login.php">Metrics & Monitoring</a>
</div>
</body>
</html>
