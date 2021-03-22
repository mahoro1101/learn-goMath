package elementary

//초등수학 정수 도형제외

func Add(a int, b int) int {
	return a + b
}

// add 더하다

func Sub(a int, b int) int {
	return a - b
}

// substract 빼다

func Multi(a int, b int) int {
	return a * b
}

// multiply 곱하다

func Div(a int, b int) int {
	return a / b
}

// divide 나누다

func Mod(a int, b int) int {
	return a % b
}

// modulo 나머지 연산하다

func Divisor(a int) []int {
	var divisor []int
	for i := 1; i <= a; i++ {
		if float64(a/i) == float64(a)/float64(i) {
			divisor = append(divisor, i)
		}
	}
	return divisor
}

// divisor a값의 약수 반환

func Multiple(a int) []int {
	var multiple []int
	for i := 1; i < 11; i++ {
		multiple = append(multiple, a*i)
	}
	return multiple
}

// multiple 10개의 배수 반환

func GreatestCommonDivisor(a int, b int) int {
	da := Divisor(a)
	db := Divisor(b)
	var commonDivisor int
	for _, v1 := range db {
		for _, v2 := range da {
			if v1 == v2 {
				commonDivisor = v1
				break
			}
		}
	}
	return commonDivisor
}

//최대 공약수

func LeastCommonMultiple(a int, b int) int {
	ma := Multiple(a)
	mb := Multiple(b)
	var commonMultiple int

	for _, v1 := range mb {
		for _, v2 := range ma {
			if v1 == v2 {
				commonMultiple = v1
				return commonMultiple
			}
		}
	}
	return -1
}

//최소 공배수

func Inequality(a int, b int) string {
	switch {
	case a > b:
		return "a>b"
	case a < b:
		return "a<b"
	default:
		return "a=b"
	}
}

//이하 이상 미만 초과

//올림 버림 반올림 math패키지 math.Ceil(a) math.Floor(a) math.Round(a)
//소수
//분수랑 가능성
//평균
//나눗셈을 곱셈으로
//분수 나눗셈
//소수 나눗셈
//비율, 백분율
//비례식
//원주율
//솟수, 합성수, 거듭제곱
//소인수분해
//최대공약수
//최소공배수
//양수,음수, 정수, 유리수
//절댓값
//부등호
//문자사용
