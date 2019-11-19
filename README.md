#새프로젝트를 만든다.
gcloud projects create [YOUR_PROJECT_ID] --set-as-default
ex) gcloud projects create heokjin-home --set-as-default

#프로젝트가 만들어졌는지 확인
gcloud projects describe [YOUR_PROJECT_ID]
ex) gcloud projects describe heokjin-home

#배포
gcloud components update

#배포구성이 변경되면 아래 명령으로 바꿔준다.
gcloud config set project heokjin-home

#그후
export GO111MODULE=on
go.mod 가 있어야한다.

gcloud app deploy