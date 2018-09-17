class Solution(object):
    def sumSubarrayMins(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        left = [None] * len(A)
        right = [None] * len(A)

        stack = []
        for i in xrange(0, len(A)):
            while stack and A[i] <= A[stack[-1]]:
                stack.pop(0)
            if stack:
                left[i] = stack[-1]
            else:
                left[i] = -1
            stack.append(i)
        
        stack = []
        for i in xrange(len(A)-1, -1 , -1):
            
            while stack and A[i] < A[stack[-1]]:
                stack.pop()
            if stack:
                right[i] = stack[-1]
            else:
                right[i] = len(A)
            stack.append(i)
        
        total = sum([A[i] * (i - left[i]) * (right[i] - i) for i in xrange(len(A))]) % (10 ** 9 + 7)
        print left,right
        return total 

    def sumSubarrayMins2(self, A):
        MOD = 10**9 + 7
        N = len(A)

        # prev has i* - 1 in increasing order of A[i* - 1]
        # where i* is the answer to query j
        stack = []
        prev = [None] * N
        for i in xrange(N):
            while stack and A[i] <= A[stack[-1]]:
                stack.pop()
            prev[i] = stack[-1] if stack else -1
            stack.append(i)

        # next has k* + 1 in increasing order of A[k* + 1]
        # where k* is the answer to query j
        stack = []
        next_ = [None] * N
        for k in xrange(N-1, -1, -1):
            while stack and A[k] < A[stack[-1]]:
                stack.pop()
            next_[k] = stack[-1] if stack else N
            stack.append(k)

        print prev,next_

        # Use prev/next array to count answer
        return [(i - prev[i]) * (next_[i] - i) * A[i]
                   for i in xrange(N)]
            
s = Solution()
print(s.sumSubarrayMins([85,93,93,90]))