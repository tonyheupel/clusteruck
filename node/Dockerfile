# Our first stage, that is the Builder
FROM node:13.6-alpine AS build
# install Python for nod-gyp to stop throwing errors
RUN apk update && apk add python make && rm -rf /var/cache/apk/*
WORKDIR /app
COPY "./*.js*" /app/
RUN npm install
# run tests, etc. here so it fails...

# Our Second stage, that creates an image for production
FROM node:13.6-alpine AS runtime
WORKDIR /app
COPY --from=build /app /app
RUN npm install --production
EXPOSE 8081
CMD npm start
