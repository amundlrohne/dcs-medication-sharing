FROM node:16-alpine 
LABEL org.opencontainers.image.source https://github.com/amundlrohne/dcs-medication-sharing


WORKDIR /app

COPY package*.json ./
 
RUN npm install
RUN mkdir -p node_modules/.cache && chmod -R 777 node_modules/.cache


COPY . .
 
EXPOSE 3000

CMD npm start