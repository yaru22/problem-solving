// https://www.hackerrank.com/challenges/and-xor-or

#include <stdio.h>

int calc(int m1, int m2) {
    return ((m1 & m2) ^ (m1 | m2)) & (m1 ^ m2);
}

int main() {
    int n, max = 0, numMin, val;
    scanf("%d", &n);
    int arr[n];
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
        if (i > 0) {
            numMin = arr[i-1];
        }
        for (int j = i - 1; j >= 0; j--) {
            val = calc(arr[i], numMin);
            if (max < val) {
                max = val;
            }
            if (i - j > 5) {
              break;
            }
            if (numMin <= arr[i]) {
                break;
            }
            if (arr[j] < numMin) {
                numMin = arr[j];
            }
            if (arr[i] + numMin <= max) {
                break;
            }
        }
    }
    printf("%d\n", max);
    return 0;
}
