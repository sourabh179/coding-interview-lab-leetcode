func rotate(nums []int, k int)  {
    if k==0||len(nums)==0{
        return
    }
    if k>len(nums){
        k=k%len(nums)
    }
    var arr []int
    if k<len(nums) && k!=0{
        for i:=len(nums)-k;i<len(nums);i++{
            arr=append(arr,nums[i])
        }

        for i:=0;i<len(nums)-k;i++ {
            arr=append(arr,nums[i])
        }
    }

    for i:=0;i<len(arr);i++{
        nums[i]=arr[i]
    }
}
