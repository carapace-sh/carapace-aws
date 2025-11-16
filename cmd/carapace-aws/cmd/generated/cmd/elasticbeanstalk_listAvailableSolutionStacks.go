package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_listAvailableSolutionStacksCmd = &cobra.Command{
	Use:   "list-available-solution-stacks",
	Short: "Returns a list of the available solution stack names, with the public version first and then in reverse chronological order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_listAvailableSolutionStacksCmd).Standalone()

	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_listAvailableSolutionStacksCmd)
}
