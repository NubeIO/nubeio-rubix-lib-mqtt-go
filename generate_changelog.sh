set -eu

# destination of the final changelog file
OUTPUT_FILE=CHANGELOG.md

# generate the changelog
git --no-pager log --no-merges --format="### %s%d%n>%aD%n%n>Author: %aN %n%n " > $OUTPUT_FILE

# prevent recursion!
# since a 'commit --amend' will trigger the post-commit script again
# we have to check if the changelog file has changed or not
res=$(git status --porcelain | grep -c $OUTPUT_FILE)
if [ "$res" -gt 0 ]; then
  git add $OUTPUT_FILE
  git commit --amend
  echo "Populated Changelog in $OUTPUT_FILE"
fi