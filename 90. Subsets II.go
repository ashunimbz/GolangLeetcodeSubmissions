

var mp map[string]bool

func rec(cur int , nums []int , curString string ){
    
    if cur == len(nums)  {  
        mp[curString] = true 
        return 
    }
    
    //take it 
    newString := ""
    if curString  == "" {
        newString = strconv.Itoa(nums[cur]) 
    }else{
        newString = curString + "," + strconv.Itoa(nums[cur])     
    }    
    rec(cur + 1, nums , newString)
    
    
    //dont take it   
    rec(cur + 1, nums , curString)
    
    
}


func subsetsWithDup(nums []int) [][]int {
    
    mp = make(map[string]bool)
    
    sort.Ints(nums)
    rec(0,nums,"")
    
    ans := [][]int{}
    
    for key := range mp {
        
        sli := strings.Split(key , ",")
        
        innerSli := []int{}
        for _ ,ele := range sli {
            if ele == "" {
                continue
            }
            i , _ := strconv.Atoi(ele)
            
            innerSli = append(innerSli , i )
        }
        ans = append(ans , innerSli)
    }
    return ans 
    
} 