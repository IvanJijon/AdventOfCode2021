from unittest import TestCase
from functools import reduce
import day4


class Day4Test(TestCase):
    def test_import_data(self):
        draw_numbers, boards = day4.import_data("day4_input_test")
        self.assertListEqual(draw_numbers, [17, 11, 45, 26, 75])

        board_1 = [[49, 0, 9, 90, 8],
                   [41, 88, 56, 13, 6],
                   [17, 11, 45, 26, 75],
                   [29, 62, 27, 83, 36],
                   [31, 78, 1, 55, 38]]
        self.assertListEqual(boards[0].get_board(), board_1)

        board_2 = [[52, 53, 19, 56, 80],
                   [94, 33, 3, 78, 32],
                   [10, 89, 66, 48, 55],
                   [99, 23, 88, 8, 39],
                   [76, 75, 44, 79, 14]]
        self.assertListEqual(boards[1].get_board(), board_2)

    def test_mark_number(self):
        before = [[49, 0, 9, 90, 8],
                  [41, 88, 56, 13, 6],
                  [17, 11, 45, 26, 75],
                  [29, 62, 27, 83, 36],
                  [31, 78, 1, 55, 38]]

        board = day4.Board(before)
        board.mark_number(88)

        after = [[49, 0, 9, 90, 8],
                 [41, 'x', 56, 13, 6],
                 [17, 11, 45, 26, 75],
                 [29, 62, 27, 83, 36],
                 [31, 78, 1, 55, 38]]

        self.assertListEqual(board.get_board(), after)

    def test_look_for_bingo(self):

        no_bingo = [[49, 0, 9, 90, 8],
                    ['x', 'x', 56, 'x', 'x'],
                    [17, 11, 45, 26, 75],
                    [29, 62, 27, 83, 36],
                    [31, 78, 1, 55, 38]]
        board = day4.Board(no_bingo)
        self.assertFalse(board.bingo())

        horizontal_bingo = [[49, 0, 9, 90, 8],
                            ['x', 'x', 'x', 'x', 'x'],
                            [17, 11, 45, 26, 75],
                            [29, 62, 27, 83, 36],
                            [31, 78, 1, 55, 38]]

        board = day4.Board(horizontal_bingo)
        self.assertTrue(board.bingo())

        vertical_bingo = [[49, 'x', 9, 90, 8],
                          [41, 'x', 56, 13, 6],
                          [17, 'x', 45, 26, 75],
                          [29, 'x', 27, 83, 36],
                          [31, 'x', 1, 55, 38]]

        board = day4.Board(vertical_bingo)
        self.assertTrue(board.bingo())

    def test_sum_of_unmarked(self):

        empty = [[]]
        board = day4.Board(empty)
        self.assertEqual(board.sum_of_unmarked(), 0)

        fully_marked = [['x', 'x', 'x', 'x', 'x']]
        board = day4.Board(fully_marked)
        self.assertEqual(board.sum_of_unmarked(), 0)

        some_marked = [[49, 'x', 'x', 90, 8]]
        board = day4.Board(some_marked)
        self.assertEqual(board.sum_of_unmarked(), 147)

        none_marked = [[49, 0, 9, 90, 8]]
        board = day4.Board(none_marked)
        self.assertEqual(board.sum_of_unmarked(), 156)

        horizontal_bingo = [[49, 0, 9, 90, 8],
                            ['x', 'x', 'x', 'x', 'x'],
                            [17, 11, 45, 26, 75],
                            [29, 62, 27, 83, 36],
                            [31, 78, 1, 55, 38]]

        board = day4.Board(horizontal_bingo)
        self.assertEqual(board.sum_of_unmarked(), 770)

    def test_answer_part_one(self):
        draw_numbers, boards = day4.import_data("day4_input")

        score_first_winner = 0
        for dn in draw_numbers:
            for board in boards:
                board.mark_number(dn)
                if board.bingo():
                    score_first_winner = dn * board.sum_of_unmarked()
                    break
            else:
                continue
            break

        self.assertEqual(score_first_winner, 2745)

    def test_answer_part_two(self):
        draw_numbers, boards = day4.import_data("day4_input")

        winners = []
        last_drawn_winning_number = 1
        for dn in draw_numbers:
            for board in boards:
                if board.available:
                    board.mark_number(dn)
                    if board.bingo():
                        board.available = False
                        winners.append(board)
                        last_drawn_winning_number = dn

        score_last_winner = winners[-1].sum_of_unmarked() * \
            last_drawn_winning_number

        self.assertEqual(score_last_winner, 6594)
