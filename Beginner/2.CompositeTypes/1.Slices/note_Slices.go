// 										Notes on Slices
//  Ways to declare slices
//   1. using slice literal
      var x := []int{1,2,3}
//   2. using nil
      var x []int{}  //nil slice 
//   3. using make
      var x = make([]int, length, capacity) //to preassign type,length and optionally capacity to increase the speed 

        x := make([]int, 5) 
            // This creates an int slice with a length of 5 and a capacity of 5. Since it has a length of
            // 5, x[0] through x[4] are valid elements, and they are all initialized to 0.

         x = append(x, 10) 
        //   [0 0 0 0 0 10] length=6 capacity=10
        
         x := make([]int, 0, 10) 
        //   This is a slice with 0 length but non-nil

//  ❌ specify capacity with numeric literal or constant that is less than length = compile time error
//  ❌  specify capacity with variable that is less than length , program will panic at runtime

 var data []int  // Nil slice
 var x = []int{} // Non-nil zero length slice ... comparing it to nil returns false ...useful when converting slice to json

 
//  Slicing slices
        x := []int{1, 2, 3, 4}
        y := x[:2]
//   output
    x: [1 2 3 4]
    y: [1 2] 
 // slices share storage ,taking a slice from a slice creates another varibale that shares the same memory

//  ## to create a slice that's independent of original use copy function
        // copy(destination ,source)
    x := []int{1, 2, 3, 4}
    y := make([]int, 4) // 0 0 0 0
      copy(y, x)
    // y becomes  [1,2,3,4]