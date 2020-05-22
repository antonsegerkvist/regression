x = [57.000000 57.340000 57.680000 58.020000 58.360001 58.700001 59.040001 59.380001 59.720001 60.060001 60.400002 60.740002 61.080002 61.420002 61.760002 62.100002 62.440002 62.780003 63.120003 63.460003 63.800003 64.139999 64.479996 64.819992 65.159988 65.499985 65.839981 66.179977 66.519974 66.859970 67.199966 67.539963 67.879959 68.219955 68.559952 68.899948 69.239944 69.579941 69.919937 70.259933 70.599930 70.939926 71.279922 71.619919 71.959915 72.299911 72.639908 72.979904 73.319901 73.659897 73.999893 74.339890 74.679886 75.019882 75.359879 75.699875 76.039871 76.379868 76.719864 77.059860 77.399857 77.739853 78.079849 78.419846 78.759842 79.099838 79.439835 79.779831 80.119827 80.459824 80.799820 81.139816 81.479813 81.819809 82.159805 82.499802 82.839798 83.179794 83.519791 83.859787 84.199783 84.539780 84.879776 85.219772 85.559769 85.899765 86.239761 86.579758 86.919754 87.259750 87.599747 87.939743 88.279739 88.619736 88.959732 89.299728 89.639725 89.979721 90.319717 90.659714 90.999710];
y = [164.271576 164.529724 164.786743 165.042648 165.297455 165.551178 165.803818 166.055389 166.305908 166.555374 166.803802 167.051193 167.297577 167.542953 167.787338 168.030716 168.273148 168.514587 168.755066 168.994598 169.233185 169.470840 169.707550 169.943359 170.178268 170.412262 170.645370 170.877594 171.108948 171.339417 171.569031 171.797806 172.025726 172.252792 172.479034 172.704468 172.929077 173.152863 173.375854 173.598068 173.819458 174.040085 174.259933 174.479004 174.697327 174.914886 175.131683 175.347733 175.563049 175.777634 175.991486 176.204636 176.417053 176.628754 176.839752 177.050049 177.259674 177.468582 177.676819 177.884384 178.091263 178.297485 178.503036 178.707932 178.912170 179.115768 179.318726 179.521042 179.722733 179.923782 180.124207 180.324020 180.523209 180.721802 180.919769 181.117157 181.313919 181.510101 181.705704 181.900696 182.095139 182.288986 182.482269 182.674973 182.867111 183.058685 183.249725 183.440201 183.630127 183.819504 184.008331 184.196625 184.384384 184.571609 184.758316 184.944489 185.130142 185.315292 185.499908 185.684036 185.867630 ];
x0 = [83 71 64 69 69 64 68 59 81 91 57 65 58 62];
y0 = [183 168 171 178 176 172 165 158 183 182 163 175 164 175];

plot(x, y);
hold;
plot(x0, y0, 'k *');

print("-r60", "power-regression.png");