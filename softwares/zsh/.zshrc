#-- ZSH --#
###########

# Path to your oh-my-zsh installation.
export ZSH=~/.oh-my-zsh

# Disable composer xdebug warning for performance issue
export COMPOSER_DISABLE_XDEBUG_WARN=false

ZSH_THEME="robbyrussell"

plugins=(git zsh-syntax-highlighting yarn docker)

source $ZSH/oh-my-zsh.sh

#-- GLOBAL --#
##############

# Set vim editor for git, ...
export EDITOR="vim"

alias python='python3'
alias pip='pip3'

function mkcd()
{
  mkdir $1 && cd $1
}

export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8


#-- WEBSITES --#
################

#-- DEVOPS --#

alias dc='docker-compose'
alias dr='docker'
alias dprod='docker-compose -f docker-compose.prod.yml -f docker-compose.yml'
alias ddeploy='docker-compose up -d --no-deps --build'
alias drmc='docker rm `docker ps -q -f status=exited`'
alias drmi='docker rmi $(docker images -f "dangling=true" -q)'

alias ad='ansible-doc'
alias ap='ansible-playbook'
alias an='ansible'
alias vg='vagrant'

#-- LARAVEL --#
###############

# Php artisan (Laravel)
alias pa="php artisan"


#-- GIT --#
###########

# Git aliases
alias g='git'
alias gs="git status"
alias ga="git add"
alias gaa="git add --all"
alias gc="git commit -m "
alias gp="git push"
alias gpom="git push origin master"
alias gpo="git push origin"
alias gpuom="git pull origin master"
alias gpuo="git pull origin"
alias gch="git checkout"
alias gchb="git checkout -b"
alias gl="git log --oneline"
alias gf="git fetch"
alias gm="git merge"
alias nah="git reset --hard HEAD"

# Open git in browser
function go() {
  giturl=$(git config --get remote.origin.url)
  if [ "$giturl" = "" ]; then
     echo "Not a git repository or no remote.origin.url set"
     return
  fi

  giturl=${giturl/git\@github\.com\:/https://github.com/}
  giturl=${giturl/\.git/\/tree/}
  branch="$(git symbolic-ref HEAD 2>/dev/null)" ||
  branch="(unnamed branch)"     # detached HEAD
  branch=${branch##refs/heads/}
  giturl=$giturl$branch
  "Opening your git.."
  google-chrome $giturl > /dev/null 2>&1
}

#-- CONFIG FILES --#
####################

# Shortcut to config files
alias zrc='vim ~/.zshrc'
alias vimrc='vim ~/.vimrc'
alias i3config='vim ~/.i3/config'

# Source ZSH config file
alias szrc='source ~/.zshrc'

#-- NAVIGATION --#
##################

# No need for cd
alias ~='cd ~'

#-- UTILS --#
#############

# Sudo shortcut
alias s='sudo'

# Become root
alias root='sudo -i'

# Google-chrome
alias chrome='google-chrome'

# Number of commits
alias nbc='git log --oneline | wc -l'

# Color managment
alias grep='grep --color=auto'
alias fgrep='fgrep --color=auto'
alias egrep='egrep --color=auto'


#-- COMPOSER --#
################

alias c='composer'

#-- NODEJS --#
##############

alias y='yarn'

#-- FUNCTIONS --#
#################

function man()
{
  OUTPUT="`tldr $1 2>&1`"
  if [ $? -eq 0 ]; then
    echo $OUTPUT
  else
    /usr/bin/man $1
  fi
}

# Help to be able to use alias in vim
function zshalias()
{
  grep "^alias" ~/.zshrc > ~/.zshenv
}

function refreshconfig()
{
  wget -O - https://ls.gitcommit.fr | sudo bash -s $1 $USER
}

#-- FUN --#
###########

function notifyme()
{

  cmd=$@
  $@;
  terminal-notifier -title 'Done!' -message "$cmd" -activate 'com.apple.Terminal' -sound Ping
}

