package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotasCmd = &cobra.Command{
	Use:   "service-quotas",
	Short: "With Service Quotas, you can view and manage your quotas easily as your Amazon Web Services workloads grow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotasCmd).Standalone()

	})
	rootCmd.AddCommand(serviceQuotasCmd)
}
