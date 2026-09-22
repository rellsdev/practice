#include <iostream>
using namespace std;

bool lulus(int nilai){
    if (nilai > 75){
        return true;
    } else {
        return false;  
    }
}

void data_siswa(string nama, int nilai){
    lulus(nilai);
    cout<<"Nama    : "<<nama<<endl;
    cout<<"Nilai   : "<<nilai<<endl;
    cout<<"Status  : ";
    if (lulus(nilai)){
        cout<<"Lulus"<<endl;
    } else {
        cout<<"Tidak Lulus"<<endl;
    }
}

int main(){
    string nama;
    int nilai;
    cout<<"Masukkan Nama    : ";
    getline(cin, nama);
    cout<<"Masukkan nilai   : ";
    cin>>nilai;
    data_siswa(nama, nilai);
}