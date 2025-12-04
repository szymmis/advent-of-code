import unittest

from main import part_one, part_two


class TestParts(unittest.TestCase):
    def test_part_one(self):
        input = """..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.""".splitlines()
        self.assertEqual(part_one(input), 13)

    def test_part_two(self):
        input = """..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.""".splitlines()
        self.assertEqual(part_two(input), 43)


if __name__ == "__main__":
    unittest.main()
