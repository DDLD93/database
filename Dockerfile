FROM ubuntu:jammy
WORKDIR /root/
COPY formServer ./
RUN mkdir images
EXPOSE 3000
CMD ["./formServer"]