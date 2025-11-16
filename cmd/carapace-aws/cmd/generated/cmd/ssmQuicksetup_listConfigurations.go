package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_listConfigurationsCmd = &cobra.Command{
	Use:   "list-configurations",
	Short: "Returns configurations deployed by Quick Setup in the requesting Amazon Web Services account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_listConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_listConfigurationsCmd).Standalone()

		ssmQuicksetup_listConfigurationsCmd.Flags().String("configuration-definition-id", "", "The ID of the configuration definition.")
		ssmQuicksetup_listConfigurationsCmd.Flags().String("filters", "", "Filters the results returned by the request.")
		ssmQuicksetup_listConfigurationsCmd.Flags().String("manager-arn", "", "The ARN of the configuration manager.")
		ssmQuicksetup_listConfigurationsCmd.Flags().String("max-items", "", "Specifies the maximum number of configurations that are returned by the request.")
		ssmQuicksetup_listConfigurationsCmd.Flags().String("starting-token", "", "The token to use when requesting a specific set of items from a list.")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_listConfigurationsCmd)
}
