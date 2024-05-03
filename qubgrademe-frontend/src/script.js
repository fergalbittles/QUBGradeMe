// This will be triggered if no servers are online
window.addEventListener("unhandledrejection", function(promiseRejectionEvent) { 
    disableButtons(false);
    document.getElementById('output-text').value = "Error: Unable to connect to a server";
    restoreConfig();
});

// Keep track of proxy arrays have been attempted for each request
var usedProxyURLs = [];

function restoreConfig() {
    var tempArray = CONFIG.proxyURLs.concat(usedProxyURLs);
    CONFIG.proxyURLs = tempArray;
    usedProxyURLs = [];
}

var unique_identifier_overwrite = "";

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

function displayResponse(j, service, dbQuery) {
    if (service == "maxmin") {
        let max_module = j.max_module;
        let min_module = j.min_module;
        displayMaxMin(max_module,min_module);
        disableButtons(false);
        return;
    }
    
    if (service == "totalmarks") {
        let total = j.result;
        displayTotalMarks(total);
        disableButtons(false);
        return;
    }

    if (service == "sortmodules") {
        let sorted_modules_returned = j.sorted_modules;
        let sorted_modules = '';
        for (let i = 0; i < sorted_modules_returned.length; i++) {
            sorted_modules += sorted_modules_returned[i]['module'] + ' - ' + sorted_modules_returned[i]['marks'] + '\r\n';
        }
        displaySortedModules(sorted_modules);
        disableButtons(false);
        return;
    }

    if (service == "classify") {
        let classification = j.result;
        displayClassification(classification);
        disableButtons(false);
        return;
    }

    if (service == "averagemark") {
        let average = j.average;
        displayAverageMark(average);
        disableButtons(false);
        return;
    }

    if (service == "classifymodules") {
        let classified_modules = '';
        classified_modules += j.module_1 + ' - ' + j.mark_1 + '\r\n';
        classified_modules += j.module_2 + ' - ' + j.mark_2 + '\r\n';
        classified_modules += j.module_3 + ' - ' + j.mark_3 + '\r\n';
        classified_modules += j.module_4 + ' - ' + j.mark_4 + '\r\n';
        classified_modules += j.module_5 + ' - ' + j.mark_5 + '\r\n';
        displayClassifiedModules(classified_modules);
        disableButtons(false);
        return;
    }

    if (service == "database") {
        if (dbQuery == "save") {
            displayNewId(j.insertedId.$oid);
            disableButtons(false);
            setTimeout(function() {
                alert("Your data has been saved. Load your data using the following identifier: " + j.insertedId.$oid);
            }, 150);
        }

        if (dbQuery == "overwrite") {
            displayOverwriteMessage();
            disableButtons(false);
            setTimeout(function() {
                alert("Your data has been overwritten. You can continue to load your data using your unique identifier: " + unique_identifier_overwrite);
            }, 150);
        }

        if (dbQuery == "load") {
            displayLoadedData(j.document);
            disableButtons(false);
            setTimeout(function() {
                alert("Your data has been loaded successfully");
            }, 150);
        }
    }
}

function httpRequest(route, service, method, dbQuery) {
    if (service != "database") {
        document.getElementById('output-text').value = "Fetching result...";
    }

    if (service == "database") {
        document.getElementById('output-text').value = "Operation in Progress...";
    }

    disableButtons(true);

    // Call recursive fetch
    const options = {
        method: method,
        headers: {
            'Access-Control-Allow-Origin': '*'
        }
    };
    fetch_retry(route, service, options, dbQuery);
}

function fetch_retry(route, service, options, dbQuery) {
    var urlIndex = getRandomInt(CONFIG.proxyURLs.length);
    var url = CONFIG.proxyURLs[urlIndex]+route;

    // Remove the selected url from the array and store it in a temp array
    usedProxyURLs.push(CONFIG.proxyURLs[urlIndex]); 
    CONFIG.proxyURLs.splice(urlIndex, 1);

    return new Promise(function(resolve, reject) {
        fetch(url, options)
            .then(function(result) {
                // Fetch was successful
                resolve(result);
                return result.json();
            })
            .then(function(json) {
                // Carry out operations on response
                var error = json.error;

                if (!error || error == false) {
                    restoreConfig();
                    displayResponse(json, service, dbQuery);
                } else {
                    restoreConfig();
                    disableButtons(false);
                    displayError(json.string);
                }
            })
            .catch(function(error) {
                // Check the base case
                if (CONFIG.proxyURLs.length == 0) {
                    // Fetch was unsuccessful
                    return reject(error);
                }
                // Recursively fetch once again
                resolve(fetch_retry(route, service, options, dbQuery))
            })
    });
}

function displayError(error) {
    document.getElementById('output-text').value = 'Error:\n\n' + error;
}

function displayNewId(id) {
    document.getElementById('output-text').value = "Data Saved!\n\nPlease make note of your unique identifier:\n\n" + id;
}

function displayOverwriteMessage() {
    document.getElementById('output-text').value = "Data Overwritten!\n\nPlease make note of your unique identifier:\n\n" + unique_identifier_overwrite;
}

function displayLoadedData(data) {
    document.getElementById('module_1').value = data.modules[0];
    document.getElementById('module_2').value = data.modules[1];
    document.getElementById('module_3').value = data.modules[2];
    document.getElementById('module_4').value = data.modules[3];
    document.getElementById('module_5').value = data.modules[4];

    document.getElementById('mark_1').value = data.marks[0];
    document.getElementById('mark_2').value = data.marks[1];
    document.getElementById('mark_3').value = data.marks[2];
    document.getElementById('mark_4').value = data.marks[3];
    document.getElementById('mark_5').value = data.marks[4];

    document.getElementById('output-text').value = "Your data has been loaded successfully!";
}

function displayMaxMin(max_module,min_module) {
    document.getElementById('output-text').value = 'Highest scoring module = ' + max_module
    + '\nLowest scoring module = ' + min_module;

}

function displaySortedModules(sorted_modules) {
    document.getElementById('output-text').value = sorted_modules;

}

function displayTotalMarks(total) {
    document.getElementById('output-text').value = 'Total marks acquired = ' + total;

}

function displayAverageMark(average) {
    document.getElementById('output-text').value = 'Average mark = ' + average;

}

function displayClassification(classification) {
    document.getElementById('output-text').value = 'Your overall classification is: ' + classification;
}

function displayClassifiedModules(classified_modules) {
    document.getElementById('output-text').value = 'Your module classifications are:\n\n' + classified_modules;
}

function clearText() {
    let ids = [
        "module_1",
        "module_2",
        "module_3",
        "module_4",
        "module_5",
        "mark_1",
        "mark_2",
        "mark_3",
        "mark_4",
        "mark_5"
    ];

    for (var i in ids) {
        document.getElementById(ids[i]).value = '';
        updateErrorStyle(ids[i], false);
    }

    document.getElementById('output-text').value = '';
    document.getElementById('unique_id').value = '';
    document.getElementById("save-input-error").style.display = null
    document.getElementById("load-input-error").style.display = null
    updateErrorStyle("unique_id", false);
}

function getMaxMin() {
    let module_1 = document.getElementById('module_1').value;
    let module_2 = document.getElementById('module_2').value;
    let module_3 = document.getElementById('module_3').value;
    let module_4 = document.getElementById('module_4').value;
    let module_5 = document.getElementById('module_5').value;

    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=maxmin&module_1=" + module_1 + "&mark_1=" + mark_1 + "&module_2=" + module_2 + "&mark_2=" + mark_2
    + "&module_3=" + module_3 + "&mark_3=" + mark_3 + "&module_4=" + module_4 + "&mark_4=" + mark_4
    + "&module_5=" + module_5 + "&mark_5=" + mark_5;

    httpRequest(route, "maxmin", "GET");

    return;
}

function getSortedModules() {
    let module_1 = document.getElementById('module_1').value;
    let module_2 = document.getElementById('module_2').value;
    let module_3 = document.getElementById('module_3').value;
    let module_4 = document.getElementById('module_4').value;
    let module_5 = document.getElementById('module_5').value;

    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=sortmodules&module_1=" + module_1 + "&mark_1=" + mark_1 + "&module_2=" + module_2 + "&mark_2=" + mark_2
    + "&module_3=" + module_3 + "&mark_3=" + mark_3 + "&module_4=" + module_4 + "&mark_4=" + mark_4
    + "&module_5=" + module_5 + "&mark_5=" + mark_5;

    httpRequest(route, "sortmodules", "GET");

    return;
}

function getTotalMark() {
    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=totalmarks&mark_1=" + mark_1 + "&mark_2=" + mark_2 + "&mark_3=" + mark_3 + "&mark_4=" + mark_4 + "&mark_5=" + mark_5;

    httpRequest(route, "totalmarks", "GET");

    return;
}

function getAverageMark() {
    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=averagemark&mark_1=" + mark_1 + "&mark_2=" + mark_2 + "&mark_3=" + mark_3 + "&mark_4=" + mark_4 + "&mark_5=" + mark_5;

    httpRequest(route, "averagemark", "GET");

    return;
}

function getClassification() {
    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=classify&mark_1=" + mark_1 + "&mark_2=" + mark_2 + "&mark_3=" + mark_3 + "&mark_4=" + mark_4 + "&mark_5=" + mark_5;

    httpRequest(route, "classify", "GET");

    return;
}

function getModuleClassifications() {
    let module_1 = document.getElementById('module_1').value;
    let module_2 = document.getElementById('module_2').value;
    let module_3 = document.getElementById('module_3').value;
    let module_4 = document.getElementById('module_4').value;
    let module_5 = document.getElementById('module_5').value;

    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    const route = "?route=classifymodules&module_1=" + module_1 + "&mark_1=" + mark_1 + "&module_2=" + module_2 + "&mark_2=" + mark_2
    + "&module_3=" + module_3 + "&mark_3=" + mark_3 + "&module_4=" + module_4 + "&mark_4=" + mark_4
    + "&module_5=" + module_5 + "&mark_5=" + mark_5;

    httpRequest(route, "classifymodules", "GET");

    return;
}

function validateInput() {
    trimInputStrings();

    var valid = true;

    let marks = {
        "mark_1": document.getElementById('mark_1').value,
        "mark_2": document.getElementById('mark_2').value,
        "mark_3": document.getElementById('mark_3').value,
        "mark_4": document.getElementById('mark_4').value,
        "mark_5": document.getElementById('mark_5').value,
    };

    let modules = {
        "module_1": document.getElementById('module_1').value,
        "module_2": document.getElementById('module_2').value,
        "module_3": document.getElementById('module_3').value,
        "module_4": document.getElementById('module_4').value,
        "module_5": document.getElementById('module_5').value,
    };

    for (var key in marks) {
        if (!marks[key]) {
            updateErrorStyle(key, true);
            valid = false;
            continue;
        }

        var mark = parseInt(marks[key]);

        if (isNaN(mark) || mark != "" + marks[key] || mark < 0 || mark > 100) {
            updateErrorStyle(key, true);
            valid = false;
            continue;
        }

        updateErrorStyle(key, false);
    }

    for (var key in modules) {
        if (!modules[key]) {
            updateErrorStyle(key, true);
            valid = false;
            continue;
        } 

        var parsed = parseInt(modules[key]);

        if (!isNaN(parsed)) {
            updateErrorStyle(key, true);
            valid = false;
            continue;
        }

        updateErrorStyle(key, false);
    }

    const msg = "Please enter valid input into the highlighted text fields.\n\n- 'Mark' values must be integers from 0-100\n- 'Module' values must not be numbers";
    document.getElementById('output-text').value = !valid ? msg : '';
  
    return valid;
}

function validateSave() {
    document.getElementById("load-input-error").style.display = null;
    updateErrorStyle("unique_id", false);

    var validInputFields =  validateInput();

    document.getElementById("save-input-error").style.display = validInputFields ? null : "block";

    return validInputFields;
}

function validateLoad() {
    updateErrorStyle("module_1", false);
    updateErrorStyle("module_2", false);
    updateErrorStyle("module_3", false);
    updateErrorStyle("module_4", false);
    updateErrorStyle("module_5", false);
    updateErrorStyle("mark_1", false);
    updateErrorStyle("mark_2", false);
    updateErrorStyle("mark_3", false);
    updateErrorStyle("mark_4", false);
    updateErrorStyle("mark_5", false);
    document.getElementById("output-text").value = "";
    document.getElementById("save-input-error").style.display = null;

    document.getElementById("unique_id").value = document.getElementById("unique_id").value.trim();
    let id  = document.getElementById("unique_id").value;

    document.getElementById("load-input-error").style.display = id == "" ? "block" : null;

    if (id == "") {
        updateErrorStyle("unique_id", true);
        return false;
    }
    updateErrorStyle("unique_id", false);

    var input = "";
    input += document.getElementById('module_1').value;
    input += document.getElementById('module_2').value;
    input += document.getElementById('module_3').value;
    input += document.getElementById('module_4').value;
    input += document.getElementById('module_5').value;
    input += document.getElementById('mark_1').value;
    input += document.getElementById('mark_2').value;
    input += document.getElementById('mark_3').value;
    input += document.getElementById('mark_4').value;
    input += document.getElementById('mark_5').value;

    if (input.length > 0) {
        if(!confirm("Are you sure you wish to load? Your unsaved input will be lost")) {
            return false;
        }
    }

    return true;
}

function saveData() {
    document.getElementById("unique_id").value = document.getElementById("unique_id").value.trim();
    let id  = document.getElementById("unique_id").value;

    let module_1 = document.getElementById('module_1').value;
    let module_2 = document.getElementById('module_2').value;
    let module_3 = document.getElementById('module_3').value;
    let module_4 = document.getElementById('module_4').value;
    let module_5 = document.getElementById('module_5').value;

    let mark_1 = document.getElementById('mark_1').value;
    let mark_2 = document.getElementById('mark_2').value;
    let mark_3 = document.getElementById('mark_3').value;
    let mark_4 = document.getElementById('mark_4').value;
    let mark_5 = document.getElementById('mark_5').value;

    var route = "?route=database&module_1=" + module_1 + "&mark_1=" + mark_1 + "&module_2=" + module_2 + "&mark_2=" + mark_2
    + "&module_3=" + module_3 + "&mark_3=" + mark_3 + "&module_4=" + module_4 + "&mark_4=" + mark_4
    + "&module_5=" + module_5 + "&mark_5=" + mark_5;

    // User is making a new save
    if (id == "") {
        httpRequest(route, "database", "PUT", "save");
    }

    // User is overwriting saved data
    if (id != "") {
        if (confirm("You are about to overwrite the saved data for identifier " + id + ". Are you sure you wish to proceed?")) {
            route += "&id=" + id;
            httpRequest(route, "database", "PUT", "overwrite");
        }
        unique_identifier_overwrite = id;
    }

    return;
}

function loadData() {
    document.getElementById("unique_id").value = document.getElementById("unique_id").value.trim();
    let id  = document.getElementById("unique_id").value;

    const route = "?route=database&id=" + id;

    httpRequest(route, "database", "GET", "load");
}

function trimInputStrings() {
    let ids = [
        "module_1",
        "module_2",
        "module_3",
        "module_4",
        "module_5",
        "mark_1",
        "mark_2",
        "mark_3",
        "mark_4",
        "mark_5"
    ];

    for (var i in ids) {
        document.getElementById(ids[i]).value = document.getElementById(ids[i]).value.trim();
        updateErrorStyle(ids[i], false);
    }
}

function updateErrorStyle(id, flag) {
    document.getElementById(id).style.borderTop = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderRight = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderBottom = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderLeft = flag ? "2px solid red" : null;
    document.getElementById(id).style.backgroundColor = flag ? "#fccaca" : null;
}

function disableButtons(flag) {
    document.getElementById("maxmin-button").disabled = flag;
    document.getElementById("sortmodules-button").disabled = flag;
    document.getElementById("totalmarks-button").disabled = flag;
    document.getElementById("classify-button").disabled = flag;
    document.getElementById("averagemark-button").disabled = flag;
    document.getElementById("classifymodules-button").disabled = flag;
    document.getElementById("clear-button").disabled = flag;
    document.getElementById("save-button").disabled = flag;
    document.getElementById("load-button").disabled = flag;

    document.getElementById("maxmin-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("maxmin-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("sortmodules-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("sortmodules-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("totalmarks-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("totalmarks-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("classify-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("classify-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("averagemark-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("averagemark-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("classifymodules-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("classifymodules-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("clear-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("clear-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("save-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("save-button").style.cursor = flag ? "not-allowed" : null;
    document.getElementById("load-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("load-button").style.cursor = flag ? "not-allowed" : null;
    
}
