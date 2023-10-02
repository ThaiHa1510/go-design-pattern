// JavaScript program to sort an array
// in Zig-Zag form

// Program for zig-zag conversion of array
function zigZag(arr, n)
{
    // sort the by using the sort function
    arr.sort();
    //traverse the array from 1 to n-1
    for(let i = 1; i <= n - 2; i=i+2)
    {
    // swap the current element with next element
        let temp = arr[i];
        arr[i] = arr[i+1];
        arr[i+1] = temp;
    }
}

function zigZag2(arr, n)
{
    for(let i = 0; i <= n - 2; i++)
    {
        if(i%2==0)
        {
            if(arr[i]>arr[i+1])
            {
                let temp = arr[i];
                arr[i] = arr[i+1];
                arr[i+1] = temp;
            }
        }
        else{
            if(arr[i]<arr[i+1])
            {
                let temp = arr[i];
                arr[i] = arr[i+1];
                arr[i+1] = temp;
            }
        }
    }
}

// Driver code
let arr = [ 4, 3, 7, 8, 6, 2, 1 ];
let n = arr.length;
zigZag2(arr, n);
// print the array
console.log(arr);

