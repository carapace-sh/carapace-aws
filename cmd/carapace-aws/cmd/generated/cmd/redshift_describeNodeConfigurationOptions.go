package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeNodeConfigurationOptionsCmd = &cobra.Command{
	Use:   "describe-node-configuration-options",
	Short: "Returns properties of possible node configurations such as node type, number of nodes, and disk usage for the specified action type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeNodeConfigurationOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeNodeConfigurationOptionsCmd).Standalone()

		redshift_describeNodeConfigurationOptionsCmd.Flags().String("action-type", "", "The action type to evaluate for possible node configurations.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to evaluate for possible node configurations.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("filters", "", "A set of name, operator, and value items to filter the results.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("owner-account", "", "The Amazon Web Services account used to create or copy the snapshot.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot associated with the message to describe node configuration.")
		redshift_describeNodeConfigurationOptionsCmd.Flags().String("snapshot-identifier", "", "The identifier of the snapshot to evaluate for possible node configurations.")
		redshift_describeNodeConfigurationOptionsCmd.MarkFlagRequired("action-type")
	})
	redshiftCmd.AddCommand(redshift_describeNodeConfigurationOptionsCmd)
}
