## How to start

```
$ npm run start
```

## Works to do

- Angular 비동기 화면 갱신에 대한 이해

## Problem

- PWA 알림

  ### AI 채팅 기억력
  - 장기 기억 메모리 요약본 DB 저장하여 동봉해서 API 발송

  ### 웹소켓 JWT 검증
  - 방법1) 쿼리 이용/ 방법2) 연결 후 메시지 이용 => 이 방법 사용
  - JWT 만료 확인
    - JWT 토큰 검증 시 유효할 경우 만료시간을 저장하여 웹소켓 close 타이머 설정

## 작업 완료

- UI 개선(form 에러 inline 표시, room 생성 form validation, 모바일 UI, CHAT 크기 좁은 문제, 마크다운 CHAT 개선, unread 메시지, 날짜 표시)
- ollama 스케줄러 timeout 에러 발생함 => ollama 응답 30초->3분 대기 변경
- 에러처리(toastr 라이브러리로 표출. commonError 타입 지정해 커스텀 에러메시지 작성. ws 연결은 에러처리x, 재연결 시도. ws연결시 토큰 에러 발생시 close.)
- Room 도 websock 실시간 업데이트, 상단 이동
- 웹소켓 close 되면 새로고침 해야하는 문제 -> reconnect 예약기능
- AI 선톡 기능(랜덤으로 AI에게 요청보내고 답변만 사용자에게 송출)
- UI 개선(스크롤 하단 고정)
- 룸 (소프트)삭제
- 로그아웃(토큰 삭제, WEBSOCKET COLSE)
- 멀티 Room 리스트
- 채팅방 구현
- HTTP JWT 검증
- 웹소켓 JWT 검증
- AI 연결

  ### AI 채팅 기억력
  - ROOM 에 저장된 AI system 동봉해서 api 발송
  - 최근 대화 10개 동봉해서 API 발송
  - 시스템(고정 명령어)를 사용자가 편집할 수 있도록 함.
  - AI 채팅 기억력(장기기억 요약본 생성기능 보류)

  ### 랜덤 선톡
  - 2~6 시간 랜덤 예약 선톡
  - 마지막 대화가 선톡 예약시간 1시간 이내이면 선톡을 보내지 않고 예약시간을 다시 설정한다.
