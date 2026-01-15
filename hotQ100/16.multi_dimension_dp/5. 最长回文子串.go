package multidimensiondp

/**
 * dp[i][j] 表示下标[i,j]是否是回文串
 * 当s[i] == s[j] 时，若[i+1, j-1]也是一个回文串，则[i,j]也是回文串
 *                   若 j-1 < i+1 即 j-i <= 1 时，[i,j]一定是回文串
 * 当s[i] != s[j]时，[i, j]不可能是回文串
 */
func longestPalindrome(s string) string {

}
