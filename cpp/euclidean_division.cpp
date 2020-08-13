#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

int main()
{
    ll a = 0, b = 0;

    cin >> a >> b;
    if (b == 0) {
        cout << "division by 0!\n";
        return (0);
    }

    ll q = 0;

    if (a >= 0) {
        while (a >= b) {
            a -= b;
            q++;
        }
    } else {
        while (a < 0) {
            a += b;
            q--;
        }
    }

    ll r = a;

    cout << "q=" << q << '\n';
    cout << "r=" << r << '\n';
}