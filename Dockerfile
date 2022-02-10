FROM alpine 
WORKDIR /app
COPY formServer ./
RUN mkdir images
EXPOSE 3000
CMD ["./formServer"]