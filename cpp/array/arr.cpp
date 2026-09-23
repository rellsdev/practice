#include <iostream>
using namespace std;

int main(){
    int jml_item;
    cout << "Masukkan jumlah item: ";
    cin >> jml_item;
    
    int harga[jml_item];
    int total;

    for (int b = 0; b < jml_item; b++){
        cout << "Harga item ke-" << b << ": ";
        cin >> harga[b];
    }

    cout<<"==============================="<<endl;
    cout<<"Jumlah Item: " << jml_item << endl;
    for (int a = 0; a < jml_item; a++){
        cout << "Harga item ke-" << a << ": Rp. " << harga[a] << endl;
        total += harga[a];
    }
    cout<<"==============================="<<endl;
    cout << "Total: Rp. " << total << endl;
}