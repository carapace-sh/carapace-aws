package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_replicateInstanceCmd = &cobra.Command{
	Use:   "replicate-instance",
	Short: "Replicates an Amazon Connect instance in the specified Amazon Web Services Region and copies configuration information for Amazon Connect resources across Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_replicateInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_replicateInstanceCmd).Standalone()

		connect_replicateInstanceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_replicateInstanceCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_replicateInstanceCmd.Flags().String("replica-alias", "", "The alias for the replicated instance.")
		connect_replicateInstanceCmd.Flags().String("replica-region", "", "The Amazon Web Services Region where to replicate the Amazon Connect instance.")
		connect_replicateInstanceCmd.MarkFlagRequired("instance-id")
		connect_replicateInstanceCmd.MarkFlagRequired("replica-alias")
		connect_replicateInstanceCmd.MarkFlagRequired("replica-region")
	})
	connectCmd.AddCommand(connect_replicateInstanceCmd)
}
