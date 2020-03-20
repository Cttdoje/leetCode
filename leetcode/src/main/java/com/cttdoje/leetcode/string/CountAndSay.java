package com.cttdoje.leetcode.string;


/**
 * @author cttdoje
 * @date 2020/3/9 下午4:07
 */
public class CountAndSay {
    /**
     * 解题思路
     * 1. 当n = 1时仅输出 1
     * 2. 当n > 1时开始描述
     * 从左往右数,相同的进行count,生成新的数字,
     * 当遇到不同时,重新进行统计, 最终拼接成一个字符串
     * 3. 将生成字符串转成int数组 由此类推...
     */
    public String countAndSay(int n) {
        if (n <= 1) {
            return "1";
        }

        int[] intArr = new int[]{1};
        for (int i = 0; i < n-1; i++) {
            int count = 0;
            int value = -1;
            StringBuilder sb = new StringBuilder();
            for(int j=0;j<intArr.length;j++){
                if(j == 0){
                    value = intArr[j];
                    count++;
                    continue;
                }
                if(value != intArr[j]){
                    sb.append(count).append(value);
                    value = intArr[j];
                    count = 1;
                }else{
                    count++;
                }
            }
            sb.append(count).append(value);
            intArr = stringParseIntArray(sb.toString());
        }
        return intArrayParseString(intArr);
    }

    private int[] stringParseIntArray(String str) {
        int[] arr = new int[str.length()];
        for (int i = 0; i < str.length(); i++) {
            arr[i] = Integer.parseInt(str.substring(i, i + 1));
        }
        return arr;
    }

    private String intArrayParseString(int[] arr){
        StringBuilder result = new StringBuilder();
        for (int value : arr) {
            result.append(value);
        }
        return result.toString();
    }

    public static void main(String[] args) {
        CountAndSay countAndSay = new CountAndSay();
        String str = countAndSay.countAndSay(4);
        System.err.println(str);
    }
}
