FROM alpine:20210804

ARG USER=developer

ENV FZF_DEFAULT_COMMAND="rg --hidden --files -g '!{.git/*}'"
ENV TERM=xterm-256color

RUN apk add \
    sudo \
    tmux \
    vim \
    git \
    openssh \
    curl \
    ctags \
    fzf \
    ripgrep

RUN adduser -D ${USER} ${USER} && \
    echo "${USER} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/010_${USER}-nopasswd

USER ${USER}

COPY --chown=${USER}:${USER} tmux.conf /home/${USER}/.tmux.conf
COPY --chown=${USER}:${USER} profile /home/${USER}/.profile
COPY --chown=${USER}:${USER} vimrc /home/${USER}/.vim/vimrc

RUN git config --global user.email "${USER}@${USER}" && \
    git config --global user.name "${USER}"

RUN curl -fLo /home/${USER}/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

RUN vim +PlugInstall +qa

RUN sed -i 's/1E1E1E/000000/g' /home/${USER}/.vim/plugged/vim-code-dark/colors/codedark.vim

RUN sudo apk add go python3 nodejs

RUN go get golang.org/x/tools/cmd/goimports@latest && \
    go get golang.org/x/tools/gopls@latest && \
    go get github.com/go-delve/delve/cmd/dlv@latest

RUN ~/.vim/plugged/vimspector/install_gadget.py --enable-go

COPY --chown=${USER}:${USER} coc-settings.json /home/${USER}/.vim/coc-settings.json

ENV PATH=$PATH:/home/${USER}/go/bin

ENTRYPOINT ["tmux", "-u"]

