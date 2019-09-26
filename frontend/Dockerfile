FROM node:10.16.2-alpine

WORKDIR /app

RUN apk update && \
    npm install -Sg npm @vue/cli --no-optional

# npm installに必要なjsonファイルを用意
ADD package*.json ./

# npm run build時にvueファイルが必要になる。
# packaege.jsonを上書きしたくないのでnpm installより先に記載。
ADD . .


RUN npm install axios -S --no-optional
RUN npm install --no-optional

