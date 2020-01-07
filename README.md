# Golang으로 작성하는 솔플 backend

## 기능 구현 목록
1. User & Auth
    - [ ] User는 id, name, username, email, password, profile(photo)를 가짐
    - [ ] id는 자동으로 부여되는 고유한 식별번호
    - [ ] 나머지 필드는 추후 변경 가능해야함
    - [ ] username, email은 unique
    - [ ] password는 암호화돼서 db에 저장됨
    - [ ] 로그인은 username/password로 이뤄지며, 성공 시 token을 제공
    - [ ] 로그인/회원가입을 제외한 다른 api는 token으로 인증을 해야 사용 가능
    - [ ] token은 24시간동안 유지되고, 이후 재발급 받아야함