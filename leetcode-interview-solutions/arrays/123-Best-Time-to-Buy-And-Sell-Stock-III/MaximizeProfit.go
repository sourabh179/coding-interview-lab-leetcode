func maxProfit(prices []int) int {
    buy1:=prices[0]
    sell1:=0
    buy2:=prices[0]
    sell2:=0

    for _,price:=range prices{
        if price<buy1{
            buy1=price
        }
        if price-buy1 > sell1{
            sell1=price-buy1
        }
        if price-sell1 < buy2{
            buy2=price-sell1
        }
        if price - buy2>sell2{
            sell2=price-buy2
        }
    
    }
    return sell2
}