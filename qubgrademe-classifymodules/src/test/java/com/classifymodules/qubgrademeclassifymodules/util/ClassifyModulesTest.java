package com.classifymodules.qubgrademeclassifymodules.util;

import org.junit.jupiter.api.Test;

import java.util.HashMap;

import static org.junit.jupiter.api.Assertions.*;

class ClassifyModulesTest {

    @Test
    void classifyModules_successfulRequest() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", 41);
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);
        parsedMarks.put("module_1", "Databases");
        parsedMarks.put("module_2", "Programming");
        parsedMarks.put("module_3", "Cyber Security");
        parsedMarks.put("module_4", "Concurrent Programming");
        parsedMarks.put("module_5", "Cloud Computing");

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(10, classification.size());
        assertEquals("Fail", classification.get("mark_1"));
        assertEquals("First-Class Honours (1st)", classification.get("mark_2"));
        assertEquals("Third-Class Honours (3rd)", classification.get("mark_3"));
        assertEquals("Lower Second-Class Honours (2:2)", classification.get("mark_4"));
        assertEquals("Upper Second-Class Honours (2:1)", classification.get("mark_5"));
    }

    @Test
    void classifyModules_nullRequest() {
        ClassifyModules classifyModules = new ClassifyModules();

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(null);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_emptyMap() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_invalidValue() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", "Hello World!");
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_nullMapValue() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", null);
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_negativeParam() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", -41);
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_largeParam() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", 41);
        parsedMarks.put("mark_4", 5298);
        parsedMarks.put("mark_5", 63);

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_emptyStringModule() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", 41);
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);
        parsedMarks.put("module_1", "Databases");
        parsedMarks.put("module_2", "Programming");
        parsedMarks.put("module_3", "");
        parsedMarks.put("module_4", "Concurrent Programming");
        parsedMarks.put("module_5", "Cloud Computing");

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

    @Test
    void classifyModules_whitespaceStringModule() {
        ClassifyModules classifyModules = new ClassifyModules();
        HashMap<String, Object> parsedMarks = new HashMap<>();

        // Add marks
        parsedMarks.put("mark_1", 13);
        parsedMarks.put("mark_2", 89);
        parsedMarks.put("mark_3", 41);
        parsedMarks.put("mark_4", 52);
        parsedMarks.put("mark_5", 63);
        parsedMarks.put("module_1", "Databases");
        parsedMarks.put("module_2", "      ");
        parsedMarks.put("module_3", "Cyber Security");
        parsedMarks.put("module_4", "Concurrent Programming");
        parsedMarks.put("module_5", "Cloud Computing");

        // Classify Modules
        HashMap<String, Object> classification = classifyModules.classifyModules(parsedMarks);

        // Make assertions
        assertEquals(classification, null);
    }

}