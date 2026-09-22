#include <iostream>
using namespace std;

int main(){
    int jml_arr;
    int angka[] = {90, 87, 45, 38, 78, 80, 120};
    jml_arr = sizeof(angka)/sizeof(*angka);

    for (int a = 0; a < jml_arr; a++){
        cout << angka[a] << endl;
    }

    cout << "Jumlah Array: " << jml_arr << endl;
}