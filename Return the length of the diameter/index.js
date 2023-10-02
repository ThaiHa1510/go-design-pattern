let root = {
    val: 1,
    left: {
        val: 2,
        left: {
            val: 4,
            left: null,
            right: null
        },
        right: {
            val: 5,
            left: {
                val: 6,
                left: {
                    val: 10,
                    left: null,
                    right: null
                },
                right: {
                    val: 9,
                    left: null,
                    right: null
                }
            },
            right: null
        }
    },
    right: {
        val: 3,
        left: null,
        right: null
    }
}

diameterOfBinaryTree = function (root) {
    let maxD = 0

    function dfs(node) {
        if (node == null) {
            return 0
        }
        let left = dfs(node.left)
        let right = dfs(node.right)
        let currMax =Math.max(left,right)
        maxD = Math.max(maxD, currMax)
        console.log(`val:${node.val} , currMax:${currMax}, maxD: ${maxD}`)
        return Math.max(left + 1, right + 1)

    }
    dfs(root)
    return maxD
}
console.log(diameterOfBinaryTree(root))