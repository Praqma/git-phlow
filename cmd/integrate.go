package cmd

import (
	"fmt"

	"github.com/praqma/git-phlow/options"
	"github.com/praqma/git-phlow/phlow"
	"github.com/praqma/git-phlow/plugins"
	"github.com/spf13/cobra"
)

// integrateCmd represents the integrate command
var integrateCmd = &cobra.Command{
	Use:   "deliver",
	Short: "deliver work in agent version",
	Long: fmt.Sprintf(`
%s delivers the changes in an agent version. The current branch if will be merged into the default branch
with your default branch and pushed to the remote.
`, options.Bold("integrate")),
	Run: func(cmd *cobra.Command, args []string) {

		defaultBranch, _ := plugins.GetDefaultBranch(plugins.RepoURL)
		phlow.LocalDeliver(defaultBranch)
	},
}

func init() {
	agentCmd.AddCommand(integrateCmd)
}
