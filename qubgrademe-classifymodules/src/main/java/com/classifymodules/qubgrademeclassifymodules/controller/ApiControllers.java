package com.classifymodules.qubgrademeclassifymodules.controller;

import com.classifymodules.qubgrademeclassifymodules.util.ClassifyModules;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;

@RestController
public class ApiControllers {

    @Autowired
    private ClassifyModules classify;

    @GetMapping(value = "/")
    public ResponseEntity<HashMap<String, Object>> getPage(@RequestParam HashMap<String, String> queryParameters) {
        HashMap<String, Object> parsedParams = new HashMap<>();
        String[] params = { "module_1", "mark_1", "module_2", "mark_2", "module_3", "mark_3", "module_4", "mark_4", "module_5", "mark_5" };

        // Error Handling
        for (int i = 0; i < params.length; i++) {
            boolean containsKey = queryParameters.containsKey(params[i]);
            boolean keyIsModule = params[i].contains("module_");
            String key = params[i].replace('_', ' ');
            key = key.substring(0, 1).toUpperCase() + key.substring(1);

            // Missing parameter error
            if (!containsKey) {
                String msg = key + " value is missing";
                return errorResponse(msg, HttpStatus.BAD_REQUEST);
            }

            // Empty string error
            String trimmedValue = queryParameters.get(params[i]).trim();
            queryParameters.put(params[i], trimmedValue);

            if (trimmedValue.equals("")) {
                String msg = key + " value is missing";
                return errorResponse(msg, HttpStatus.BAD_REQUEST);
            }

            // Key is a module
            if (keyIsModule) {
                parsedParams.put(params[i], queryParameters.get(params[i]));
                continue;
            }

            // Invalid input error
            int value;
            try {
                value = Integer.parseInt(queryParameters.get(params[i]));
            } catch (final NumberFormatException e) {
                String msg = "You must provide a valid integer for " + key;
                return errorResponse(msg, HttpStatus.BAD_REQUEST);
            }

            // Negative integer error
            if (value < 0) {
                String msg = "You must provide a non-negative integer for " + key;
                return errorResponse(msg, HttpStatus.BAD_REQUEST);
            }

            // Integer exceeds 100
            if (value > 100) {
                String msg = "You cannot exceed 100 marks for " + key;
                return errorResponse(msg, HttpStatus.BAD_REQUEST);
            }

            parsedParams.put(params[i], value);
        }

        // Classify the marks
        HashMap<String, Object> classifiedMarks = classify.classifyModules(parsedParams);

        // Ensure that the calculation was successful
        if (classifiedMarks == null) {
            String msg = "Error occurred while performing calculation, ensure that input is valid";
            return errorResponse(msg, HttpStatus.BAD_REQUEST);
        }

        classifiedMarks.put("error", false);
        return ResponseEntity.status(HttpStatus.OK).body(classifiedMarks);
    }

    private ResponseEntity<HashMap<String, Object>> errorResponse(String msg, HttpStatus code) {
        HashMap<String, Object> response = new HashMap<>();

        response.put("error", true);
        response.put("string", msg);

        return ResponseEntity.status(code).body(response);
    }

}
