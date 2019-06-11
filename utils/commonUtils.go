package utils

func ConvertStringTobyte32(str string)([32]byte){

	var bytes32 [32]byte
	for i := 0; i< len(str); i++{
		bytes32[i] = str[i]
	}

	return bytes32
}

func GetValidLen(str [32]byte)(int){

	validLen := 0
	for i:=0; i<len(str); i++{
		if str[i] == 0{
			break
		}
		validLen++
	}
	return validLen

}


func Add(numa string)(string){

	numb := config.Ec.EthereumIdStep

	for len(numa) > len(numb){
		numb = "0"+numb
	}

	for len(numb) > len(numa){
		numa = "0"+numa
	}
	ans := ""

	lena, lenb, flag := len(numa)-1, len(numb)-1, 0

	for lena >= 0 && lenb >= 0{
		midNumber := int(numa[lena]-'0') + int(numb[lenb]-'0') + flag
		ans = string(midNumber % 10 + '0') + ans
		flag = midNumber / 10
		lena--
		lenb--
	}
	if flag != 0{
		ans = string(flag + '0') + ans
	}
	return ans
}