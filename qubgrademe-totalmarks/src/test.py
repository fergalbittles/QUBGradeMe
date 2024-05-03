import unittest
import sum
from app import app

# Test the totalMarks function
class TestSum(unittest.TestCase):
    # Test correct input
    def test_totalMarks1(self):
        marks = {"m1": 20, "m2": 30, "m3": 10, "m4": 5, "m5": 5}

        self.assertEqual(sum.totalMarks(marks), 70)

    
    # Test with null parameter
    def test_totalMarks2(self):
        self.assertEqual(sum.totalMarks(None), False)


    # Test with empty dictionary
    def test_totalMarks3(self):
        marks = {}

        self.assertEqual(sum.totalMarks(marks), False)

    
    # Test with wrong data type
    def test_totalMarks4(self):
        marks = "Hello World!"

        self.assertEqual(sum.totalMarks(marks), False)


    # Test with a value less than 0
    def test_totalMarks5(self):
        marks = {"m1": 20, "m2": -30, "m3": 10, "m4": 5, "m5": 5}

        self.assertEqual(sum.totalMarks(marks), False)


    # Test with a value greater than 100
    def test_totalMarks6(self):
        marks = {"m1": 20, "m2": 9999, "m3": 10, "m4": 5, "m5": 5}

        self.assertEqual(sum.totalMarks(marks), False)


# Test the API endpoint
class TestAPI(unittest.TestCase):
    # Test correct input
    def test_calculateTotal1(self):
        tester = app.test_client(self)
        response = tester.get(path="/", query_string="mark_1=23&mark_2=1&mark_3=76&mark_4=60&mark_5=20")
        statuscode = response.status_code

        self.assertEqual(statuscode, 200)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": false' in response.data)

        self.assertTrue(b'"string": "Total Marks Acquired = 180", "result": 180' in response.data)


    # Test with missing parameter
    def test_calculateTotal2(self):
        tester = app.test_client(self)
        response = tester.get(path="/", query_string="mark_1=23&mark_3=76&mark_4=60&mark_5=20")
        statuscode = response.status_code

        self.assertEqual(statuscode, 400)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": true' in response.data)

        self.assertTrue(b'"string": "Mark 2 value is missing"' in response.data)


    # Test with no parameters
    def test_calculateTotal3(self):
        tester = app.test_client(self)
        response = tester.get(path="/")
        statuscode = response.status_code

        self.assertEqual(statuscode, 400)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": true' in response.data)

        self.assertTrue(b'value is missing' in response.data)


    # Test with invalid input
    def test_calculateTotal4(self):
        tester = app.test_client(self)
        response = tester.get(path="/", query_string="mark_1=23&mark_2=1&mark_3=76&mark_4=sdfsdfsdfs&mark_5=20")
        statuscode = response.status_code

        self.assertEqual(statuscode, 400)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": true' in response.data)

        self.assertTrue(b'"string": "You must provide a valid integer for Mark 4"' in response.data)

    
    # Test with negative value
    def test_calculateTotal5(self):
        tester = app.test_client(self)
        response = tester.get(path="/", query_string="mark_1=23&mark_2=1&mark_3=76&mark_4=-56&mark_5=20")
        statuscode = response.status_code

        self.assertEqual(statuscode, 400)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": true' in response.data)

        self.assertTrue(b'"string": "You must provide a non-negative integer for Mark 4"' in response.data)

    
    # Test with value that exceeds 100
    def test_calculateTotal6(self):
        tester = app.test_client(self)
        response = tester.get(path="/", query_string="mark_1=23&mark_2=1&mark_3=76&mark_4=8900&mark_5=20")
        statuscode = response.status_code

        self.assertEqual(statuscode, 400)

        self.assertEqual(response.content_type, "application/json")

        self.assertTrue(b'"error": true' in response.data)

        self.assertTrue(b'string": "You cannot exceed 100 marks for Mark 4' in response.data)


if __name__ == '__main__':
    unittest.main()
