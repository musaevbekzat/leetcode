void main() {
    IO.println(reverse(123));
}

public int reverse(int x) {
    int result = 0;
    while (x != 0) {
        if (((long) result) * 10 > Integer.MAX_VALUE) {
            return 0;
        }
        if (((long) result) * 10 < Integer.MIN_VALUE) {
            return 0;
        }
        result *= 10;
        result += x % 10;
        x /= 10;
    }

    return result;
}

