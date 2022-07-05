require 'git'
require 'logger'
require 'active_support/all'

days = 0
g = Git.open(ENV['PWD'], :log => Logger.new(STDOUT))

g.log(14).each do |commit|
  actual = commit.author.date
  author_name = "luka"
  github_name = "R0bertDenir0"
  author_mail = "luka.michelson@gmail.com"
  

  if rand(4) == 2
    author_name = "Yolan Maldonado"
    github_name = "Y0lan"
    author_mail = "register@yolan.dev"
  end
  puts(author_name)
  puts(author_mail)
  puts(github_name)
  puts("-------------")
  %x(git filter-branch -f --commit-filter '
    if [ "$GIT_COMMIT" = "#{commit.sha}" ]
    then
        export GIT_AUTHOR_NAME="#{author_name}"
        export GIT_AUTHOR_EMAIL="#{author_mail}"
        export GIT_COMMITTER_NAME="#{github_name}"
        export GIT_COMMITTER_EMAIL="#{author_mail}"
        git commit-tree "$@";
    else
        git commit-tree "$@";
    fi' HEAD)
end
