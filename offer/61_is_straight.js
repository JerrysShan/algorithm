var isStraight = function (nums) {
    let count = 0;
    let sum = 0;
    nums.sort((a, b) => a - b);
    for (let i = 0; i < nums.length - 1; i++) {
        if (nums[i] == 0) {
            count++;
            continue;
        }
        const t = nums[i + 1] - nums[i];
        if (t == 0) {
            return false;
        }
        if (t > 1) {
            sum += nums[i + 1] - nums[i] - 1;
        }
    }
    if (count >= sum) {
        return true;
    }
    return false
};