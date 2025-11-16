package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbrecommendationCmd = &cobra.Command{
	Use:   "modify-dbrecommendation",
	Short: "Updates the recommendation status and recommended action status for the specified recommendation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbrecommendationCmd).Standalone()

	rds_modifyDbrecommendationCmd.Flags().String("locale", "", "The language of the modified recommendation.")
	rds_modifyDbrecommendationCmd.Flags().String("recommendation-id", "", "The identifier of the recommendation to update.")
	rds_modifyDbrecommendationCmd.Flags().String("recommended-action-updates", "", "The list of recommended action status to update.")
	rds_modifyDbrecommendationCmd.Flags().String("status", "", "The recommendation status to update.")
	rds_modifyDbrecommendationCmd.MarkFlagRequired("recommendation-id")
	rdsCmd.AddCommand(rds_modifyDbrecommendationCmd)
}
