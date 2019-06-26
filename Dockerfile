FROM scratch
ENV foo 10
EXPOSE 8080
ENTRYPOINT ["/project2"]
COPY ./bin/ /