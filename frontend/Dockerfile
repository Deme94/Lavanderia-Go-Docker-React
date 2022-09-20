# REACT APP
# Specify a base image
FROM node:16.14.2-alpine

WORKDIR /react-app

# Install dependencies
COPY package*.json ./

RUN npm install

COPY . .

# Default command
CMD [ "npm", "start" ]