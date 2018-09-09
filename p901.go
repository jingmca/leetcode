package main

//StockSpanner ...
type StockSpanner struct {
	call  int
	price []int
	spans []int
}

//Constructor ...
func Constructor() StockSpanner {
	return StockSpanner{call: 0, price: []int{}, spans: []int{}}
}

func (this *StockSpanner) Next(price int) int {

	if this.call == 0 || this.price[this.call-1] > price {
		this.price = append(this.price, price)
		this.spans = append(this.spans, 1)
		this.call++
		return 1
	}

	k := 1
	for i := this.call - 1; i >= 0; {
		if this.price[i] <= price {
			if this.spans[i] > 1 {
				k += this.spans[i]
				i -= this.spans[i]
			} else {
				k++
				i--
			}
		} else {
			break
		}
	}

	this.price = append(this.price, price)
	this.spans = append(this.spans, k)
	this.call++

	return k
}

func main() {
	obj := Constructor()
	
}
