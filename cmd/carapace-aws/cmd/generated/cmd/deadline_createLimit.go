package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createLimitCmd = &cobra.Command{
	Use:   "create-limit",
	Short: "Creates a limit that manages the distribution of shared resources, such as floating licenses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createLimitCmd).Standalone()

		deadline_createLimitCmd.Flags().String("amount-requirement-name", "", "The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.")
		deadline_createLimitCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_createLimitCmd.Flags().String("description", "", "A description of the limit.")
		deadline_createLimitCmd.Flags().String("display-name", "", "The display name of the limit.")
		deadline_createLimitCmd.Flags().String("farm-id", "", "The farm ID of the farm that contains the limit.")
		deadline_createLimitCmd.Flags().String("max-count", "", "The maximum number of resources constrained by this limit.")
		deadline_createLimitCmd.MarkFlagRequired("amount-requirement-name")
		deadline_createLimitCmd.MarkFlagRequired("display-name")
		deadline_createLimitCmd.MarkFlagRequired("farm-id")
		deadline_createLimitCmd.MarkFlagRequired("max-count")
	})
	deadlineCmd.AddCommand(deadline_createLimitCmd)
}
