
export function getElement(arr, index) {
    if (typeof index !== "number" || !Number.isInteger(index)) {
        throw new Error("Index must be an integer");
    }
    if (index < 0 || index >= arr.length) {
      throw new Error("Index out of bounds");
    }
    return arr[index];
  }
  