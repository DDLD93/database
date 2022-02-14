FROM ubuntu:jammy
WORKDIR /app/
COPY forms ./
RUN mkdir images
EXPOSE 3000
CMD ["./forms"]
