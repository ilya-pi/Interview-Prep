import java.util.*;
import java.io.*;

class Node {
    Node left;
    Node right;
    int data;
    
    Node(int data) {
        this.data = data;
        left = null;
        right = null;
    }
}

class Solution {
   
  	public static void preOrder( Node root ) {
      
    	if( root == null)
        	return;
      
        System.out.print(root.data + " ");
        preOrder(root.left);
        preOrder(root.right);
     
    }

 /* Node is defined as :
 class Node 
    int data;
    Node left;
    Node right;
    
    */

	public static Node insert(Node root, int data) {
        if (root == null) {
            Node n = new Node(data);
            return n;
        }
        if (data < root.data) {
            if (root.left != null) {
                insert(root.left, data);
                return root;
            } else {
                Node n = new Node(data);
                root.left = n;
                return root;
            }
        } else if (data > root.data) {
            if (root.right != null) {
                insert(root.right, data);
                return root;
            } else {
                Node n = new Node(data);
                root.right = n;
                return root;
            }
        }
        // We already have that value in the tree 
        return root;
    }

	public static void main(String[] args) {
