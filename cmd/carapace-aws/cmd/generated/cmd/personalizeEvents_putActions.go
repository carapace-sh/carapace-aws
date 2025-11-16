package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEvents_putActionsCmd = &cobra.Command{
	Use:   "put-actions",
	Short: "Adds one or more actions to an Actions dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEvents_putActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalizeEvents_putActionsCmd).Standalone()

		personalizeEvents_putActionsCmd.Flags().String("actions", "", "A list of action data.")
		personalizeEvents_putActionsCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the Actions dataset you are adding the action or actions to.")
		personalizeEvents_putActionsCmd.MarkFlagRequired("actions")
		personalizeEvents_putActionsCmd.MarkFlagRequired("dataset-arn")
	})
	personalizeEventsCmd.AddCommand(personalizeEvents_putActionsCmd)
}
