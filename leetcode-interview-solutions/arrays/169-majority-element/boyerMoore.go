//Only Fucntion is given for end to end testing main function has to be written.

func majorityElement(nums []int) int {
   ele:=nums[0]
   c:=1
   for _,v:=range nums{
    if ele==v{
        c++
    }else{
        c--
        if c<=0{
            ele=v
            c=1
        }
    }
   }
   return ele
}
