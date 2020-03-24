var minIncrementForUnique = function (A) {
    A.sort((a, b) => a - b);
    let count = 0;
    for (let i = 1; i < A.length; i++) {
        while (A[i] <= A[i - 1]) {
            A[i] = A[i] + 1;
            count++;
        }
    }

    return count;
};

console.log(minIncrementForUnique([1, 2, 2]));
// console.log(minIncrementForUnique([3, 2, 1, 2, , 1, 7]))