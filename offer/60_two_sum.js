var twoSum = function (n) {
    const dp = [];
    for (let i = 1; i <= n; i++) {
        dp[i] = new Array(6 * n + 1).fill(0);
    }
    for (let i = 1; i <= 6; i++) {
        dp[1][i] = 1;
    }
    for (let i = 2; i <= n; i++) {
        for (let j = i; j <= 6 * i; j++) {
            for (let cur = 1; cur <= 6; cur++) {
                if (j - cur <= 0) {
                    break;
                }

                dp[i][j] += dp[i - 1][j - cur];
            }
        }
    }
    const all = Math.pow(6, n);
    const ret = [];
    for (let i = n; i <= 6 * n; i++) {
        ret.push((dp[n][i] / all).toFixed(5))
    }

    return ret;
};

// console.log(twoSum(1));
console.log(twoSum(9));
// console.log(twoSum(3));