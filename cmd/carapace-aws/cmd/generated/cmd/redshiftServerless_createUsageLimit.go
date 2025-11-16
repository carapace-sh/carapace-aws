package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createUsageLimitCmd = &cobra.Command{
	Use:   "create-usage-limit",
	Short: "Creates a usage limit for a specified Amazon Redshift Serverless usage type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createUsageLimitCmd).Standalone()

	redshiftServerless_createUsageLimitCmd.Flags().String("amount", "", "The limit amount.")
	redshiftServerless_createUsageLimitCmd.Flags().String("breach-action", "", "The action that Amazon Redshift Serverless takes when the limit is reached.")
	redshiftServerless_createUsageLimitCmd.Flags().String("period", "", "The time period that the amount applies to.")
	redshiftServerless_createUsageLimitCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Redshift Serverless resource to create the usage limit for.")
	redshiftServerless_createUsageLimitCmd.Flags().String("usage-type", "", "The type of Amazon Redshift Serverless usage to create a usage limit for.")
	redshiftServerless_createUsageLimitCmd.MarkFlagRequired("amount")
	redshiftServerless_createUsageLimitCmd.MarkFlagRequired("resource-arn")
	redshiftServerless_createUsageLimitCmd.MarkFlagRequired("usage-type")
	redshiftServerlessCmd.AddCommand(redshiftServerless_createUsageLimitCmd)
}
