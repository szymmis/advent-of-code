import unittest

from main import part_one, part_two


class TestParts(unittest.TestCase):
    def test_part_one(self):
        input = """987654321111111
811111111111119
234234234234278
818181911112111""".splitlines()
        self.assertEqual(part_one(input), 357)

    def test_part_two(self):
        input = """987654321111111
811111111111119
234234234234278
818181911112111""".splitlines()
        self.assertEqual(part_two(input), 3121910778619)


if __name__ == "__main__":
    unittest.main()
