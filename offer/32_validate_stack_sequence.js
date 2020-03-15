
var validateStackSequences = function (pushed, popped) {
    if (!pushed.length || !popped.length) {
        return false;
    }
    const stack = [];
    for (const item of pushed) {
        stack.push(item);
        while (stack.length && (stack[stack.length - 1] == popped[0])) {
            popped.shift();
            stack.pop();
        }
    }
    return !stack.length;
};

console.log(validateStackSequences([1, 2, 3, 4, 5], [4, 5, 3, 2, 1]));

console.log(validateStackSequences([1, 2, 3, 4, 5], [4, 3, 5, 1, 2]));
