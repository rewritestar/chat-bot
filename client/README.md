## How to start

```
$ npm run start
```

## Works to do

- 채팅방 구현
- AI 연결
- 에러처리
- UI 개선
- AI 기억력 개선

## Problem

- AI 가 선톡을 할때, UI 를 어떻게 업데이트 시키는가? => http 통신은 요청-응답 구조임
  => checkOrigin 강제 true 수정하기
  => 웹소켓 코드 구조 개선하기
  => 채팅창 마지막 채팅 스크롤 고정하기

## 웹소켓 JWT 검증

- 방법1) 쿼리 이용/ 방법2) 연결 후 메시지 이용 => 이 방법 사용
- JWT 만료 확인
  - JWT 토큰 검증 시 유효할 경우 만료시간을 저장하여 웹소켓 close 타이머 설정
