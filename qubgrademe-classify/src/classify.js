module.exports.classifyMarks = function(marks) { 
    // Ensure marks is not null
    if (!marks || Object.keys(marks).length === 0) {
        return null;
    }

    var sum = 0;
    var count = 0;
    var average = 0;

    for (const key in marks) {
        // Ensure value is number
        if (isNaN(marks[key])) {
            return null;
        }

        // Ensure value is within suitable range
        if (marks[key] < 0 || marks[key] > 100) {
            return null;
        }

        sum += marks[key];
        count++;
    }

    // Calculate average
    if (sum == 0) {
        average = 0;
    } else {
        average = sum / count;
    }

    // Assign classification
    if (average >= 70) {
        return "First-Class Honours (1st)";
    } else if (average >= 60) {
        return "Upper Second-Class Honours (2:1)";
    } else if (average >= 50) {
        return "Lower Second-Class Honours (2:2)";
    } else if (average >= 40) {
        return "Third-Class Honours (3rd)";
    } else {
        return "Fail";
    }
}