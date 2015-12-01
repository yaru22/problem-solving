// https://www.hackerrank.com/challenges/harry-potter-and-the-floating-rocks

#include <stdio.h>

int gcd(int a, int b) {
    if (b == 0) {
        return a;
    }
    return gcd(b, a%b);
}

int main() {
    int t, x1, y1, x2, y2, g;
    scanf("%d", &t);
    for (int i = 0; i < t; i++) {
        scanf("%d %d %d %d", &x1, &y1, &x2, &y2);
        g = gcd(y2 - y1, x2 - x1);
        printf("%d\n", g > 0 ? g - 1 : -g - 1);
    }
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    return 0;
}
