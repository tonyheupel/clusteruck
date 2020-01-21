# Our first stage, that is the Builder
FROM node:13.6-alpine AS build
WORKDIR /app
COPY "./*.js*" /app/
RUN npm install
# run tests, etc. here so it fails...

# Our Second stage, that creates an image for production
FROM node:13.6-alpine AS runtime
WORKDIR /app
COPY --from=build /app /app
RUN npm install --production
EXPOSE 8080 8081
CMD npm start