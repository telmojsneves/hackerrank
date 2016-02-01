#include <stdio.h>

int main()
{
    int t;
    int i,j;
    scanf("%d",&t);
    for (;t!=0;t--){
        int n,k,index;
        long long int max=0;

        scanf("%d%d\n",&n,&k);
        char a[n];
        int b[n];

        long long int x=1;

        for(i=0;i<n;i++){
            scanf("%c\n",&a[i]);
            b[i] = a[i] - '0';
        }

        for(i=0;i<n-k;i++)
        {
            x=1;
            for(j=i;j<i+k;j++)
                x = x * b[j];
            if(x>max)
            {
                max=x;
                index=i;
            }
        }

        printf("%lld\n",max);
    }
    return 0;
}
