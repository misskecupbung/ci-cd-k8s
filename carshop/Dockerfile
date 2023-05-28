FROM alpine:latest
 
# install node
RUN apk update && \
    apk add --no-cache nodejs npm

WORKDIR /app

# copy project file
COPY package.json .

# install node packages
RUN npm set progress=false && \
    npm config set depth 0 && \
    npm install

# copy app files
COPY . .

# application server port
EXPOSE 3001

# default run command
CMD npm run start
