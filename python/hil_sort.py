# -*- coding: utf-8 -*-

def hil_sort(l):
    length = len(l)
    cap = length / 2
    while cap > 0:
        groups = length / cap + (1 if length % cap else 0)
        for i in range(cap):
            for j in range(1, groups):
                if (cap*j + i) >= length:
                    break;
                t = j
                tmp = l[cap*t + i]
                while t >= 0:
                    if tmp < l[cap*(t-1) + i]:
                        l[cap*t + i] = l[cap*(t-1) + i]
                        t -= 1
                    break
                l[cap*t + i] = tmp
        cap = cap / 2
    return l

if __name__ == '__main__':
    l = [5, 2, 3, 1, 9, 10, 14]
    l = [14, 10, 9, -4, 1, 3, 2, 5]
    print hil_sort(l)
