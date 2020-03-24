var lastRemaining = function (n, m) {
    const arr = Array.from(new Array(n), (val, index) => {
        return index;
    });
    let index = 0;
    while (arr.length != 1) {
        index = (index + m - 1) % arr.length
        arr.splice(index, 1);
    }
    return arr[0];
};

console.log(lastRemaining(5, 3));
console.log(lastRemaining(10, 17));