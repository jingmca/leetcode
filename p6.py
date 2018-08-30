class Solution(object):
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if (len(s) <= numRows) or numRows < 2:
            return s
        
        last_spaces = 0
        keep_zeros = 0
        strList = [[] for i in xrange(numRows)]
        
        for i in xrange(len(s)):
            last_spaces, after = self.compute_space(last_spaces, numRows, keep_zeros)
            print(s[i], last_spaces, after)
            if last_spaces == 0:
                if after != 0:
                    keep_zeros = 0
                strList[keep_zeros].append(s[i])
                keep_zeros += 1
            if last_spaces != 0:
                if keep_zeros != 0:
                    keep_zeros = 0
                strList[last_spaces].append(s[i])

        # return "".join(map(lambda x:"".join(x), strList))        
        ret = ""
        for i in strList:
            for j in i:
                ret += j
        return ret

    
    def compute_space(self, last_spaces, num_rows, keep_zeros):
        if last_spaces == 0 and keep_zeros == num_rows:
            return(num_rows - 2, 1)
        elif last_spaces == 0 or last_spaces == 1:
            return (0,0)
        else:
            return (last_spaces - 1, num_rows - last_spaces)


s = Solution()
print(s.convert("ABC",2))
