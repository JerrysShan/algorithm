var topKFrequent = function (nums, k) {
    const ret = [];
    const m = new Map();
    nums.forEach(ele => {
        if (m.has(ele)) {
            m.set(ele, m.get(ele) + 1);
        } else {
            m.set(ele, 1);
        }
    });
    let i = 0;
    for (const [key, val] of m) {
        if (i < k) {
            ret[i++] = key;
        } else {
            ret.sort((a, b) => m.get(a) - m.get(b));
            if (val > m.get(ret[0])) {
                ret[0] = key;
            }
        }
    }
    return ret;
};

console.log(topKFrequent([1, 1, 1, 2, 2, 3], 2));