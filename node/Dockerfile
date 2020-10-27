FROM ubuntu:18.04
WORKDIR /usr/project/
RUN apt update && apt install git -y
RUN git clone https://github.com/nahhoj/k8s.git
RUN cd node && npm install
CMD ["node", "index"]
EXPOSE 8080