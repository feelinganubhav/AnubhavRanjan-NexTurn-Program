export function capitalize(word) {
    if (!word) return "";
    return word[0].toUpperCase() + word.slice(1);
}
  
export function reverseString(str) {
    if (!str) return ""; // Handle null or undefined
    return str.split("").reverse().join("");
}
  
  
  