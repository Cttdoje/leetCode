package com.cttdoje.leetcode.string;

/**
 * @author cttdoje
 * @date 2020/3/4 下午5:59
 */
public class StrStr {
    /**
     * 解题思路：
     * 1. 子串为空 则是0
     * 2. 父串 首位 star 末位 end  初始 star = end = 0
     * 从父串开始遍历 若匹配，则父串 end = index+1，子串index+1
     * 3. 若不匹配，则 父串star = star +1 ,end = star 且 子串index从0开始
     * 4. 结束状态 end = 父串length 或者 子串index = 子串length
     * 5. 最终结果
     * 父串end - 父串star == 子串length 则成功匹配 返回父串star
     * 否则 则匹配失败 返回-1
     */
    public int strStr(String haystack, String needle) {
        if (needle == null || needle.length() == 0) {
            return 0;
        }
        char[] hayCharArray = haystack.toCharArray();
        char[] needleCharArray = needle.toCharArray();
        int star = 0;
        int end = 0;
        int childIndex = 0;
        int parentLength = hayCharArray.length;
        int childLength = needleCharArray.length;
        while (end < parentLength && childIndex < childLength) {
            if (hayCharArray[end] == needleCharArray[childIndex]) {
                end++;
                childIndex++;
            } else {
                star = star + 1;
                end = star;
                childIndex = 0;
            }
        }
        if (end - star == childLength) {
            return star;
        } else {
            return -1;
        }
    }

    public static void main(String[] args) {
        StrStr tool = new StrStr();
        System.err.println(tool.strStr("mississippi", "issip"));
        System.err.println(tool.strStr("hello", "ll"));
    }
}
