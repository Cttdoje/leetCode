package com.cttdoje.leetcode.string;

/**
 * @author cttdoje
 * @date 2020/3/20 下午3:06
 */
public class LongestCommonPrefix {
    /**
     *  解题思路:
     *  1. 使用第一个字符串作为基础字符串
     *  2. 遍历比较
     *  3. 从前往后比较，遇到不一致的停止，截取剩下相同的
     *  4. 如果基础字符串为空，则直接退出，无匹配
     */
    public String longestCommonPrefix(String[] strs) {
        if(strs == null || strs.length == 0){
            return "";
        }
        char[] temp = strs[0].toCharArray();
        int endIndex = temp.length;
        for (int i=1;i<strs.length;i++){
            if(endIndex == 0){
                break;
            }
            char[] thisCharArray = strs[i].toCharArray();
            if(thisCharArray.length < endIndex){
                endIndex = thisCharArray.length;
            }
            for (int j=0;j<endIndex;j++){
                if(temp[j] != thisCharArray[j]){
                    endIndex = j;
                    break;
                }
            }
        }
        if(endIndex == 0){
            return "";
        }else{
            String result ="";
            for(int i=0;i<endIndex;i++){
                result += temp[i];
            }
            return result;
        }
    }

    public static void main(String[] args) {
        LongestCommonPrefix longestCommonPrefix = new LongestCommonPrefix();
        String[] strs = {"dog","racecar","car"};
        System.err.println(longestCommonPrefix.longestCommonPrefix(strs));
    }
}
