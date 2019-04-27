FROM node:alpine
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY captcha-proxy.js /usr/src/app/captcha-proxy.js
EXPOSE 5000

# Serve the app
CMD ["node", "captcha-proxy.js"]