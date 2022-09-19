# Here we use multi-stage to build image.
# For more information about multi-stage builds, see official doc:
# https://docs.docker.com/develop/develop-images/multistage-build/

# Here is builder kit doc for frontend
# https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md

################################
# STAGE 0: build production code
################################
FROM node:14.20.0-alpine3.16 AS builder

WORKDIR /app

ENV NODE_ENV=production

COPY ./f2e/package*.json .
RUN npm ci

COPY ./f2e/ .
RUN npm run build

################################
# STAGE 1: build nginx image
################################
FROM nginx:1.23.1
# copy build result from STAGE 0
COPY --from=builder /app/dist/ /usr/share/nginx/html

# Since nginx 1.19, we can use config template to generate config dynamically which
# can apply environment variable, so that we can setup listening port number from
# docker-componse.yml directly.
# For more information about using environment variables in nginx configuration,
# see official document: https://hub.docker.com/_/nginx
COPY ./docker/f2e/nginx.conf /etc/nginx/nginx.conf
COPY ./docker/f2e/nginx.conf.template /etc/nginx/templates/default.conf.template

WORKDIR /app
