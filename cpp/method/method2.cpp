#include <iostream>
using namespace std;

void siswa(string nama, string jk){
    cout<<nama<<" adalah seorang mahasiswa Kelas "<<jk<<endl;
}

int main(){
    string nama, kelas;
    cout<< "Nama = "; 
    getline(cin, nama);
    cout<< "Kelas = ";
    getline(cin, kelas);

    siswa(nama, kelas); 
}