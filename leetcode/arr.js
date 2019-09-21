
function getIndex(arr, num) {
    const map = new Map();
    for (const index in arr) {
        map.set(arr[index], index);
    }

    for (const item of arr) {
        if (map.has(num - item)) {
            return [map.get(item), map.get(num - item)]
        }
    }
    return [];

}

console.log(getIndex([2, 7, 11, 13], 18));