#include <iostream>
using namespace std;

int main(){
    int jml_arr;
    cout<<"Masukkan Jumlah Index: ";
    cin>>jml_arr;
    
    int angka[jml_arr];
    //int jml_arr = sizeof(angka)/sizeof(*angka);

    for (int a = 0; a < jml_arr; a++){
        cout<<"Masukkan Nilai Index ke- "<<a<<": ";
        cin>>angka[a];
    }
    cout<<"==============================="<<endl;
    cout<<"Nilai Yang Terimpan"<<endl;
    cout<<"==============================="<<endl;

    for(int b = 0; b < jml_arr; b++){
        cout<<"Nilai Index ke- "<<b<<" : "<<angka[b]<<endl;
    }
}