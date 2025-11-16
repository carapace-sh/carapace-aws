package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateUsageLimitCmd = &cobra.Command{
	Use:   "update-usage-limit",
	Short: "Update a usage limit in Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateUsageLimitCmd).Standalone()

	redshiftServerless_updateUsageLimitCmd.Flags().String("amount", "", "The new limit amount.")
	redshiftServerless_updateUsageLimitCmd.Flags().String("breach-action", "", "The new action that Amazon Redshift Serverless takes when the limit is reached.")
	redshiftServerless_updateUsageLimitCmd.Flags().String("usage-limit-id", "", "The identifier of the usage limit to update.")
	redshiftServerless_updateUsageLimitCmd.MarkFlagRequired("usage-limit-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateUsageLimitCmd)
}
