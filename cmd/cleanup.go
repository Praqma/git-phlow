package cmd

import (
	"fmt"

	"github.com/code-cafe/git-phlow/cmd/cmdperm"
	"github.com/code-cafe/git-phlow/flags"
	"github.com/code-cafe/git-phlow/options"
	"github.com/code-cafe/git-phlow/phlow"
	"github.com/code-cafe/git-phlow/ui"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var cleanCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "cleanup removes all delivered branches",
	Long: fmt.Sprintf(`
%s is for tidying up the git workspace. 
As you follow the workflow a lot of branches prefixed with 'delivered/' will be in the workspace, and should just be deleted if they have been successfully integrated. 
Running the command will remove these branches locally and remote as well. A local version of this command is available, and will just remove the local branches.
Some branches can not be deleted, because git cannot detect if they have been integrated into the integration branch. That can be due to a rebase or squash. Those can be deleted with the 'force' flag. 
`, ui.Format.Bold("cleanup")),
	PreRun: func(cmd *cobra.Command, args []string) {
		cmdperm.RequiredCurDirRepository()
	},
	Run: func(cmd *cobra.Command, args []string) {
		phlow.CleanCaller(options.GlobalFlagTarget)
	},
}

func init() {
	RootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().BoolVar(&flags.CleanUpTidy, "tidy", false, "prune remote branches")

	cleanCmd.Flags().BoolVarP(&flags.CleanupForce, "force", "f", false, "similar to git branch -D to delete an unmerged branch")

	cleanCmd.Flags().BoolVarP(&flags.CleanUpDryRun, "dry-run", "n", false, "output files that will be deleted in non dry run")

	cleanCmd.Flags().BoolVar(&flags.CleanupDelivered, "delivered", false, "remove delivered branches. No prompt")

}
