import java.io.*;
import java.util.*;

public class Solution {
    
    boolean isBST(int left, int right, Node n) {
        if (n == null) {
            return true;
        }
        if (n.data > left && n.data < right) {
            return isBST(left, n.data, n.left) && isBST(n.data, right, n.right);
        } else {
            return false;
        }
    }

    boolean checkBST(Node root) {
        return isBST(Integer.MIN_VALUE, Integer.MAX_VALUE, root);        
    }
    class Node {
        int data;
        Node left;
        Node right;
    }
    
    public Node buildTree(String [] arr, int from, int to) {
        int middle = from + (to - from) / 2;
        Node node = new Node();
        node.data = Integer.valueOf(arr[middle]);
        if (middle > from) {
            node.left = buildTree(arr, from, middle - 1);
            node.right = buildTree(arr, middle + 1, to);
        }
        return node;
    }

    public static void main(String[] args) throws IOException {
        try (BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
             PrintWriter out = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)))) {
            int maxHeight = Integer.valueOf(in.readLine());
            String [] line = in.readLine().split(" ");
            Solution solution = new Solution();
            Node root = solution.buildTree(line, 0, line.length);
            if (solution.checkBST(root)) {
                out.println("Yes");
            } else {
                out.println("No");
            }
        }
    }
}
