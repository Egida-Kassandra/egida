import sys


class ErrorListener():

    def __init__(self, mode):
        self.mode = mode

    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        print(self.mode)
        print(str(line) + ":" + str(column) + ": sintax ERROR, " + str(msg))
        sys.exit()

    def reportAmbiguity(self, recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs):
        print(self.mode)
        print("Ambiguity ERROR, " + str(configs))
        sys.exit()

    def reportAttemptingFullContext(self, recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs):
        print(self.mode)
        print("Attempting full context ERROR, " + str(configs))
        sys.exit()

    def reportContextSensitivity(self, recognizer, dfa, startIndex, stopIndex, prediction, configs):
        print(self.mode)
        print("Context ERROR, " + str(configs))
        sys.exit()
