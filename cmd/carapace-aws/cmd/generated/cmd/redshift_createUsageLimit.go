package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createUsageLimitCmd = &cobra.Command{
	Use:   "create-usage-limit",
	Short: "Creates a usage limit for a specified Amazon Redshift feature on a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createUsageLimitCmd).Standalone()

	redshift_createUsageLimitCmd.Flags().String("amount", "", "The limit amount.")
	redshift_createUsageLimitCmd.Flags().String("breach-action", "", "The action that Amazon Redshift takes when the limit is reached.")
	redshift_createUsageLimitCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster that you want to limit usage.")
	redshift_createUsageLimitCmd.Flags().String("feature-type", "", "The Amazon Redshift feature that you want to limit.")
	redshift_createUsageLimitCmd.Flags().String("limit-type", "", "The type of limit.")
	redshift_createUsageLimitCmd.Flags().String("period", "", "The time period that the amount applies to.")
	redshift_createUsageLimitCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createUsageLimitCmd.MarkFlagRequired("amount")
	redshift_createUsageLimitCmd.MarkFlagRequired("cluster-identifier")
	redshift_createUsageLimitCmd.MarkFlagRequired("feature-type")
	redshift_createUsageLimitCmd.MarkFlagRequired("limit-type")
	redshiftCmd.AddCommand(redshift_createUsageLimitCmd)
}
