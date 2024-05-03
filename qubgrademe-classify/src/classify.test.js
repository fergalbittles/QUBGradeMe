const classify = require('./classify.js');

describe("successfully obtaining a grade classification", () => {
    test("marks should result in a 'First-Class Honours (1st)' classification", () => {
        var marks = { "m1": 30, "m2": 89, "m3": 75, "m4": 92, "m4": 94};
    
        expect(classify.classifyMarks(marks)).toBe("First-Class Honours (1st)");
    });
    
    test("marks should result in a 'Upper Second-Class Honours (2:1)' classification", () => {
        var marks = { "m1": 60, "m2": 69, "m3": 65, "m4": 62, "m4": 64};
    
        expect(classify.classifyMarks(marks)).toBe("Upper Second-Class Honours (2:1)");
    });
    
    test("marks should result in a 'Lower Second-Class Honours (2:2)' classification", () => {
        var marks = { "m1": 50, "m2": 59, "m3": 55, "m4": 52, "m4": 54};
    
        expect(classify.classifyMarks(marks)).toBe("Lower Second-Class Honours (2:2)");
    });
    
    test("marks should result in a 'Third-Class Honours (3rd)' classification", () => {
        var marks = { "m1": 43, "m2": 45, "m3": 40, "m4": 50, "m4": 39};
    
        expect(classify.classifyMarks(marks)).toBe("Third-Class Honours (3rd)");
    });
    
    test("marks should result in a 'Fail' classification", () => {
        var marks = { "m1": 30, "m2": 29, "m3": 15, "m4": 2, "m4": 12};
    
        expect(classify.classifyMarks(marks)).toBe("Fail");
    });
});

describe("passing bad input to the classifyMarks function", () => {
    test("passing an empty object", () => {
        var marks = {};
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });
    
    test("passing an undefined object", () => {
        var marks = undefined;
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });
    
    test("passing a null value", () => {
        var marks = null;
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });
    
    test("passing marks array containing invalid value", () => {
        var marks = { "m1": "asdasda", "m2": 45, "m3": 40, "m4": 50, "m4": 39};
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });
    
    test("passing marks array containing negative number", () => {
        var marks = { "m1": -30, "m2": 29, "m3": 15, "m4": 2, "m4": 12};
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });

    test("passing marks array containing a value exceeding 100", () => {
        var marks = { "m1": 3000, "m2": 29, "m3": 15, "m4": 2, "m4": 12};
    
        expect(classify.classifyMarks(marks)).toBe(null);
    });
});
