# 구글 프로토콜 버퍼 사용 테스트

## 개요
구글 프로토콜 버퍼는 구글에서 개발하고 오픈소스로 공개한, 직렬화 데이터 구조 입니다. 쉽게 말해서 XML, JSON 등과 같은 데이터를 저장하는 하나의 포맷이라고 합니다. proto3부터 다양한 언어를 지원하고 직렬화 속도가 빠르고 크기도 작습니다. 하나의 파일에 최대 64M까지 지원할 수 있으며, JSON 파일로 변환이 가능합니다. 사용이 간단하고 BSD License 정책으로 100% 무료라고 합니다.

## 사용 방법
  1. Protoc 설치
  [https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases)
  2. proto 파일 작성
  데이터를 구조화 합니다. 아래 코드는 Protocol Buffer 공식 홈페이지에 있는 예제 코드입니다.
  ```
  ### test.proto 파일
  message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;
}

// Our address book file is just one of these.
message AddressBook {
  repeated Person people = 1;
}
  ```
  3. Proto Complier를 통해 정의한 언어가 자동 생성 됩니다.
  형식은
  * protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
  go 언어 코드를 생성하기 위해
  ```
  # protoc. -I=. --go_out=./tutorial ./test.proto
  ```
  4. 테스트는 임의로 데이터를 넣어 binary data를 만들고 그 data를 읽어보았습니다.

## 결과
일단 proto 파일 형식에 대해 좀 더 공부할 필요가 있습니다. 사용하는데에 어렵지 않아보입니다. 프로젝트를 하면서 모델을 저장하는 부분이 복잡하다면 사용해보려고 합니다.
