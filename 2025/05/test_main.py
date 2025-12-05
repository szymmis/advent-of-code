import unittest

from main import part_one, part_two


class TestParts(unittest.TestCase):
    def test_part_one(self):
        input = """3-5
10-14
16-20
12-18

1
5
8
11
17
32""".splitlines()
        self.assertEqual(part_one(input), 3)

    def test_part_two(self):
        input = """3-5
10-14
16-20
12-18

1
5
8
11
17
32""".splitlines()

        self.assertEqual(part_two(input), 14)


if __name__ == "__main__":
    unittest.main()
