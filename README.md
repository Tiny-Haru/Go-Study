# Go-Study

## Go
![Gopher image](https://sdtimes.com/wp-content/uploads/2014/12/1212.sdt-github.jpg)

```
Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.
```
- 공식 홈페이지 : https://golang.org/

- 다운로드 : https://golang.org/dl/

- [The Go Playground](https://play.golang.org/)를 통해 설치 없이 웹에서 간단한 컴파일, 코드 공유가 가능하다.

- [A Tour of Go](https://go-tour-kr.appspot.com/#1)(한국어 버전)을 통해 Go의 기본적인 문법부터 웹 크롤러 만들기까지 다양한 예제를 통해 튜토리얼을 진행할 수 있다.

---

Go언어는 2007년 구글에서 개발을 시작하여 2012년에 버전 1.0을 완성한 오픈소스 프로그래밍 언어이다. 구글의 V8 엔진에 참여했던 Robert Griesemer, 유닉스 개발에 참여했던 Rob Pike, B 언어의 개발자인 Ken Thompson이 함께 개발했다.

고언어의 특징으로는 다음와 같은 것들이 있다.

- 바이너리 컴파일
- 타입 시스템 : 강타입, 정적 타이핑
- 명시적 타입 추론
- Scope : 정적, 블록 스코프
- 일급 함수
- 멀티 페러다임
- 구조체와 인터페이스, duck typing 기반 객체지향
- 별도의 VM이 없는 가비지 컬렉션으로 메모리 관리
- 특정 스타일의 코딩을 Go에 내장되어 있는 'go fmt' 라는 유틸리티를 통해 강력하게 제안
- 특별한 점
    - Loop는 for만 지원 ([예시](https://github.com/Tiny-Haru/Go-Study/blob/master/For%20Loop/main.go))
    - if 앞에도 for처럼 전처리 statement 지원 ([예시](https://github.com/Tiny-Haru/Go-Study/blob/master/If%20Statement/main.go))
    - 포인터를 지원하지만, 포인터에 관한 복잡성을 없애기 위해 포인터간 연산은 불가능
    - 구조체를 지원하고, 태그 개념을 통해 재사용성을 높임.
    - 초 경량 스레드를 사용하는 '고루틴'과 '채널'이라는 강력한 동시성을 지원.

## GOROOT & GOPATH
나도 처음 시작할 때 어려워했고, 많은 사람들이 어려워하는 GOROOT와 GOPATH에 대해 간단하게 알아보자.

### GOROOT
GOROOT의 개념은 JAVA의 JAVA_HOME과 같은 개념이라고 보면 된다. Go 공식 홈페이지에서 .tar.gz로 된 파일을 내려받아 유저 홈 디렉토리에 압축을 풀었다면 GOROOT는 `C:\Users\(사용자 이름)\go`가 된다. (윈도우)

정의하면, GOROOT는 GO를 설치한 위치이다.

### GOPATH
```
The GOPATH environment variable is used to specify directories outside of $GOROOT that contain the source for Go projects and their binaries.
```

대부분의 프로그래밍 언어에서 프로젝트 루트는 때에 따라 바뀐다. 하지만 Go의 경우에는 프로젝트 루트가 GOPATH/src로 고정된다. 예를 들어 'hello'라는 프로젝트를 만든다고 가정하자. 위와 같이 GOROOT를 설정해 놓은 상태라면, 이 프로젝트의 경로는 `C:\Users\(사용자 이름)\go\src\hello`가 되는 것이다.

### Go의 웹 프레임워크
파이썬에는 Django, Flask, 루비의 Ruby on Rails, Node.js에는 Express와 같이 메인 프레임워크들이 대체적으로 정해져 있는 반면에 Go는 그렇지 않다. 대표적 풀스텍 프레임워크로는 beego와 revel이 있고, 마이크로 웹 프레임워크로는 martini, iris, gin, echo가 있다.

프레임워크의 대부분이 속도를 강조한다. 하지만 Go를 사용한다는 것 자체가 성능에서 큰 이점을 가지고 가기에, 성능보다는 '활발한 활동이 이루어지는가?'라는 질문이 중요하다고 생각한다.

[martini](https://github.com/go-martini/martini)는 초반에 인기를 끌었으나, Go언어의 사상과 맞지 않는다는 [의견](https://www.reddit.com/r/golang/comments/25oxg0/three_reasons_you_should_not_use_martini/)도 존재했다.

[iris](https://github.com/kataras/iris)는 다른 Go 프레임워크들 사이에서 좋은 성능을 보여준다. 하지만 개발자의 도덕적 논란이 있었다는 [의견](http://www.florinpatan.ro/2016/10/why-you-should-not-use-iris-for-your-go.html)도 존재했다.

[gin](https://github.com/gin-gonic/gin)는 martini와 비슷한 API를 가지고 있지만, 더 빠르다고 소개하고 있다.

[echo](https://github.com/labstack/echo)는 gin보다 빠르다고 소개하고 있다. 설명만 보면 `martini < gin < echo`이다.

---
## 기타 정보들
1. Echo : Follow the Echo [Guide](https://echo.labstack.com/guide).

2. 참고 레포지토리 : Plan-B's [repository](https://github.com/JoMingyu/--Awesome-Go--).
