
var myAtoi = function (str) {
    str = str.trim();
    const reg = /^[-+]?\d+/;
    if (!reg.test(str)) {
        return 0;
    }
    let temp = reg.exec(str);
    let ret = Number.parseInt(temp[0]);
    // console.log(ret)
    if (ret < -Math.pow(2, 31)) {
        return -Math.pow(2, 31);
    }
    if (ret > Math.pow(2, 31) - 1) {
        return Math.pow(2, 31) - 1;
    }
    return ret;
};

console.log(myAtoi("42"));
console.log(myAtoi("  -42"));
console.log(myAtoi("words 42"))
console.log(myAtoi("-91283472332"));
console.log(myAtoi("4193 with words"));
