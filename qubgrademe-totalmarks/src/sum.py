def totalMarks(marks):
    # Ensure marks is a dictionary
    if not isinstance(marks, dict):
        return False

    # Ensure marks is not empty
    if len(marks) == 0:
        return False

    sum = 0

    for key in marks:
        # Ensure value is a valid int
        try:
            marks[key] = int(marks[key])
        except ValueError:
            return False

        # Ensure value is within suitable range
        if marks[key] < 0 or marks[key] > 100:
            return False

        sum += marks[key]

    return sum