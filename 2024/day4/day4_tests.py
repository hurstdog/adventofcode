import unittest
from day4 import (
    countLTR,
    countRTL,
    countDown,
    countUp,
    countDiagSE,
    countDiagSW,
    countDiagNE,
    countDiagNW,
    countAll,
    countXmas
)

# Test-related constants
_EXAMPLE_TEXT = """
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
"""

_EX_LTR = 3
_EX_RTL = 2
_EX_DOWN = 1
_EX_UP = 2
_EX_DIAG_SE = 1
_EX_DIAG_NE = 4
_EX_DIAG_NW = 4
_EX_DIAG_SW = 1
_EX_TOTAL = 18
_EX_XMAS = 9

def getExampleInput():
    arr = []
    for line in _EXAMPLE_TEXT.splitlines():
        if len(line) == 0:
            continue
        arr.append(list(line))
    return arr

class TestDay4(unittest.TestCase):
    def setUp(self):
        self.example_input = getExampleInput()

    def test_count_ltr(self):
        self.assertEqual(countLTR(self.example_input), _EX_LTR)

    def test_count_rtl(self):
        self.assertEqual(countRTL(self.example_input), _EX_RTL)

    def test_count_down(self):
        self.assertEqual(countDown(self.example_input), _EX_DOWN)

    def test_count_up(self):
        self.assertEqual(countUp(self.example_input), _EX_UP)

    def test_count_diag_se(self):
        self.assertEqual(countDiagSE(self.example_input), _EX_DIAG_SE)

    def test_count_diag_sw(self):
        self.assertEqual(countDiagSW(self.example_input), _EX_DIAG_SW)

    def test_count_diag_ne(self):
        self.assertEqual(countDiagNE(self.example_input), _EX_DIAG_NE)

    def test_count_diag_nw(self):
        self.assertEqual(countDiagNW(self.example_input), _EX_DIAG_NW)

    def test_count_all(self):
        self.assertEqual(countAll(self.example_input), _EX_TOTAL)

    def test_count_xmas(self):
        self.assertEqual(countXmas(self.example_input), _EX_XMAS)

    def test_example_input_structure(self):
        # Test that example input is properly structured
        self.assertTrue(len(self.example_input) > 0)
        self.assertTrue(all(len(row) > 0 for row in self.example_input))
        self.assertTrue(all(len(row) == len(self.example_input[0]) for row in self.example_input))

if __name__ == '__main__':
    unittest.main() 