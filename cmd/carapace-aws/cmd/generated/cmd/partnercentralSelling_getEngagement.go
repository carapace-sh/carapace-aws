package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getEngagementCmd = &cobra.Command{
	Use:   "get-engagement",
	Short: "Use this action to retrieve the engagement record for a given `EngagementIdentifier`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getEngagementCmd).Standalone()

	partnercentralSelling_getEngagementCmd.Flags().String("catalog", "", "Specifies the catalog related to the engagement request.")
	partnercentralSelling_getEngagementCmd.Flags().String("identifier", "", "Specifies the identifier of the Engagement record to retrieve.")
	partnercentralSelling_getEngagementCmd.MarkFlagRequired("catalog")
	partnercentralSelling_getEngagementCmd.MarkFlagRequired("identifier")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getEngagementCmd)
}
