
function backtrack(str, index, ret, arr, flags) {
    if (index === str.length) {
        return ret.push(arr.join(''));
    }
    for (let i = 0; i < str.length; i++) {
        if (!flags[i]) {
            flags[i] = true;
            arr[index] = str[i];
            backtrack(str, index + 1, ret, arr, flags);
            flags[i] = false;
        }
    }
}

function permute(str) {
    const ret = [];
    backtrack(str, 0, ret, [], []);
    return ret;
}
