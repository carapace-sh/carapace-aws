package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateCrawlerCmd = &cobra.Command{
	Use:   "update-crawler",
	Short: "Updates a crawler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateCrawlerCmd).Standalone()

	glue_updateCrawlerCmd.Flags().String("classifiers", "", "A list of custom classifiers that the user has registered.")
	glue_updateCrawlerCmd.Flags().String("configuration", "", "Crawler configuration information.")
	glue_updateCrawlerCmd.Flags().String("crawler-security-configuration", "", "The name of the `SecurityConfiguration` structure to be used by this crawler.")
	glue_updateCrawlerCmd.Flags().String("database-name", "", "The Glue database where results are stored, such as: `arn:aws:daylight:us-east-1::database/sometable/*`.")
	glue_updateCrawlerCmd.Flags().String("description", "", "A description of the new crawler.")
	glue_updateCrawlerCmd.Flags().String("lake-formation-configuration", "", "Specifies Lake Formation configuration settings for the crawler.")
	glue_updateCrawlerCmd.Flags().String("lineage-configuration", "", "Specifies data lineage configuration settings for the crawler.")
	glue_updateCrawlerCmd.Flags().String("name", "", "Name of the new crawler.")
	glue_updateCrawlerCmd.Flags().String("recrawl-policy", "", "A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.")
	glue_updateCrawlerCmd.Flags().String("role", "", "The IAM role or Amazon Resource Name (ARN) of an IAM role that is used by the new crawler to access customer resources.")
	glue_updateCrawlerCmd.Flags().String("schedule", "", "A `cron` expression used to specify the schedule (see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ?")
	glue_updateCrawlerCmd.Flags().String("schema-change-policy", "", "The policy for the crawler's update and deletion behavior.")
	glue_updateCrawlerCmd.Flags().String("table-prefix", "", "The table prefix used for catalog tables that are created.")
	glue_updateCrawlerCmd.Flags().String("targets", "", "A list of targets to crawl.")
	glue_updateCrawlerCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_updateCrawlerCmd)
}
