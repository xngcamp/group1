package probl1

//Maxlength
func Maxlength(a []int) int {
	arr_length:=len(a)
	max:=1
	end:=-1
	len:=1
	for i:=0;i<arr_length;i++{
		if a[i+1] - a[i]==1{
			len++
		} else{
		if len>max{
			max=len
			end=i;
		}
		}
	}
	if end==-1{
		end =arr_length-1
		max=len
	}
	result :=max-2
	if a[end]==1000{
		result+=1;
	}
	if a[end-max+1] == 1{
		result+=1
		}
	return result
}

