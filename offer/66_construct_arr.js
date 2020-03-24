var constructArr = function (a) {
    const arr = [];
    for (let i = 0; i < a.length; i++) {
        let temp = 1;
        for (let j = 0; j < a.length; j++) {
            if (j !== i) {
                temp *= a[j]
            }
        }
        arr.push(temp);
    }
    return arr;
};
console.log(constructArr([1, 2, 3, 4, 5]));