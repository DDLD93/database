FROM ubuntu:18.04  
WORKDIR /root/
COPY formServer ./
RUN mkdir images
EXPOSE 5000
CMD ["./formServer"]