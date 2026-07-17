FROM squidfunk/mkdocs-material:9.7.7

COPY action.sh /action.sh

RUN apk add --no-cache bash \
  && chmod +x /action.sh

ENTRYPOINT ["/action.sh"]
