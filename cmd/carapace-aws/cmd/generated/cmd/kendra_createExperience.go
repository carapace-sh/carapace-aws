package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createExperienceCmd = &cobra.Command{
	Use:   "create-experience",
	Short: "Creates an Amazon Kendra experience such as a search application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createExperienceCmd).Standalone()

	kendra_createExperienceCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create your Amazon Kendra experience.")
	kendra_createExperienceCmd.Flags().String("configuration", "", "Configuration information for your Amazon Kendra experience.")
	kendra_createExperienceCmd.Flags().String("description", "", "A description for your Amazon Kendra experience.")
	kendra_createExperienceCmd.Flags().String("index-id", "", "The identifier of the index for your Amazon Kendra experience.")
	kendra_createExperienceCmd.Flags().String("name", "", "A name for your Amazon Kendra experience.")
	kendra_createExperienceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access `Query` API, `GetQuerySuggestions` API, and other required APIs.")
	kendra_createExperienceCmd.MarkFlagRequired("index-id")
	kendra_createExperienceCmd.MarkFlagRequired("name")
	kendraCmd.AddCommand(kendra_createExperienceCmd)
}
