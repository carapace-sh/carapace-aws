package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackEventsCmd = &cobra.Command{
	Use:   "describe-stack-events",
	Short: "Returns all stack related events for a specified stack in reverse chronological order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackEventsCmd).Standalone()

	cloudformation_describeStackEventsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_describeStackEventsCmd.Flags().String("stack-name", "", "The name or the unique stack ID that's associated with the stack, which aren't always interchangeable:")
	cloudformation_describeStackEventsCmd.MarkFlagRequired("stack-name")
	cloudformationCmd.AddCommand(cloudformation_describeStackEventsCmd)
}
