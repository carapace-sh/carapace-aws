package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_deleteGraphCmd = &cobra.Command{
	Use:   "delete-graph",
	Short: "Disables the specified behavior graph and queues it to be deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_deleteGraphCmd).Standalone()

	detective_deleteGraphCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph to disable.")
	detective_deleteGraphCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_deleteGraphCmd)
}
