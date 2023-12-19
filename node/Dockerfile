FROM node:20-alpine

WORKDIR /home/node

COPY . .

RUN npm install
RUN npm run build

CMD ["npm",  "run" ,"start"]