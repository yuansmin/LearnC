# -*- coding: utf-8 -*-
# @Date    : 2017-03-27 12:00:01
# @Author  : fancy (fancy@thecover.cn)

import unittest

# 逆波兰表示转换

class Converter(object):
    ops = ['+', '-', '*', '/']

    pri = {
        '+': 1,
        '-': 1,
        '*': 2,
        '/': 2,
        '(': 3,
    }

    def convert(self, s):
        oper = []
        result = ''
        for i in xrange(len(s)):
            if s[i] == ' ':
                continue
            if s[i] in self.ops:
                cur_pri = self.pri[s[i]]
                while (oper):
                    op = oper.pop()
                    if (op == '(') or (self.pri[op] < cur_pri):
                        oper.append(op)
                        break
                    result += ' %s' % op
                oper.append(s[i])
            elif s[i] == '(':
                oper.append(s[i])
            elif s[i] == ')':
                op = oper.pop()
                while op != '(':
                    result += ' %s' % op
                    op = oper.pop()
            else:
                result += ' %s' % s[i]
        while oper:
            op = oper.pop()
            result += ' %s' % op
        return result.strip()


class ConverterTest(unittest.TestCase):

    def testConvert(self):
        s = 'a + b * c + (d * e + f) * g'
        result = 'a b c * + d e * f + g * +'
        con = Converter()
        self.assertEqual(con.convert(s), result)

if __name__ == '__main__':
    unittest.main()
