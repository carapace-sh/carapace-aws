package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listUniqueProblemsCmd = &cobra.Command{
	Use:   "list-unique-problems",
	Short: "Gets information about unique problems, such as exceptions or crashes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listUniqueProblemsCmd).Standalone()

	devicefarm_listUniqueProblemsCmd.Flags().String("arn", "", "The unique problems' ARNs.")
	devicefarm_listUniqueProblemsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listUniqueProblemsCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_listUniqueProblemsCmd)
}
