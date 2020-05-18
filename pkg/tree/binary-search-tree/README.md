# What is binary search tree?

Binary search tree is a data structure that allows efficient search, adding and removing of items. A binary tree has at most two child nodes, usually distinguished as "left" and "right". Nodes with children are parent nodes. Outside the tree, there is often a reference to the "root" node (the ancestor of all nodes), if it exists. Any node in the data structure can be reached by starting at root node and repeatedly following references to either the left or right child.

Key properties of tree include:

- There is no duplicate nodes
- The left node contains key that is smaller
- The right node contains key that is larger
- Contains a `key` - The key is used as the identifying property to order the entries appropriately as they are inserted within the tree so that a search in the binary tree can be achieved efficiently.
- Contains a `value` - The value can be any kind of data type, it could store a string which can be the text of a page of a book or a boolean value to be used in a finite state machine.


## Insert into a tree 

A new key is always inserted at the leaf. We start from the root and keep going until we found a leaf node and then insert it as its child.

```
         100                               100
        /   \        Insert 40            /    \
      20     500    --------->          20     500 
     /  \                              /  \  
    10   30                           10   30
                                              \   
                                              40
```

## Search in a tree

Start from the root, and check current node and go to left or right tree depend on the value being searched.

## Delete in a tree 

1. Node to be deleted is a leaf node - just remove the node 

```

             50                             50
           /     \         delete(20)      /   \
          30      70       --------->    30     70 
         /  \    /  \                     \    /  \ 
       20   40  60   80                   40  60   80
```

2. Node to be deleted has one child only - remove the child to the position of the deleted node

```
              50                            50
           /     \         delete(30)      /   \
          30      70       --------->    40     70 
            \    /  \                          /  \ 
            40  60   80                       60   80

```

3. Node to be deleted has 2 children - We ned to find the smallest node from the right tree (the most left leaf node). In this case, it is node with key 60.


```
              50                            60
           /     \         delete(50)      /   \
          40      70       --------->    40    70 
                 /  \                            \ 
                60   80                           80

```

### Time and space Complexity

| Operation     | Big O       | 
| ------------- |-------------| 
| Insert        | O(logN)     | 
| Search        | O(logN)     | 
| Delete        | O(logN)     |

Why it is O(logN)? The number of elements is N, and it has a power of 2 (as it is binary: left or right). Let's say we have 8 nodes, if we eliminate half of the nodes by going into a subtree, we have 4 nodes. If we keep going we are left with 2 nodes and again we finally reached the desired node. Thus O(log8) = O(3), which is the worst case. 

It is the same case regardless whether it's insert, search or delete option.