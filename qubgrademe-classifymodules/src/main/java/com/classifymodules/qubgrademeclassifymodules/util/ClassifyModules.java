package com.classifymodules.qubgrademeclassifymodules.util;

import org.springframework.stereotype.Component;

import java.util.HashMap;

@Component
public class ClassifyModules {

    public HashMap<String, Object> classifyModules(HashMap<String, Object> marks) {
        // Ensure marks is not null
        if (marks == null) {
            return null;
        }

        // Ensure marks is not empty
        if (marks.isEmpty()) {
            return null;
        }

        for (String key : marks.keySet()) {
            if (marks.get(key) == null) {
                return null;
            }

            // Ensure value is not whitespace or empty string
            if (marks.get(key) instanceof String) {
                String trimmedValue = ((String) marks.get(key)).trim();
                marks.put(key, trimmedValue);
            }

            if (marks.get(key).equals("")) {
                return null;
            }

            if (key.contains("module_")) {
                continue;
            }

            // Ensure value is valid integer
            int mark;
            try {
                mark = (int)marks.get(key);
            } catch (final ClassCastException e) {
                return null;
            }

            // Ensure mark is within suitable range
            if (mark < 0 || mark > 100) {
                return null;
            }

            if (mark >= 70) {
                marks.put(key, "First-Class Honours (1st)");
            } else if (mark >= 60) {
                marks.put(key, "Upper Second-Class Honours (2:1)");
            } else if (mark >= 50) {
                marks.put(key, "Lower Second-Class Honours (2:2)");
            } else if (mark >= 40) {
                marks.put(key, "Third-Class Honours (3rd)");
            } else {
                marks.put(key, "Fail");
            }
        }

        return marks;
    }

}
