package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_createEngagementCmd = &cobra.Command{
	Use:   "create-engagement",
	Short: "The `CreateEngagement` action allows you to create an `Engagement`, which serves as a collaborative space between different parties such as AWS Partners and AWS Sellers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_createEngagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_createEngagementCmd).Standalone()

		partnercentralSelling_createEngagementCmd.Flags().String("catalog", "", "The `CreateEngagementRequest$Catalog` parameter specifies the catalog related to the engagement.")
		partnercentralSelling_createEngagementCmd.Flags().String("client-token", "", "The `CreateEngagementRequest$ClientToken` parameter specifies a unique, case-sensitive identifier to ensure that the request is handled exactly once.")
		partnercentralSelling_createEngagementCmd.Flags().String("contexts", "", "The `Contexts` field is a required array of objects, with a maximum of 5 contexts allowed, specifying detailed information about customer projects associated with the Engagement.")
		partnercentralSelling_createEngagementCmd.Flags().String("description", "", "Provides a description of the `Engagement`.")
		partnercentralSelling_createEngagementCmd.Flags().String("title", "", "Specifies the title of the `Engagement`.")
		partnercentralSelling_createEngagementCmd.MarkFlagRequired("catalog")
		partnercentralSelling_createEngagementCmd.MarkFlagRequired("client-token")
		partnercentralSelling_createEngagementCmd.MarkFlagRequired("description")
		partnercentralSelling_createEngagementCmd.MarkFlagRequired("title")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_createEngagementCmd)
}
