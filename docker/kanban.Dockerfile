FROM alpine:latest

COPY  kanban /bin/kanban

RUN adduser -D -u 1000 kanban
USER kanban

CMD /bin/kanban
