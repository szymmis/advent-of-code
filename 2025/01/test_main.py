import unittest

from main import part_one, part_two


class TestParts(unittest.TestCase):
    def test_part_one(self):
        input = """L68
L30
R48
L5
R60
L55
L1
L99
R14
L82""".splitlines()
        self.assertEqual(part_one(input), 3)

    def test_part_two(self):
        input = """L68
L30
R48
L5
R60
L55
L1
L99
R14
L82""".splitlines()
        self.assertEqual(part_two(input), 6)


if __name__ == "__main__":
    unittest.main()
