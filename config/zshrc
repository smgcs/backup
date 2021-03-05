# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH
#add

# Path to your oh-my-zsh installation.
export ZSH="/Users/sp.dong/.oh-my-zsh"

# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes

#ZSH_THEME="af-magic"
ZSH_THEME="myagnoster"
# Set list of themes to pick from when loading at random
# Setting this variable when ZSH_THEME=random will cause zsh to load
# a theme from this variable instead of looking in ~/.oh-my-zsh/themes/
# If set to an empty array, this variable will have no effect.
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion.
# Case-sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# You can set one of the optional three formats:
# "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# or set a custom format using the strftime function format specifications,
# see 'man strftime' for details.
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load?
# Standard plugins can be found in ~/.oh-my-zsh/plugins/*
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.

plugins=(git zsh-autosuggestions  zsh-syntax-highlighting)


source $ZSH/oh-my-zsh.sh
#source ~/.zsh/git-completion.zsh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/rsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"
# golang config
alias cdi="cd ~/go/src/gf/information"
alias gfinfo="cd ~/go/src/gf/information/ && vim"
alias gitdiff="git difftool -t vimdiff"
export GOPATH=/Users/sp.dong/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
#export GOROOT=/usr/local/Cellar/go
export GOROOT=/usr/local/go
export PATH="/usr/local/opt/bison/bin:$PATH:/usr/local/instantclient_12_2"
export GOPROXY=https://goproxy.cn
#export GO111MODULE=on
# oracle config
export DYLD_LIBRARY_PATH="/usr/local/instantclient_12_2"
export PKG_CONFIG_PATH="/usr/lib/pkgconfig"
export ORACLE_HOME=$DYLD_LIBRARY_PATH
export TNS_ADMIN="/usr/local/instantclient_12_2/tns"
export NLS_LANG="AMERICAN_AMERICA.UTF8"

# grafana
export PATH="$PATH:/usr/bin/grafana-5.4.2/bin"
export SSLKEYLOGFILE=/Users/sp.dong/Documents/sslkeylogout.log

export ZOOKEEPER_HOME=/Users/sp.dong/opt/zookeeper
export PATH=$PATH:$ZOOKEEPER_HOME/bin
export FZF_DEFAULT_OPTS="--height 40% --layout=reverse --preview '(highlight -O ansi {} || cat {}) 2> /dev/null | head -500'"
alias vif='vim $(fzf)'
source ~/.zsh-interactive-cd.plugin.zsh
#export http_proxy=http://10.2.237.202:7777
#export https_proxy=http://10.2.237.202:7777

[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.aliyun.com/homebrew/homebrew-bottles
alias gp='git push'
SRC_ENDPOINT=https://sourcegraph.com
SRC_ACCESS_TOKEN=70dfa220ae5a55002c819b495e65deed6b11bebb
ZSH_AUTOSUGGEST_HIGHLIGHT_STYLE='fg=6'

export PATH="$PATH:/usr/local/Cellar/qt/5.13.0/bin/"

#bindkey
bindkey '^w' vi-forward-blank-word
bindkey '^b' backward-delete-word

export PATH=$PATH:/usr/local/Cellar/mongodb-community@3.2/3.2.22/bin

#Python3虚拟环境,virtualenvwrapper.sh中50行的python改为python3
export  WORKON_HOME='~/.virtualenvs'   #子虚拟环境输入路径
source /usr/local/bin/virtualenvwrapper.sh
