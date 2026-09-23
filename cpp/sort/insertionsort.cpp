#include <iostream>
using namespace std;

int main(){
    int x[] = {3, 2, 6, 5, 4, 10, 1, 9, 8, 7};
    int j = size(x) / sizeof(*x);
    int temp;

    cout << "x: ";
    for (int a = 0; a < j; a++){
        cout << x[a] << " ";
    }

    cout << endl;
    cout<< endl;

    for(int b = 1; b < j; b++){
        for(int c = b; c > 0; c--){
            if(x[c] < x[c-1]){
                temp = x[c];
                x[c] = x[c-1];
                x[c-1] = temp;
            } else {
                break;
            }
            cout << "b: " << b << " ->";
            for(int i = 0; i < j; i++){
                cout << x[i] << " ";
            }
        }
        cout << endl;
    }
    cout << endl;
    cout << endl;
    for(int d = 0; d < j; d++){
        cout << x[d] << " ";
    }
} 