package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeProtectedResourceCmd = &cobra.Command{
	Use:   "describe-protected-resource",
	Short: "Returns information about a saved resource, including the last time it was backed up, its Amazon Resource Name (ARN), and the Amazon Web Services service type of the saved resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeProtectedResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeProtectedResourceCmd).Standalone()

		backup_describeProtectedResourceCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies a resource.")
		backup_describeProtectedResourceCmd.MarkFlagRequired("resource-arn")
	})
	backupCmd.AddCommand(backup_describeProtectedResourceCmd)
}
