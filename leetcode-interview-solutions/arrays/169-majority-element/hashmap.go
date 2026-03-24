// main function can be written. here only logical function has been given

func majorityElement(nums []int) int {
    mp:=make(map[int]int)
    max:=0
    if(len(nums)==1){
       return nums[0]
    }
    for _,v:=range nums{
        if mp[v]==0{
            mp[v]=1
        }else{
            mp[v]=mp[v]+1
            if mp[v]>len(nums)/2{
                max=v
                break
            }
        }
    }
    return max
}
