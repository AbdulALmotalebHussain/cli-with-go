#include <iostream>
#include <set>
#include <iterator>
using namespace std;

int main() {
    set <int, greater <int> > Tele1;
    Tele1.insert(40); Tele1.insert(30); Tele1.insert(60);
    Tele1.insert(20); Tele1.insert(10);
    Tele1.insert(50); Tele1.insert(50); // only one 50 will be added to the set

    set <int> :: iterator itr;
    cout << "\nThe set Tele1 is : ";
    for (itr = Tele1.begin(); itr != Tele1.end(); ++itr) {
        cout << '\t' << *itr;
    }
    cout << endl;

    set <int> Tele2(Tele1.begin(), Tele1.end());
    cout << "\nThe set Tele2 after assign from Tele1 is : ";
    for (itr = Tele2.begin(); itr != Tele2.end(); ++itr) {
        cout << '\t' << *itr;
    }
    cout << endl;

    cout << "\nTele2.erase(50) : ";
    Tele2.erase(50);
    for (itr = Tele2.begin(); itr != Tele2.end(); ++itr) {
        cout << '\t' << *itr;
    }
    cout << endl;

    cout << "\nTele2 after removal of elements less than 30 : ";
    Tele2.erase(Tele2.begin(), Tele2.find(30));
    for (itr = Tele2.begin(); itr != Tele2.end(); ++itr) {
        cout << '\t' << *itr;
    }
    cout << endl;

    // lower bound and upper bound for set Tele1
    cout << "Tele1.lower_bound(40) : " << *Tele1.lower_bound(40) << endl;
    cout << "Tele1.upper_bound(40) : " << *Tele1.upper_bound(40) << endl;

    // lower bound and upper bound for set Tele2
    cout << "Tele2.lower_bound(40) : " << *Tele2.lower_bound(40) << endl;
    cout << "Tele2.upper_bound(40) : " << *Tele2.upper_bound(40) << endl;
    return 0;
}
