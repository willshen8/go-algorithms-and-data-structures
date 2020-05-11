# What is linked list?

A linked list is a linear collection of data elements, and its order is not based on physical placement in memory. Instead each element (node) points to the next one. In the most basic form, a node contains a data and a reference/link to the next node. 

## What is doubly linked list?

A doubly linked list contains links to elements before and after, as oppose to normal linked list. 

## Doubly linked list vs linked list?

Doubly linked list contains two links, so it has the advantages of being able to traverse in both directions. It is also easy to insert new nodes before existing ones. 

The disadvantages of doubly linked list is obviously having to manage extra pointer, hence more memory.

## Run time complexities for doubly linked list  

| Operation     | Big O       | Description |
| ------------- |-------------| ------------- |
| Traverse      | O(n)        | Traverse the list in order one by one |
| PushFront     | O(1)        | Insert new node at the front of the list |
| PushBack      | O(1)        | Insert new node at the back of the list |
| PopFront      | O(1)        | Delete node from the front of the list |
| PopBack       | O(1)        | Delete node from the back of the list |

## Why do we use linked list instead of array/slice

Similar to linked list, array can achieve the same result by able to access data before and after using index very quickly. In fact, we can easily access an item at a certain index using `array[i]` for example, so what's different these two and why should I use linked list? 

### Differences between linked list and array

Arrays can created with a fixed size, and any unused elements could be potentially wasted. Arrays are organised in continuous block of memory. If original array is too small to store additional elements,  then we need to create a new array and copy the old data into new array. 

Linked list on the other hands consists of data elements that are stored in different memory addresses. We can create and delete nodes without having to worry about memory. There is also no reserved/unused memory, so it's not wasted.

## Applications of linked list

Linked list can also be used to implement stack and queue, and it also can be used to implement graphs - adjacency list representation of graphs. 