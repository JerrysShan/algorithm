
var intersection = function (nums1, nums2) {
    const set = new Set();
    nums1.forEach(element => {
        if (nums2.indexOf(element) > -1) {
            set.add(element);
        }
    });
    return [...set];
};

console.log(intersection([1, 2, 2, 1], [2, 2]));

console.log(intersection([4, 9, 5], [9, 4, 9, 8, 4]));