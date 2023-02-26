package cmd

import (
	"github.com/spf13/cobra"
)

// linkedinCmd represents the linkedin command
var linkedinCmd = &cobra.Command{
	Use:   "linkedin",
	Short: "opens Arjun Danadagi's linkedin profile in your browser",
	Run: func(cmd *cobra.Command, args []string) {
        openbrowser("https://linkedin.com/in/arjundandagi")
	},
}

// twitterCmd represents the twitter command
var twitterCmd = &cobra.Command{
	Use:   "twitter",
	Short: "opens Arjun Danadagi's twitter profile in your browser",
	Run: func(cmd *cobra.Command, args []string) {
		 openbrowser("https://twitter.com/arjundandagi")
	},
}

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "opens Arjun Danadagi's github profile in your browser",
	Run: func(cmd *cobra.Command, args []string) {
		 openbrowser("https://github.com/arjundandagi")
	},
}


// githubCmd represents the github command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "opens Arjun Danadagi's all profile's in your browser",
	Run: func(cmd *cobra.Command, args []string) {

	    //TODO: learn how to reuse the above defined funcs to invoke?
	    // check if its possible
	    openbrowser("https://github.com/arjundandagi")
	    openbrowser("https://linkedin.com/in/arjundandagi")
	    openbrowser("https://twitter.com/arjundandagi")

	},
}

func init() {
    rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(linkedinCmd)
    rootCmd.AddCommand(twitterCmd)
    rootCmd.AddCommand(githubCmd)
}
