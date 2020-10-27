FROM node
WORKDIR /src/
RUN apt update && apt install git -y
RUN git clone https://github.com/nahhoj/k8s.git
WORKDIR /src/k8s/node
RUN npm install
CMD ["node", "index"]
EXPOSE 8080