#include <bits/stdc++.h>
#include <fstream>

using namespace std;

const string FIRST = "Pishty";
const string SECOND = "Lotsy";

int dp[(1 << 16)];

int bit(int mask, int i) /// i-th bit of mask
{
    return (mask & (1 << i)) > 0;
}


int a[5][5], p[5][5];

int s(int x1, int y1, int x2, int y2) ///sum at submatrix
{
    return p[x2][y2] - p[x1 - 1][y2] - p[x2][y1 - 1] + p[x1 - 1][y1 - 1];
}

/// calculating of Sprague–Grundy numbers
bool used[(1 << 16)];
int forClear[(1 << 16)], szForClear;

void markPosition()
{
    FILE * fp = fopen("test.txt", "w");
    dp[0] = 0;
    
    for(int mask = 1; mask < (1 << 16); mask++)
    {
        int firstNotUsed = 0;

        for(int i = 0; i < szForClear; i++)
            used[forClear[i]] = false;
	for(int i = 0; i < szForClear; i++) {
		fprintf(fp, "used[forClear[%d]]==%d, forClear[%d]=%d\n", i, used[forClear[i]], i, forClear[i]);
	}
        szForClear = 0;

        for(int i = 0; i < 4; i++)
            for(int j = 0; j < 4; j++)
                a[i + 1][j + 1] =  bit(mask, i * 4 + j);

	 char str[74];
         for(int i = 1; i <= 4; i++) {
            for(int j = 1; j <= 4; j++) {
		fprintf(fp, "%d ", a[i][j]);
	    }
	    fprintf(fp, "\n");
	}
        fprintf(fp, "\n");
	
	for(int i = 1; i <= 4; i++)
            for(int j = 1; j <= 4; j++)
                p[i][j] = a[i][j] + p[i - 1][j] + p[i][j - 1] - p[i - 1][j - 1];

        for(int x1 = 1; x1 <= 4; x1++)
            for(int y1 = 1; y1 <= 4; y1++)
                for(int x2 = x1; x2 <= 4; x2++)
                    for(int y2 = y1; y2 <= 4; y2++) {
                if(s(x1, y1, x2, y2) == (x2 - x1 + 1) * (y2 - y1 + 1))
                {
		    fprintf(fp, "(x1, y1) == (%d, %d)\n", x1, y1);
                    fprintf(fp, "(x2, y2) == (%d, %d)\n", x2, y2);

                    int newMask = mask;
		    int oldMask;
                    for(int i = x1; i <= x2; i++)
                        for(int j = y1; j <= y2; j++) {
			    oldMask = newMask; 
                            newMask -= (1 << ((i - 1) * 4 + j - 1));
			    fprintf(fp, "(i, j) == (%d, %d)\n", i, j);
			    fprintf(fp, "oldMask=%d, newMask=%d, val=%d\n", oldMask, newMask, 1 << ((i - 1) * 4 + j - 1));
		    }
		    fprintf(fp, "dp[%d]=%d, used[%d]=%d\n", newMask, dp[newMask], dp[newMask], used[dp[newMask]]);
                    used[dp[newMask]] = true;
                    forClear[szForClear++] = dp[newMask];
                    while(used[firstNotUsed]) firstNotUsed++;
                }
	}
	fprintf(fp, "dp[%d]=%d\n", mask, firstNotUsed);
        dp[mask] = firstNotUsed;
    }
    fclose(fp);
    exit(1);
}

const int MAXN = 1e5 + 10;

int n, m;

///Fenwick tree
int t[MAXN], b[MAXN];

int getPrefix(int r)
{
    int total = 0;
    for(; r > 0; r = (r & (r + 1)) - 1)
        total ^= t[r];
    return total;
}

int get(int l, int r)
{
    return getPrefix(r) ^ getPrefix(l - 1);
}

void update(int pos, int newMask)
{
    int oldMask = b[pos];
    for(int i = pos; i <= n; i = (i | (i + 1)))
        t[i] = t[i] ^ oldMask ^ newMask;
    b[pos] = newMask;
}

void solve()
{
    scanf("%d%d",&n, &m);

    for(int i = 1; i <= n; i++)
        t[i] = 0, b[i] = 0;
    for(int i = 1; i <= n; i++)
    {
        char c[4][4];
        for(int j = 0; j < 4; j++)
            scanf("%s", c[j]);

        int mask = 0;
        for(int x = 0; x < 4; x++)
            for(int y = 0; y < 4; y++)
                mask += (c[x][y] - '0') * (1 << (x * 4 + y));
        update(i, dp[mask]);
    }

    while(m--)
    {
        int type;
        scanf("%d", &type);
        if(type == 1)
        {
            int l, r;
            scanf("%d%d", &l, &r);
            cout << (get(l, r) > 0 ? FIRST : SECOND) << "\n";
        }
        else
        {
            int i;
            scanf("%d", &i);
            char c[4][4];
            for(int j = 0; j < 4; j++)
                scanf("%s", c[j]);

            int mask = 0;
            for(int x = 0; x < 4; x++)
                for(int y = 0; y < 4; y++)
                    mask += (c[x][y] - '0') * (1 << (x * 4 + y));
            update(i, dp[mask]);
        }
    }
}


int main()
{
    ios :: sync_with_stdio(0);

    markPosition();
    int tt;
    scanf("%d", &tt);
    while(tt--)
        solve();

    return 0;
}
