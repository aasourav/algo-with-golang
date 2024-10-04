the problem statement is : given an array [1, -1 , 3, 5, 6 , 10], find arr[i]+arr[j] == 11

### approach 1 (Hash Table):
time O(n)
space O(n)

    hash_table = {
        
    }
    y = arr[i] + 11
 if hash_table[y] then return the result (arr[i], y)
 else push arr[i] into the hash_table

### approach 2 (two pointer)
 time O(nlogn)
 space O(1)

 two pointer works only for sorted array . so we have to sort the array first.
 now  our array will be [-1, 1, 3, 5, 6, 10]

 now we have two section Left=0(first index) and Right=len(arr)-1 (last index)

run a loop :
if arr[left]+arr[right]==11
    return arr[left], arr[right]
else if arr[left]+arr[right] < 11
    then left++
else if arr[left]+arr[right] > 11
    then right++

if left >= right return false


 