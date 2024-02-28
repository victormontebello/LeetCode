isPalindrome = function(num) {
    if (num < 0) return false;
    return num === reverse(num);
}  

reverse = function(num) {
    var reversed = 0;
    while (num > 0) {
        reversed = (reversed * 10) + (num % 10);
        num = Math.floor(num / 10);
    }
    return reversed;
}
