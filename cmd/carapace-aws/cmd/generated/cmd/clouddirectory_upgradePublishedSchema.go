package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_upgradePublishedSchemaCmd = &cobra.Command{
	Use:   "upgrade-published-schema",
	Short: "Upgrades a published schema under a new minor version revision using the current contents of `DevelopmentSchemaArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_upgradePublishedSchemaCmd).Standalone()

	clouddirectory_upgradePublishedSchemaCmd.Flags().String("development-schema-arn", "", "The ARN of the development schema with the changes used for the upgrade.")
	clouddirectory_upgradePublishedSchemaCmd.Flags().String("dry-run", "", "Used for testing whether the Development schema provided is backwards compatible, or not, with the publish schema provided by the user to be upgraded.")
	clouddirectory_upgradePublishedSchemaCmd.Flags().String("minor-version", "", "Identifies the minor version of the published schema that will be created.")
	clouddirectory_upgradePublishedSchemaCmd.Flags().String("published-schema-arn", "", "The ARN of the published schema to be upgraded.")
	clouddirectory_upgradePublishedSchemaCmd.MarkFlagRequired("development-schema-arn")
	clouddirectory_upgradePublishedSchemaCmd.MarkFlagRequired("minor-version")
	clouddirectory_upgradePublishedSchemaCmd.MarkFlagRequired("published-schema-arn")
	clouddirectoryCmd.AddCommand(clouddirectory_upgradePublishedSchemaCmd)
}
