x = [0.000000 0.400000 0.800000 1.200000 1.600000 2.000000 2.400000 2.800000 3.200000 3.600000];
y = [0.988568 1.359080 1.868457 2.568748 3.531504 4.855098 6.674770 9.176447 12.615745 17.344076];
x0 = [4 4 4 1 1 1 2 2 2 3 3 3];
y0 = [20 30 24 1 2 4 5 7 6 7 10 12];

plot(x, y);
hold;
plot(x0, y0, 'k *');

print("-r60", "exponential-regression.png");