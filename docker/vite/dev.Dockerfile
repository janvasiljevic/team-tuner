FROM node:16-alpine3.11

WORKDIR /app
COPY . .

CMD [ "yarn", "run", "dev" ]