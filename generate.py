import sys


def main(type_, start, end):
    assert start % 10 == 0
    assert end % 10 == 9
    with open('{}_{}-{}.txt'.format(type_, start, end), 'w') as f:
        f.write('\n'.join('{}:{}-{}'.format(type_, a, a+9) for a in range(start, end, 10)))
        f.write('\n')

if __name__ == '__main__':
    main(sys.argv[1], *[int(i) for i in sys.argv[2:]])

