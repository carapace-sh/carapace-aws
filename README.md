# carapace-aws

Carapace-aws is an enriched completer for [aws-cli].

![](./carapace-aws.png)

## How it works

The subcommands in [aws-cli] are mostly based on the [botocore] service definitions.
These are `json` files that contain additional information and static completions.
But these are not fully exposed to the shell completion and only accessible in [auto-prompt].

Carapace-aws parses the [botocore] service definitions into a [carapace] based completer.
This enables support for descriptions, colored hightlighting, and custom completions.
Undefined completions are simply delegated to the official [aws_completer].

## Getting Started

When using [carapace-bin] this happens implicitly when then `carapace-aws` binary is in [PATH].

Otherwise the completions can be loaded directly with [`carapace-aws _carapace`](https://carapace-sh.github.io/carapace/carapace/gen.html#hidden-subcommand).

[auto-prompt]:https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-prompting.html
[aws-cli]:https://github.com/aws/aws-cli
[aws_completer]:https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-completion.html
[botocore]:https://github.com/boto/botocore
[carapace-bin]:https://github.com/carapace-sh/carapace-bin
[carapace]:https://carapace.sh
[PATH]:https://en.wikipedia.org/wiki/PATH_(variable)
