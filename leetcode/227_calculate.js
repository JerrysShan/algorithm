var calculate = function (s) {
    const stack = [];
    s = s.replace(/\s+/g, '');
    for (let i = 0; i < s.length;) {
        if (/\d/.test(s[i])) {
            let char = s[i];
            while (i + 1 < s.length && /\d/.test(s[i + 1])) {
                char += s[i + 1];
                i++;
            }
            if (!stack.length) {
                stack.push(char);
            } else {
                let num2 = char;
                while (stack.length && (i === s.length - 1 || !calculateOrder(s[i + 1], stack[stack.length - 1]))) {
                    const sign = stack.pop();
                    const num1 = stack.pop();
                    switch (sign) {
                        case '+': num2 = Number.parseInt(num1) + Number.parseInt(num2); break;
                        case '-': num2 = Number.parseInt(num1) - Number.parseInt(num2); break;
                        case '*': num2 = Number.parseInt(num1) * Number.parseInt(num2); break;
                        case '/': num2 = Number.parseInt(Number.parseInt(num1) / Number.parseInt(num2)); break;
                    }
                }
                stack.push(num2);
            }
            i++;
        } else {
            stack.push(s[i++]);
        }
    }
    return stack.pop();
};

function calculateOrder(sign1, sign2) {
    if ((sign1 === '*' || sign1 === '/') && (sign2 === "+" || sign2 === '-')) {
        return true;
    }
    return false;
}

console.log(calculate('3+2*2'));
console.log(calculate(' 3+5 / 2 '));

console.log(calculate('42'));
console.log(calculate("1+2*5/3+6/4*2"));