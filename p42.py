class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        if len(height) <= 2:
            return 0

        top = []
        max_height = max(height)
        curr = 0
        curr_position = 0
        before = True

        for i in xrange(len(height)):
            val = height[i]
            if val != 0:
                if val == max_height:
                    if before:
                        before = False    
                    curr = 0
                    curr_position = i
                    top.append(i)

                else:
                    if before:
                        if val >= curr :
                            curr = val
                            curr_position = i
                            top.append(i)
                    else:
                        if i >= curr_position and i <= len(height) - 1:
                            print(i)
                            curr = max(height[i:])
                            if curr:
                                curr_position = height.index(curr,i)
                                top.append(curr_position)
                                print(i,curr,curr_position)
                                if curr_position == (len(height) - 1):
                                    break

        print top

        if len(top) <= 1:
            return 0

        cap = 0
        for j in range(len(top) - 1):
            head, tail  = top[j], top[j + 1]
            if head != tail:
                head_val, tail_val = height[head], height[tail]
                margin = min(head_val, tail_val) * (tail - head - 1)
                bottom = sum(height[head+1 : tail])
                cap += margin - bottom
                print(margin,bottom)
        
        print cap
        return cap

        

s = Solution()
s.trap([5,4,1,2])