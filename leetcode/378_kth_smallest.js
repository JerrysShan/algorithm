var kthSmallest = function (matrix, k) {
    let arr = [];
    for (let i = 0; i < matrix.length; i++) {
        arr = arr.concat(matrix[i]);
    }

    arr.sort((a, b) => a - b);
    return arr[k - 1];
};


console.log(kthSmallest([
    [1, 5, 9],
    [10, 11, 13],
    [12, 13, 15]
], 6));