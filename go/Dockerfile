FROM golang
WORKDIR /src/
RUN apt update && apt install git -y
RUN git clone https://github.com/nahhoj/k8s.git
WORKDIR /src/k8s/go
RUN go build index.go
CMD ./index
EXPOSE 3000