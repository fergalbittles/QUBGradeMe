// This will be triggered if no servers are online
window.addEventListener("unhandledrejection", function(promiseRejectionEvent) { 
    if (LOGIN_ATTEMPT == true) {
        disableLoginButton(false);
        displayLoginError("Error: Unable to connect to a server", true);
    } else {
        disableTestButtons(false);
        document.getElementById("tests_are_running").style.display = null;
        document.getElementById("logs_are_fetching").style.display = null;
        document.getElementById("server_error").style.display = "block";
    }

    restoreConfig();
});

// Set to true when a login is attempted, then the listener above can catch any login server errors
var LOGIN_ATTEMPT = false;

// Keep track of proxy arrays have been attempted for each request
var usedProxyURLs = [];

function restoreConfig() {
    var tempArray = CONFIG.proxyURLs.concat(usedProxyURLs);
    CONFIG.proxyURLs = tempArray;
    usedProxyURLs = []
}

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

function injectTestResults(results) {
    const chartLabels = [
        "Proxy Route", 
        "Max/Min", 
        "Sort Modules", 
        "Total Marks", 
        "Average Mark", 
        "Overall Classify", 
        "Module Classify", 
        "Database Save", 
        "Database Load"
    ];
    const chartValues = [];

    var totalTests = results.length;
    var testsPassed = 0;
    var testsTimedOut = 0;

    for (let i = 0; i < results.length; i++) {
        var id = "test-" + (i+1) + "-result";

        if (results[i] == null || results[i] == undefined) {
            document.getElementById(id+"-status").innerHTML = "TIMEOUT DURING EXECUTION";
            chartValues.push(0);
            testsTimedOut++;
            continue;
        }

        chartValues.push(parseFloat(results[i].time_taken));

        if (results[i].result == "PASSED") {
            testsPassed++;
        }

        // Inject the status of the test
        document.getElementById(id+"-status").innerHTML = results[i].result;
        document.getElementById(id+"-status").style.backgroundColor = results[i].result == "PASSED" ? "#52ff63" : "#ff5a52";

        // Inject the date of the test 
        var dateArr = results[i].date.split("T");
        dateArr = dateArr[0].split("-");
        document.getElementById(id+"-date").innerHTML = dateArr[2] + "/" + dateArr[1] + "/" + dateArr[0];

        // Inject the result of the test
        document.getElementById(id+"-result").innerHTML = results[i].result;
        document.getElementById(id+"-result").style.color = results[i].result == "PASSED" ? null : "red";

        // Inject the time of the test
        document.getElementById(id+"-time").innerHTML = Number((parseFloat(results[i].time_taken)).toFixed(5));

        // Inject the expected status of the test
        document.getElementById(id+"-ex-status").innerHTML = results[i].expected_status;

        // Inject the actual status of the test
        document.getElementById(id+"-ac-status").innerHTML = results[i].actual_status;
        document.getElementById(id+"-ac-status").style.color = results[i].actual_status == results[i].expected_status ? null : "red";

        // Inject the expected response of the test
        document.getElementById(id+"-ex-res").innerHTML = results[i].expected_response;

        // Inject the actual response of the test
        document.getElementById(id+"-ac-res").innerHTML = results[i].actual_response;
        document.getElementById(id+"-ac-res").style.color = results[i].actual_response == results[i].expected_response ? null : "red";

        // Display an overall message
        document.getElementById("tests_passed_message").style.display = "block";
        document.getElementById("tests_passed_message").innerHTML = "<b>Tests Passed = " + testsPassed + "/" + totalTests + "</b>";
        
        if (testsPassed == totalTests) {
            document.getElementById("tests_passed_message").style.color = "green";
        }

        if (testsPassed != totalTests) {
            var testsFailed = (totalTests-testsTimedOut) - testsPassed;
            document.getElementById("tests_passed_message").innerHTML = document.getElementById("tests_passed_message").innerHTML + "<br><b style=\"color: red;\">Tests Failed = " + testsFailed + "/" + totalTests + "</b>";
        }

        if (testsTimedOut > 0) {
            document.getElementById("tests_passed_message").innerHTML = document.getElementById("tests_passed_message").innerHTML + "<br><b style=\"color: #f08102;\">Timeout Errors = " + testsTimedOut + "/" + totalTests + "</b>";
        }
    }

    showChart(chartLabels, chartValues, "");
}

function showChart(chartLabels, chartValues, type) {
    var chart = document.getElementById("myChart");
    chart.remove();

    var newCanvas = document.createElement("canvas");
    newCanvas.setAttribute("id", "myChart");
    document.getElementById("chart-wrapper-inner").appendChild(newCanvas); 

    var labelString = 'Execution Time (s)';

    if (type == "average") {
        labelString = 'Average ' + labelString;
    }

    chart = document.getElementById("myChart");

    new Chart(chart, {
        type: 'bar',
        data: {
            labels: chartLabels,
            datasets: [{
                label: labelString,
                data: chartValues,
                borderWidth: 1,
                backgroundColor: ["#fd7f6f", "#7eb0d5", "#b2e061", "#bd7ebe", "#ffb55a", "#ffee65", "#beb9db", "#fdcce5", "#8bd3c7"]
            }]
        },
        options: {
            plugins: {
                legend: {
                    display: false,
                }
            },
            scales: {
                y: {
                    type: 'logarithmic',
                    title: {
                        display: true,
                        text: labelString,
                        padding: 15,
                        font: {
                            weight: "bold",
                            size: 14
                        }
                    }
                },
                x: {
                    title: {
                        display: true,
                        text: "Test Name",
                        padding: 15,
                        font: {
                            weight: "bold",
                            size: 14
                        }
                    }
                }
            }
        }
    });

    if (type == "average") {
        document.getElementById("average_execution_time").style.display = "block";
    } else {
        document.getElementById("execution_time").style.display = "block";
    }

    document.getElementById("chart-wrapper").style.display = "block";
}

function  displayTestResponse(j, service) {
    if (service == "fetch-logs") {
        displayLogs(j.logs);
        return;
    }

    // Innect the test results into the UI
    injectTestResults(j.results);

    // Show the results for each test
    for (let i = 0; i < j.results.length; i++) {
        var id = "test-" + (i+1) + "-result";
        document.getElementById(id).style.display = "block";
    }

    // Remove the "tests are running" notification
    document.getElementById("tests_are_running").style.display = null;
    disableTestButtons(false);
}

function displayLogs(logs) {
    const chartLabels = [
        "Proxy Route", 
        "Max/Min", 
        "Sort Modules", 
        "Total Marks", 
        "Average Mark", 
        "Overall Classify", 
        "Module Classify", 
        "Database Save", 
        "Database Load"
    ];

    const totalExecutions = [0, 0, 0, 0, 0, 0, 0, 0, 0];
    const totalTime = [0, 0, 0, 0, 0, 0, 0, 0, 0];

    // Store dates and values for all tests
    const test1Dates = [];
    const test1Values = [];
    const test2Dates = [];
    const test2Values = [];
    const test3Dates = [];
    const test3Values = [];
    const test4Dates = [];
    const test4Values = [];
    const test5Dates = [];
    const test5Values = [];
    const test6Dates = [];
    const test6Values = [];
    const test7Dates = [];
    const test7Values = [];
    const test8Dates = [];
    const test8Values = [];
    const test9Dates = [];
    const test9Values = [];

    for (let i = 0; i < logs.length; i++) {
        for (let j = 0; j < logs[i].results.length; j++) {
            if (logs[i].results[j] == null || logs[i].results[j] == undefined) {
                continue;
            }

            totalExecutions[j] += 1;
            totalTime[j] += parseFloat(logs[i].results[j].time_taken);

            var dateTime = logs[i].date.split("T");
            var date = dateTime[0].split("-");
            date = date[2] + "/" + date[1] + " " + dateTime[1].substring(0, 5);

            switch(j) {
                case 0: 
                    test1Dates.push(date);
                    test1Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 1: 
                    test2Dates.push(date);
                    test2Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 2: 
                    test3Dates.push(date);
                    test3Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 3: 
                    test4Dates.push(date);
                    test4Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 4: 
                    test5Dates.push(date);
                    test5Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 5: 
                    test6Dates.push(date);
                    test6Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 6: 
                    test7Dates.push(date);
                    test7Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 7: 
                    test8Dates.push(date);
                    test8Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
                case 8: 
                    test9Dates.push(date);
                    test9Values.push(parseFloat(logs[i].results[j].time_taken));
                    break;
            }
        }
    }

    // Calculate average execution time for each test
    for (let i = 0; i < totalTime.length; i++) {
        totalTime[i] = totalTime[i] / totalExecutions[i];
    }

    // Show average execution time in a chart
    showChart(chartLabels, totalTime, "average");

    // Show all of the line graphs
    updateGraph(test1Dates, test1Values, 1);
    updateGraph(test2Dates, test2Values, 2);
    updateGraph(test2Dates, test3Values, 3);
    updateGraph(test2Dates, test4Values, 4);
    updateGraph(test2Dates, test5Values, 5);
    updateGraph(test2Dates, test6Values, 6);
    updateGraph(test2Dates, test7Values, 7);
    updateGraph(test2Dates, test8Values, 8);
    updateGraph(test2Dates, test9Values, 9);

    for (let i = 0; i < 9; i++) {
        document.getElementById("graph-wrapper-" + (i+1)).style.display = "block";
    }

    // Remove the "fetching logs" notification
    document.getElementById("logs_are_fetching").style.display = null;
    disableTestButtons(false);
}

function updateGraph(graphLabels, graphValues, index) {
    const graphId = "myGraph-" + index;
    const graphWrapperId = "graph-wrapper-inner-" + index;

    var graph = document.getElementById(graphId);
    graph.remove();

    var newCanvas = document.createElement("canvas");
    newCanvas.setAttribute("id", graphId);
    document.getElementById(graphWrapperId).appendChild(newCanvas); 

    graph = document.getElementById(graphId);

    new Chart(graph, {
        type: "line",
        data: {
            labels: graphLabels,
            datasets: [{
                label: 'Execution Time (s)',
                data: graphValues,
                borderWidth: 1,
            }]
        },
        options: {
            plugins: {
                legend: {
                    display: false,
                }
            },
            scales: {
                y: {
                    type: 'logarithmic',
                    title: {
                        display: true,
                        text: 'Execution Time (s)',
                        padding: 15,
                        font: {
                            weight: "bold",
                            size: 14
                        }
                    }
                },
                x: {
                    title: {
                        display: true,
                        text: "Date",
                        padding: 15,
                        font: {
                            weight: "bold",
                            size: 14
                        }
                    }
                }
            }
        }
    });
}

function displayTestError(error, flag) {
    document.getElementById("tests_are_running").style.display = null;
    document.getElementById("logs_are_fetching").style.display = null;
    document.getElementById("generic_error").innerHTML = error;
    document.getElementById("generic_error").style.display = flag ? "block" : null;
}

function httpTestRequest(route, service, method) {
    disableTestButtons(true);

    // Call recursive fetch
    const options = {
        method: method,
        headers: {
            'Access-Control-Allow-Origin': '*'
        }
    };
    fetch_test_retry(route, service, options);
}

function fetch_test_retry(route, service, options) {
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
                    if (service == "login") { 
                        restoreConfig();
                        loginAccepted(json.auth);
                    } else {
                        restoreConfig();
                        displayTestResponse(json, service);
                    }
                } else {
                    if (service == "login") {
                        restoreConfig();
                        disableLoginButton(false);
                        displayLoginError(json.string, true);
                    } else {
                        restoreConfig();
                        disableTestButtons(false);
                        displayTestError(json.string, true);
                    }
                }
            })
            .catch(function(error) {
                // Check the base case
                if (CONFIG.proxyURLs.length == 0) {
                    // Fetch was unsuccessful
                    return reject(error);
                }
                // Recursively fetch once again
                resolve(fetch_test_retry(route, service, options))
            })
    });
}

function disableTestButtons(flag) {
    document.getElementById("runtests-button").disabled = flag;
    document.getElementById("runtests-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("runtests-button").style.cursor = flag ? "not-allowed" : null;    

    document.getElementById("analysedata-button").disabled = flag;
    document.getElementById("analysedata-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("analysedata-button").style.cursor = flag ? "not-allowed" : null;    
}

function runTests() {
    displayTestError("", false);
    document.getElementById("server_error").style.display = null;
    document.getElementById("tests_passed_message").style.display = null;
    document.getElementById("tests_are_running").style.display = "block";
    document.getElementById("logs_are_fetching").style.display = null;
    document.getElementById("chart-wrapper").style.display = null;
    document.getElementById("average_execution_time").style.display = null;
    document.getElementById("execution_time").style.display = null;

    for (let i = 0; i < 9; i++) {
        document.getElementById("graph-wrapper-" + (i+1)).style.display = null;
    }

    // Hide the results for each test
    for (let i = 0; i < 9; i++) {
        var id = "test-" + (i+1) + "-result";
        document.getElementById(id).style.display = null;
    }

    httpTestRequest("?route=monitor", "run-tests", "GET");
}

function analyseTestData() {
    displayTestError("", false);
    document.getElementById("server_error").style.display = null;
    document.getElementById("tests_passed_message").style.display = null;
    document.getElementById("tests_are_running").style.display = null;
    document.getElementById("logs_are_fetching").style.display = "block";
    document.getElementById("chart-wrapper").style.display = null;
    document.getElementById("average_execution_time").style.display = null;
    document.getElementById("execution_time").style.display = null;

    for (let i = 0; i < 9; i++) {
        document.getElementById("graph-wrapper-" + (i+1)).style.display = null;
    }

    // Hide the results for each test
    for (let i = 0; i < 9; i++) {
        var id = "test-" + (i+1) + "-result";
        document.getElementById(id).style.display = null;
    }

    httpTestRequest("?route=monitorlogs", "fetch-logs", "GET");
}

function validateLogin() {
    document.getElementById("username").value = document.getElementById("username").value.trim();

    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value

    if (username == "" || password == "") {
        if (username == "") {
            updateLoginErrorStyle("username", true);
        } else {
            updateLoginErrorStyle("username", false);
        }

        if (password == "") {
            updateLoginErrorStyle("password", true);
        } else {
            updateLoginErrorStyle("password", false);
        }

        displayLoginError("Enter a username and password", true);
        return false;
    } else {
        updateLoginErrorStyle("username", false);
        updateLoginErrorStyle("password", false);
        displayLoginError("", false);
    }

    return true;
}

function signIn() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value

    disableLoginButton(true);

    // Call recursive fetch
    const route = "admin/login?username="+username+"&password="+password;
    const service = "login";
    const options = {
        method: 'GET',
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Authorization': 'Basic YWRtaW4uYm9zczpzdXBlcnNlY3JldA=='
        }
    };
    LOGIN_ATTEMPT = true;
    fetch_test_retry(route, service, options);
}

function disableLoginButton(flag) {
    document.getElementById("login-button").disabled = flag;
    document.getElementById("login-button").style.backgroundColor = flag ? "#595959" : null;
    document.getElementById("login-button").style.cursor = flag ? "not-allowed" : null;   
}

function displayLoginError(error, flag) {
    document.getElementById("login_error").innerHTML = error;
    document.getElementById("login_error").style.display = flag ? "block" : null;
}

function loginAccepted(auth) {
    LOGIN_ATTEMPT = false;

    disableLoginButton(false);
    document.getElementById("username").value = "";
    document.getElementById("password").value = "";

    if (auth == true) {
        document.location.href = "/metrics.php";
    }
}

function updateLoginErrorStyle(id, flag) {
    document.getElementById(id).style.borderTop = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderRight = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderBottom = flag ? "2px solid red" : null;
    document.getElementById(id).style.borderLeft = flag ? "2px solid red" : null;
    document.getElementById(id).style.backgroundColor = flag ? "#fccaca" : null;
}
