

function isValidBST(root) {
	let lastVal = Number.MIN_SAFE_INTEGER;
	const stack = [];
	while (root || stack.length) {
		if (root) {
			stack.push(root);
			root = root.left;
		} else {
			const temp = stack.pop();
			if (temp.val < lastVal) {
				return false;
			}
			lastVal = temp.val;
			root = temp.right;
		}
	}
	return true
}
