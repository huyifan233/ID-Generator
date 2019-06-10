package utils




func ConvertStringTobyte32(str string)([32]byte){

	var bytes32 [32]byte
	sym := 32
	for i := len(str)-1; i>=0 ;i--{
		bytes32[sym-1] = str[i]
		sym--
	}

	return bytes32
}

func Add(numa string, numb string)(string){

	for len(numa) > len(numb){
		numb = "0"+numb
	}

	for len(numb) > len(numa){
		numa = "0"+numa
	}
	ans := ""

	lena, lenb, flag := len(numa)-1, len(numb)-1, 0

	for lena >= 0 && lenb >= 0{
		midNumber := int(numa[lena]) + int(numb[lenb]) + flag
		ans = string(midNumber % 10) + ans
		flag = midNumber / 10
		lena--
		lenb--
	}
	if flag != 0{
		ans = string(flag) + ans
	}
	return ans
}