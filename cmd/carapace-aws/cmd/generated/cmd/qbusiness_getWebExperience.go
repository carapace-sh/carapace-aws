package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getWebExperienceCmd = &cobra.Command{
	Use:   "get-web-experience",
	Short: "Gets information about an existing Amazon Q Business web experience.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getWebExperienceCmd).Standalone()

	qbusiness_getWebExperienceCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the web experience.")
	qbusiness_getWebExperienceCmd.Flags().String("web-experience-id", "", "The identifier of the Amazon Q Business web experience.")
	qbusiness_getWebExperienceCmd.MarkFlagRequired("application-id")
	qbusiness_getWebExperienceCmd.MarkFlagRequired("web-experience-id")
	qbusinessCmd.AddCommand(qbusiness_getWebExperienceCmd)
}
