package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deregisterManagedInstanceCmd = &cobra.Command{
	Use:   "deregister-managed-instance",
	Short: "Removes the server or virtual machine from the list of registered servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deregisterManagedInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deregisterManagedInstanceCmd).Standalone()

		ssm_deregisterManagedInstanceCmd.Flags().String("instance-id", "", "The ID assigned to the managed node when you registered it using the activation process.")
		ssm_deregisterManagedInstanceCmd.MarkFlagRequired("instance-id")
	})
	ssmCmd.AddCommand(ssm_deregisterManagedInstanceCmd)
}
