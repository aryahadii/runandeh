FROM ubuntu:16.04

RUN apt-get update -y

# Install packages
RUN apt-get install -y curl build-essential

# Remove apt cache to make the image smaller
RUN rm -rf /var/lib/apt/lists/*

# Add runandeh
ADD runandehd /home/app/runandehd
WORKDIR /home/app

ENTRYPOINT [ "bash" ]
