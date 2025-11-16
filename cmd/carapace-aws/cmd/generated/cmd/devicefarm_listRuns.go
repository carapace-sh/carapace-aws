package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listRunsCmd = &cobra.Command{
	Use:   "list-runs",
	Short: "Gets information about runs, given an AWS Device Farm project ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listRunsCmd).Standalone()

	devicefarm_listRunsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project for which you want to list runs.")
	devicefarm_listRunsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listRunsCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_listRunsCmd)
}
