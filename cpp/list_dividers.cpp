#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

int main()
{
    ll n = 0;

    cin >> n;
    if (n == 0) {
        cout << "Z\n";
        return (0);
    }
    n = abs(n);

    ll root = sqrt(n);

    for (ll i = 1; i <= root; i++) {
        ll ratio = n / i;
        if (ratio * i == n) {
            cout << i << '\n';
            if (ratio != i) {
                cout << ratio << '\n';
            }
        }
    }
}