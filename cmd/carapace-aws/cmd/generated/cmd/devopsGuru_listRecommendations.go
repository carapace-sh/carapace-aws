package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listRecommendationsCmd = &cobra.Command{
	Use:   "list-recommendations",
	Short: "Returns a list of a specified insight's recommendations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listRecommendationsCmd).Standalone()

	devopsGuru_listRecommendationsCmd.Flags().String("account-id", "", "The ID of the Amazon Web Services account.")
	devopsGuru_listRecommendationsCmd.Flags().String("insight-id", "", "The ID of the requested insight.")
	devopsGuru_listRecommendationsCmd.Flags().String("locale", "", "A locale that specifies the language to use for recommendations.")
	devopsGuru_listRecommendationsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_listRecommendationsCmd.MarkFlagRequired("insight-id")
	devopsGuruCmd.AddCommand(devopsGuru_listRecommendationsCmd)
}
