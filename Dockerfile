FROM ubuntu 
WORKDIR /app
COPY formServer ./
RUN mkdir images
EXPOSE 3000
CMD ["./forms"]
