class RLEIterator(object):

    def __init__(self, A):
        """
        :type A: List[int]
        """
        self.index = []
        self.data = []
        for i,r in enumerate(A):
            if i % 2 == 0 and r > 0 :
                self.index.append(r)
                self.data.append(A[i+1])
        self.dlen = sum(self.index)
        

    def next(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n > self.dlen:
            self.data = []
            self.index = []
            self.dlen = 0
            return -1
        else:
            for i,no in enumerate(self.index):
                if no < n:
                    self.dlen -= no
                    n -= no
                else:
                    self.dlen -= n
                    self.index[i] -= n
                    break
            
            self.data = self.data[i:]
            self.index = self.index[i:]
            return self.data[0]

                




obj = RLEIterator([3,8,0,9,2,5])
# print(obj.dlen)
for i in [2,1,1,2]:
    print(obj.next(i))


