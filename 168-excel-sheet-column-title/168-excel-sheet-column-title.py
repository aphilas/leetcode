class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        title = ''
        q = columnNumber

        while q > 0:
            q, r = divmod(q, 26)

            if r == 0:
                q -= 1
                r = 26

            title += chr(64+r)

        return title[::-1]
