package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeServiceIntegrationCmd = &cobra.Command{
	Use:   "describe-service-integration",
	Short: "Returns the integration status of services that are integrated with DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeServiceIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeServiceIntegrationCmd).Standalone()

	})
	devopsGuruCmd.AddCommand(devopsGuru_describeServiceIntegrationCmd)
}
