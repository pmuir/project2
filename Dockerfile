FROM scratch
ENV foo 11
EXPOSE 8080
ENTRYPOINT ["/project2"]
COPY ./bin/ /