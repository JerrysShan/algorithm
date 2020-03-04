var intersect = function (nums1, nums2) {
    const arr = [];
    const m = new Map();
    nums1.forEach(element => {
        if (m.has(element)) {
            m.set(element, m.get(element) + 1);
        } else {
            m.set(element, 1);
        }
    });
    nums2.forEach(ele => {
        const count = m.get(ele) || 0;
        if (count > 0) {
            arr.push(ele);
        }
        m.set(ele, count - 1);
    });
    return arr;
};

console.log(intersection([1, 2, 2, 1], [2, 2]));

console.log(intersection([4, 9, 5], [9, 4, 9, 8, 4]));