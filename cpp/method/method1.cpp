#include<iostream>
using namespace std;

//procedure (gaperlu di panggil lagi atau di return, langsung cout aja)
void siswa(){
    string nama = "Farel";
    string jk = "Laki-laki";
    cout<<nama<<" adalah seorang "<<jk<<endl;
}

//function (harus manggil lagi fungsi nya yaitu di return)
string siswi(){
    string nama = "Khai";
    string jk = "Perempuan";
    return(nama+" adalah seorang "+jk); //kalo return makenya + bukan << atau >>
}

int main(){
    siswa(); //procedure (gaperlu make cout langsung panggil namanya aja)
    cout<<siswi()<<endl; //function, make cout baru manggil namanya)
}