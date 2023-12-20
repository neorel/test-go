FROM node:20-alpine

WORKDIR /home/node

# CMD ["sleep", "1d"]
COPY . .

RUN npm install
RUN npm run build

CMD ["npm",  "run" ,"start"]