from functools import reduce

# ToDo list:
# ----------
# - Import data : retrieve both drawn numbers and boards -> check
# - A function that marks the drawn number in all the boards where
#   the drawn number is present -> check
# - A function that verifies if a board has bingo -> check
# - A function that calculates the sum_of_unmarked -> check
#


class Board:
    def __init__(self, board):
        self.board = board
        self.available = True

    def _is_bingo_in(self, board):
        for line in board:
            if self._is_whole_line_marked(line):
                return True

        return False

    def _is_whole_line_marked(self, line):
        return self._only_unmarked_values(line) == []

    def _only_unmarked_values(self, line):
        return list(filter(lambda x: x != 'x', line))

    def get_board(self):
        return self.board

    def mark_number(self, n):
        for line in self.board:
            if n in line:
                line[line.index(n)] = 'x'
                break

    def bingo(self):
        if self._is_bingo_in(self.board):
            return True

        transposed_board = [list(row) for row in zip(*self.board)]

        if self._is_bingo_in(transposed_board):
            return True

        return False

    def sum_of_unmarked(self):
        s = 0
        for line in self.board:
            s += reduce(lambda x, y: x + y,
                        self._only_unmarked_values(line), 0)
        return s


def import_data(file):
    f = open(file, "r")
    draw_numbers = [int(x) for x in f.readline().split(",")]
    boards = []
    while f.readline():  # blank line
        b = []
        for _ in range(0, 5):
            b.append([int(x) for x in f.readline().split()])
        boards.append(Board(b))

    f.close()
    return (draw_numbers, boards)
