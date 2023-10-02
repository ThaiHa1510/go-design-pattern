function findMaxDiff(arr, n){
    let min = Number.MAX_VALUE
    let max = Number.MIN_VALUE
    for(i=0;i<n;i++){
        if ((arr[i]-i)<min){
            min = arr[i]-i;
        }
        if ((arr[i]-i)>max){
            max = arr[i]-i;
        }
    }
    return max-min
}
let arr = [-1, -2, -3, 4, 10]
let n = arr.length;
console.log(findMaxDiff(arr, n));