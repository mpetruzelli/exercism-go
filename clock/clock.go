package clock
import "fmt"
const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct{
	h int
	m int
}

func New(hour, minute int) Clock {
	h,m := fix_time(hour,minute);	
	return Clock{h,m};
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m);
}

func (c Clock) Add(minutes int) Clock {
	return New(c.h,c.m+minutes);
}

func fix_time(h,m int) (int,int){
	//if m == -60 do not adjust impact on hours.
	if(m<0 && m==-60){
	}else if(m<0){//adjust impact on hours
		m-=60
	}
	h += m / 60 
	m=mod(m , 60)
	h=mod(h , 24)
	
	return h,m;
}

func mod(x, y int) int {

  mod := x % y; 
  if mod < 0 { 
    mod += y 
  }
  return mod
}




