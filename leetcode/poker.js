
function poke(arr) {
    const out = [];
    let i = 1;
    while (arr.length) {
        if (i % 2 !== 0) {
            out.unshift(arr.pop());
        } else {
            arr.unshift(arr.pop());
        }
        i++;
    }
    return out;
}

function reverse(arr) {
    const out = [];
    let i = 1;
    while (arr.length) {
        if (i % 2) {
            out.push(arr.shift());
        } else {
            out.push(out.shift());
        }
        i++;
    }
    return out;
}


// console.log(poke([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]));

console.log(poke([1, 2, 3, 4, 5, 6]));

console.log(reverse([3, 1, 5, 2, 4, 6]));