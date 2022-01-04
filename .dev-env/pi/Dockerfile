FROM golang

ENV GOOS=linux \
    GOARCH=arm \
    GOARM=5 \
    SHELL=/bin/bash

ARG USER_ID=1000
ARG GROUP_ID=1000

RUN apt-get update && \
    apt-get install -y sudo tmux vim && \
    apt-get clean

RUN groupadd -g ${GROUP_ID} pi && \
    useradd -l -u ${USER_ID} -g pi pi && \
    usermod -aG sudo pi && \
    install -d -m 0755 -o pi -g pi /home/pi && \
    echo "pi ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/010_pi-nopasswd

USER pi

COPY tmux.conf /home/pi/.tmux.conf
COPY bashrc /home/pi/.bashrc

RUN echo source ~/.bashrc > ~/.bash_profile

RUN git config --global user.email "pi@pi" && \
    git config --global user.name "pi"

ENTRYPOINT ["tmux"]
