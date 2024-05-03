// Import the dependencies
'use strict'; 
const classify = require('./classify.js');
const express = require('express');
const cors = require('cors');
const morgan = require('morgan');

// Initialise the app
const app = express();
app.use(cors());
app.use(morgan('combined'));

// Assign a port number
const port = 3000;

app.get('/', (req, res) => {
    const marks = {
        "Mark 1": req.query.mark_1,
        "Mark 2": req.query.mark_2,
        "Mark 3": req.query.mark_3,
        "Mark 4": req.query.mark_4,
        "Mark 5": req.query.mark_5
    }

    // Error handling
    for (const key in marks) {
        // Missing parameter error
        if (!marks[key]) {
            var msg = key + " value is missing";
            res.statusCode = 400;
            res.type('json') 
            return res.json({ error: true, string: msg, result: null });
        }

        // Parse mark to int
        var mark = parseInt(marks[key]);
        
        // Invalid input error
        if(isNaN(mark) || mark != "" + marks[key]) {
            var msg = "You must provide a valid integer for " + key;
            res.statusCode = 400;
            res.type('json') 
            return res.json({ error: true, string: msg, result: null });
        }

        if (mark < 0) {
            var msg = "You must provide a non-negative integer for " + key;
            res.statusCode = 400;
            res.type('json') 
            return res.json({ error: true, string: msg, result: null });
        }

        if (mark > 100) {
            var msg = "You cannot exceed 100 marks for " + key;
            res.statusCode = 400;
            res.type('json') 
            return res.json({ error: true, string: msg, result: null });
        }

        // Replace mark with int value
        marks[key] = mark;
    }

    // Perform calculation
    var classification = classify.classifyMarks(marks);

    // Ensure calculation was successful
    if (classification == null) {
        var msg = "Error while performing calculation, ensure all input is valid";
        res.statusCode = 400;
        res.type('json') 
        return res.json({ error: true, string: msg, result: null });
    }

    var msg = "Your overall classification is " + classification;

    // Return response
    res.statusCode = 200;
    res.type('json') 
    return res.json({ error: false, string: msg, result: classification });
});

module.exports = app;


