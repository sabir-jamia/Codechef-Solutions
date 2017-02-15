package main 

import (
    "fmt"
)

func merge(int arr[], int l, int m, int r) {
     int arrL[m - l + 1], arrR[r - m + 2], i;
     
     for(i = 0; i < m - l + 1; i++) {
         arrL[i] = arr[i];
     }
     
     for(i = 0; i < r - m + 2; i++) {
         arrR[i] = arr[i];
     }
     
     int j =0, k =0;
     for(i = 0; i < r - l + 1; i++) {
         if(arrL[j] > arrR[k]) {
             k++;
             arr[i] = arrR[k];
         } else {
             arr[i] = arrL[j];
             j++;
         }
     }
}

func mergeSort(int arr[], int l, int r) {
    if (l < r)   {
        int m = l+(r-l)/2;
        mergeSort(arr, l, m);
        mergeSort(arr, m+1, r);
        merge(arr, l, m, r);
    }
}

func main() {
    var  numElements, arr[1000], i, j, k;
    reader := bufio.NewReaderSize(os.Stdin, 1000000)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanWords)
    testCases := nextInt(scanner)

    for i := 0; i < testCases; i++ {
        numElements = nextInt(scanner)
        for j = 0; j < numElements; j++) {
            scanf("%d", &arr[j]);
        }
        mergeSort(arr, 0, (numElements - 1));
        for (k =0; k < numElements; k++) {
            printf("%d ", arr[k]);   
        }
        printf("\n");
    }
    return 0;
}