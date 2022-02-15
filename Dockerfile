FROM ubuntu:jammy
WORKDIR /app
COPY form ./
RUN mkdir images
EXPOSE 3000
CMD ["./form"]
