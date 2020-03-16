package phlow

import (
	"fmt"
	"sort"

	"strings"

	"github.com/code-cafe/git-phlow/executor"
	"github.com/code-cafe/git-phlow/githandler"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/setting"
)

//UpNext ...
//Returns the next branch ready for integration based on time of creation
//Oldest branches gets integrated first.
func UpNext(prefix string) (name string) {
	git := githandler.Git{Run: executor.RunGit}
	conf := setting.NewProjectStg("phlow")
	if prefix == "" {
		prefix = "ready/"
	}

	out, err := git.Branch("-a")
	if err != nil {
		fmt.Println(err)
		return
	}

	branches := githandler.Ready(githandler.AsList(out), conf.Remote, conf.DeliveryBranchPrefix)

	if len(branches) != 0 {
		if options.GlobalFlagHumanReadable {
			fmt.Println("Found 'ready/' branches on remote")
		}

		name = getNextBranch(branches, conf.Remote)
		return
	}

	if options.GlobalFlagHumanReadable {
		fmt.Println("No 'ready/' branches found on remote")
	}
	return ""
}

//getNextBranch
//Sort branches and returns the oldest ready branch
func getNextBranch(branches []string, origin string) string {
	m := make(map[int]string)
	var time int
	var err error

	//Create map with time and branch name
	for _, br := range branches {
		if time, err = githandler.BranchTime(br); err == nil {
			m[time] = br

			if options.GlobalFlagHumanReadable {
				fmt.Printf("%s : %d \n", br, time)
			}
		}
	}

	//Order the keys in a separate list
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	if len(keys) > 0 {
		res := removeRemoteFromUpNext(m[keys[0]], origin)
		return res
	}
	return ""
}

//remoteRemoteFromUpNext ...
func removeRemoteFromUpNext(name string, origin string) string {

	if strings.HasPrefix(name, origin+"/") {
		name = strings.TrimPrefix(name, origin+"/")
		return name
	}
	return name
}
