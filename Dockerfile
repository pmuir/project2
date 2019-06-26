FROM scratch
EXPOSE 8080
ENTRYPOINT ["/project2"]
COPY ./bin/ /