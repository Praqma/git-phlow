#!/bin/sh


DIR_NAME=build/github
GITHUB_FAKE=$GOPATH/src/github.com/code-cafe/git-phlow/build

#CREATE TEST DIRECTORY
cd $GOPATH/src/github.com/code-cafe/git-phlow
mkdir -p -m 0755 $DIR_NAME
cd $DIR_NAME

#CREATE TEST FILES
touch README.md file1.txt file2.txt

#DECLARE GIT FUNCTIONS
create_inital_commit () {
    git init
    git config user.email "john@doe.com"
    git config user.name "johndoe"
    git add .
    git commit -m "initial commit"
}

two_commits_on_master () {
    echo "hello" > README.md
    add_all_and_commit "1 commit - changes to readme"
    echo "file1" > file1.txt
    add_all_and_commit "2 commit - changes to file1"
    echo "#!/bin/bash  \n exit 0" > test.sh
    chmod 755 ./test.sh
    echo "#!/bin/bash  \n exit 1" > testerr.sh
    chmod 755 ./testerr.sh
    add_all_and_commit "3 commit - test script added"
}

branch_foo_additions (){
    git checkout -b  foo
    echo "hello world to README" > README.md
    echo "on branch foo" > file1.txt
    add_all_and_commit "changes on branch foo"
}


branch_master_additions () {
    git checkout master
    echo "hello world to README from master" > README.md
    add_all_and_commit "changes on branch master"
}


branch_issue_additions () {
    git checkout -b 11-issue-bar
    echo "changes in file 2" > file2.txt
    echo "changes in file 1" > file1.txt
    add_all_and_commit "changes on branch 11-issue"
}

return_to_master () {
    git checkout master
}


create_delivered_branches () {
    git checkout -b delivered/42-issue-branch
    echo "on branch delivered/42-issue-branch" > README.md
    add_all_and_commit "delivered/42 branch commit"
    return_to_master
    git checkout -b delivered/24-issue-branch
    echo "on branch delivered/24-issue-branch" > README.md
    add_all_and_commit "delivered/24 branch commit"
}

create_ready_branches () {
    git checkout -b ready/15-issue-branch
    echo "on branch ready/15-issue-branch" > README.md
    add_all_and_commit "ready/16 branch commit"

    return_to_master
    git checkout -b ready/99-issue-branch
    echo "on branch ready/99-issue-branch" > README.md
    add_all_and_commit "ready/99 branch commit"
}

add_all_and_commit (){
    git add .
    git commit -m "$1"
}

create_origin () {
    git clone $GITHUB_FAKE/github $GITHUB_FAKE/phlow-test-pkg
    cd $GITHUB_FAKE/phlow-test-pkg
    git branch bar
    git branch delivered/1-issue-branch
}


#ACTUAL SCRIPT
echo "CREATING TEST REPOSITORY"
create_inital_commit
two_commits_on_master
branch_foo_additions
branch_master_additions
branch_issue_additions
return_to_master
create_delivered_branches
return_to_master
create_ready_branches
return_to_master
create_origin
echo "WRAPPING UP JOB"
