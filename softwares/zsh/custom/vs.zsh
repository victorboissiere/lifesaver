function vs()
{
  RESULTS=`find . -type f -not -path '*/\.*' -not -path '*/vendor/*' -not -path '*/packages/*' -ipath "*$1*"`
  red() { echo -e "\033[00;31m$1\033[0m"; }
  NB_FILES=$(echo $RESULTS | wc -w)
  if [ $NB_FILES -eq "1" ]; then
    $EDITOR $(echo "$RESULTS" | head -n1 | cut -d " " -f1)
  elif [ $NB_FILES -eq "0" ]; then
    >&2 red "No matching file"
  else
    >&2 red "Error. Found more than one file"
    column <<< "$(printf '%s\n' $RESULTS)"
  fi
}

_vs() {
  local curcontext="$curcontext" state line expl

  _arguments -C \
    '*:: :->open_files'

  case "$state" in
    open_files)
      local file=${words[CURRENT]}
      compadd -U - `find . -type f -not -path '*/\.*' -not -path '*/vendor/*' -not -path '*/packages/*' -ipath "*$file*" | sed "s|^\./||"`
      compstate[insert]=menu # no expand
      ;;
  esac
  return 0
}

compdef _vs vs


