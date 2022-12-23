package main

import (
	"memo/writer"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		// 현재 날짜의 메모 추가
		// 마지막 라인에 적은만큼 추가
		// 마지막 라인에 추가 전 그만큼 공백작성
		writer.DefaultDateWriter()
	} else {
		println("옵션 표시하기")
		println("혹은 키워드를 기반으로 메모작성하기")
		println("혹은 어떤 키워드가있는지 검색하기")
	}

}

// 옵션 표시하기
// 현재 있는 메모의 뒤에 추가하기 (공백, 시간, 내용)
// 메모 병합하기
// 키워드를 기반으로 메모 작성하기
// 어떤 키워드가 있는지 확인하기
// 상세적으로 조작하게 하기 (어떤 메모가 있는지 정확하게 파악하기)
// 메모 일부 수정
// 메모 보기
// - 인자를 넣으면 인터페이스