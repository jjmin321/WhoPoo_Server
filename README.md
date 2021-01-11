# WhoPoo_Server

WHOPOO는 리그오브레전드 API를 사용하여 데이터를 정리하고 보여주는 서비스입니다.

## Stack
|                      | Android     | Window | Server        | 
|:--------------------:|:---------------:|:------------------:|:--------------------:|
| Developer | 남규석 | 김경훈 | 제정민 | 
| Develop Language | Kotlin | WPF | GO| 
| Develop Tool     | Android Studio  | Visual Studio | Visual Studio Code | 

## API

## /matches 
    # 특정 게임 ID를 입력받아 해당 게임에 대한 정보를 반환해줌 

    - Parameter 
        - gameId (Required)
    - Return
        - MatchDTO
            - TeamsDTO
            - ParticipantsDTO
            - ParticipantIdentitiesDTO

## /matchlists
    # 유저 이름을 입력받아 최근 게임 전적을 반환해줌

    - Parameter
        - name (Required)
    - Query Parameter
        - startIndex (Optional)
        - endIndex (Optional)  
    - Return 
        - MatchlistsDTO