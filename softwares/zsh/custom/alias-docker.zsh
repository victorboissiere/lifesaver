# ------------------------------------
# Docker functions
# ------------------------------------

docker_alias_stop_all_containers() { docker stop $(docker ps -a -q); }
docker_alias_remove_all_containers() { docker rm $(docker ps -a -q); }
docker_alias_remove_all_empty_images() {docker images | awk '{print $2 " " $3}' | grep '^<none>' | awk '{print $2}' | xargs -I{} docker rmi {}; }
docker_alias_docker_file_build() { docker build -t=$1 .; }
docker_alias_show_all_docker_related_alias() { alias | grep 'docker' | sed "s/^\([^=]*\)=\(.*\)/\1 => \2/"| sed "s/['|\']//g" | sort; }
docker_alias_bash_into_running_container() { docker exec -it $(docker ps -aqf "name=$1") bash; }

# ------------------------------------
# Docker alias
# ------------------------------------

# Stop all containers
alias dstop='docker_alias_stop_all_containers'

# Remove all containers
alias drm='docker_alias_remove_all_containers'

# Dockerfile build, e.g., $dbu tcnksm/test
alias dbu='docker_alias_docker_file_build'

# Show all alias related docker
alias dalias='docker_alias_show_all_docker_related_alias'

# Bash into running container
alias dbash='docker_alias_bash_into_running_container'

# Get images
alias di="docker images"

# Get container IP
alias dip="docker inspect --format '{{ .NetworkSettings.IPAddress }}'"

# Remove all images
alias drmimages='docker rmi $(docker images -a -q)'

alias dc='docker-compose'
alias dprod='docker-compose -f docker-compose.prod.yml -f docker-compose.yml'
alias ddeploy='docker-compose up -d --no-deps --build'

