package rangeSumBS

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil {
        return 0
    }

    var result int
    if root.Val < L {
        result += rangeSumBST(root.Right, L, R)
    }
    if root.Val > R {
        result += rangeSumBST(root.Left, L, R)
    }
    if root.Val >= L && root.Val <= R {
        result += root.Val
        if root.Val != R {
            result += rangeSumBST(root.Right, L, R)
        }
        if root.Val != L {
            result += rangeSumBST(root.Left, L, R)
        }
    }
    return result
}
