FROM aryaha/runandeh-base:latest

ADD runandehd /runandehd
WORKDIR /

ENTRYPOINT [ "sh" ]
