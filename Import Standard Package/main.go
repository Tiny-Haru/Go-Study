package main

//소프트웨어의 중요한 부분 중 하나는 코드 재사용을 통한 중복의 제거이다.
//패키지는 함수처럼 코드 재소용을 돕기 위한 수단이다.

import (
	//외부 패키지를 가져와 사용하기 위해 import를 사용한다.
	//기본형은 import "[name]"이며, 이 코드와 같이 괄호로 묶어 한 번에 여러 패키지를 import 할 수도 있다.
	"fmt"
	mt "math"
	//Go언어의 내장 패키지인 fmt와 math를 import 했으며, 7행과 같이 패키지 이름 앞에 문자열을 정의하여 alias를 사용할 수 있다.
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(mt.Abs(-3))
	//외부 패키지의 모든 요소 이름이 대문자로 시작한다.
	//패키지의 요소를 다른 패키지에서 사용할 수 있게 하려면, 요소의 이름이 이처럼 대문자로 시작해야 한다.
	//이것을 exported name이라고 부른다.
}
