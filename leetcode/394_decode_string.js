function decodeString(s) {
    let ret = '';
    let multi = 0;
    const strs = [];
    const ints = [];
    for (let i = 0; i < s.length; i++) {
        let temp = s[i];
        if (temp === '[') {
            strs.push(ret);
            ints.push(multi);
            ret = '';
            multi = 0;
        } else if (temp === ']') {
            let s = '';
            const t = ints.pop();
            for (let j = 0; j < t; j++) {
                s += ret;
            }
            ret = strs.pop() + s
        } else if (temp >= "0" && temp <= "9") {
            multi = multi * 10 + Number.parseInt(temp)
        } else {
            ret = ret + temp
        }
    }
    return ret;
}

// console.log(decodeString('3[a]2[bc]'));
// console.log(decodeString('3[a2[c]]'))