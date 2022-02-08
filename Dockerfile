FROM ubuntu:18.04  
WORKDIR /root/
COPY formServer ./
RUN mkdir images
EXPOSE 3000
CMD ["./formServer"]