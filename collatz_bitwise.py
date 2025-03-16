import time

def collatz_iter_trad(n):
    while n != 1:
        if n % 2 == 1:
            n = 3*n + 1
        else:
            n = n//2

def collatz_iter_bit(n):
    while n != 1:
        if n & 1:
            n = ((n >> 1) | 1) + n
        else:
            n >>= 1

def collatz_recur_trad(n):
    if n == 1:
        return
    if n % 2 == 1:
        collatz_recur_trad(3*n + 1)
    else:
        collatz_recur_trad(n//2)

def collatz_recur_bit(n):
    if n == 1:
        return
    if n & 1:
        collatz_recur_bit(((n >> 1) | 1) + n)
    else:
        collatz_recur_bit(n >> 1)

def time_func(func, n):
    now = time.time()
    for i in range(1, n+1):
        func(i)
    return time.time() - now

def main():
    n = 10_000_000
    recur_trad_time = time_func(collatz_recur_trad, n)
    recur_bit_time = time_func(collatz_recur_bit, n)
    iter_trad_time = time_func(collatz_iter_trad, n)
    iter_bit_time = time_func(collatz_iter_bit, n)
    print("Traditional (recursive): ", recur_trad_time)
    print("Bitwise (recursive): ", recur_bit_time)
    print("Traditional (iterative): ", iter_trad_time)
    print("Bitwise (iterative): ", iter_bit_time)

if __name__ == "__main__":
    main()