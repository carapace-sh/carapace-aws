package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteUsageLimitCmd = &cobra.Command{
	Use:   "delete-usage-limit",
	Short: "Deletes a usage limit from Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteUsageLimitCmd).Standalone()

	redshiftServerless_deleteUsageLimitCmd.Flags().String("usage-limit-id", "", "The unique identifier of the usage limit to delete.")
	redshiftServerless_deleteUsageLimitCmd.MarkFlagRequired("usage-limit-id")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteUsageLimitCmd)
}
