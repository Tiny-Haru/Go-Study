package main

import "fmt"

func main() {
	//map은 순서 없이 키와 값이 쌍을 이룬 것들의 집합
	var m map[string]int
	//map[A]B에서 A를 key로, B를 값으로 가짐.
	//map에서도 zero value는 슬라이스와 같이 nil이다.

	fmt.Println(m == nil) //true

	m2 := map[string]int{}
	//slice와 비슷하게 리터럴을 사용하여 초기값을 정의
	//string을 key로, int를 값으로 가지는 길이가 0인 map
	fmt.Println(m2 == nil) //false

	m3 := make(map[string]int)
	fmt.Println(m3) //map[]
	//empty map
	//map의 길이는 key-value 매핑이 변화할 때 바뀌므로, 별도로 길이를 설정할 수 없다.

	mapping := map[string]int {
		"key1": 123,
		"key2": 1234,
	}
	//리터럴을 통한 초기화
	fmt.Println(mapping) //map[key1: 123 key2: 1234]

	mapping["key3"] = 1234
	//요소 삽입 또는 수정

	v, exist := m["key3"]
	//요소 가져오기
	//map을 읽는 경우 해당 key의 값과 존재 여부를 함께 리턴함.

	fmt.Println(v, exist) //0 false

	delete(mapping, "key3")
	//요소 지우기

}
