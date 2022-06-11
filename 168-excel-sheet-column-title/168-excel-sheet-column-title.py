class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        segments = []
        q = columnNumber

        while q > 0:
            q, r = divmod(q, 26)

            if r == 0:
                q -= 1
                r = 26

            segments.insert(0, chr(64+r))

        return ''.join(segments)
