package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteWebExperienceCmd = &cobra.Command{
	Use:   "delete-web-experience",
	Short: "Deletes an Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteWebExperienceCmd).Standalone()

	qbusiness_deleteWebExperienceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the Amazon Q Business web experience.")
	qbusiness_deleteWebExperienceCmd.Flags().String("web-experience-id", "", "The identifier of the Amazon Q Business web experience being deleted.")
	qbusiness_deleteWebExperienceCmd.MarkFlagRequired("application-id")
	qbusiness_deleteWebExperienceCmd.MarkFlagRequired("web-experience-id")
	qbusinessCmd.AddCommand(qbusiness_deleteWebExperienceCmd)
}
