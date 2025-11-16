package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createFarmCmd = &cobra.Command{
	Use:   "create-farm",
	Short: "Creates a farm to allow space for queues and fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createFarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_createFarmCmd).Standalone()

		deadline_createFarmCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_createFarmCmd.Flags().String("description", "", "The description of the farm.")
		deadline_createFarmCmd.Flags().String("display-name", "", "The display name of the farm.")
		deadline_createFarmCmd.Flags().String("kms-key-arn", "", "The ARN of the KMS key to use on the farm.")
		deadline_createFarmCmd.Flags().String("tags", "", "The tags to add to your farm.")
		deadline_createFarmCmd.MarkFlagRequired("display-name")
	})
	deadlineCmd.AddCommand(deadline_createFarmCmd)
}
