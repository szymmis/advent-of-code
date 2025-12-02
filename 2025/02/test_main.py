import unittest

from main import part_one, part_two


class TestParts(unittest.TestCase):
    def test_part_one(self):
        input = """11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124""".splitlines()
        self.assertEqual(part_one(input), 1227775554)

    def test_part_two(self):
        input = """11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124""".splitlines()
        self.assertEqual(part_two(input), 4174379265)


if __name__ == "__main__":
    unittest.main()
