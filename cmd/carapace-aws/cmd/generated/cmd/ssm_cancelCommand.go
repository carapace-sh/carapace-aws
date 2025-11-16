package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_cancelCommandCmd = &cobra.Command{
	Use:   "cancel-command",
	Short: "Attempts to cancel the command specified by the Command ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_cancelCommandCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_cancelCommandCmd).Standalone()

		ssm_cancelCommandCmd.Flags().String("command-id", "", "The ID of the command you want to cancel.")
		ssm_cancelCommandCmd.Flags().String("instance-ids", "", "(Optional) A list of managed node IDs on which you want to cancel the command.")
		ssm_cancelCommandCmd.MarkFlagRequired("command-id")
	})
	ssmCmd.AddCommand(ssm_cancelCommandCmd)
}
