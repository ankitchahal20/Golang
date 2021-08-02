// Your First C++ Program

#include <iostream>
#include "math.h"
using namespace std;

// int ribbonCutting(int n,int a,int b,int c){
//   //base condition
//   if (n<0 || n<0 || n<0){
//     return INT_MIN;
//   }
//   if ( n==0){
//     return 0;
//   }
//
//
//   return 1 + max(ribbonCutting(n-a,a,b,c), max( ribbonCutting(n-b,a,b,c), ribbonCutting(n-c,a,b,c)));
//   //return ans;
//
// }

int dp[4001];

int ribbonCutting(int n,int a,int b,int c){
  //base condition



  if (n<0 || n<0 || n<0){
    return INT_MIN;
  }
  if (n==0){
    return 0;
  }


  if(dp[n] != 0){
    //cout<<"dp[n]"<<dp[n];
    return dp[n];
  }

  dp[n]= 1 + max(ribbonCutting(n-a,a,b,c), max( ribbonCutting(n-b,a,b,c), ribbonCutting(n-c,a,b,c)));

  return dp[n];

}

int main() {
    int n,a,b,c;
    cin>>n>>a>>b>>c;
    int ans = ribbonCutting(n,a,b,c);
    std::cout <<ans<<endl;
    return 0;
}
