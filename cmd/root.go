// Copyright © 2020 Patrick Lawrence <patrick.lawrence@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "clean-docker-images",
		Short: "Remove older docker images",
		Run:   rootRun,
	}
	rootCmdFlags = struct {
		endpoint   string
		apiVersion string
		numToKeep  int
		remove     bool
	}{}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.Flags().StringVarP(&rootCmdFlags.endpoint, "endpoint", "e", "", "The docker endpoint tcp://...:2375")
	rootCmd.Flags().StringVarP(&rootCmdFlags.apiVersion, "api-version", "a", "1.40", "The minimum docker API version")
	rootCmd.Flags().IntVarP(&rootCmdFlags.numToKeep, "num-to-keep", "n", 2, "Number of images of the same type to keep")
	rootCmd.Flags().BoolVarP(&rootCmdFlags.remove, "remove", "r", false, "Actually remove older images")
}

// ByCreated implements sort.Interface for []types.ImageSummary based on
// the Created field (newest first)
type ByCreated []types.ImageSummary

func (b ByCreated) Len() int           { return len(b) }
func (b ByCreated) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByCreated) Less(i, j int) bool { return b[i].Created > b[j].Created }

func rootRun(cobraCmd *cobra.Command, args []string) {
	ctx := context.Background()
	cli, err := client.NewClient(rootCmdFlags.endpoint, rootCmdFlags.apiVersion, nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}

	imageMap := make(map[string][]types.ImageSummary)

	for _, image := range images {
		for _, repoDigest := range image.RepoDigests {
			k := strings.Split(repoDigest, ":")[0]
			imageMap[k] = append(imageMap[k], image)
		}
	}

	for _, v := range imageMap {
		sort.Sort(ByCreated(v))
	}

	for k, v := range imageMap {
		if len(v) > rootCmdFlags.numToKeep {
			for _, image := range v[rootCmdFlags.numToKeep:] {
				fmt.Printf("Can remove image %s for %s\n", image.ID, k)
				if rootCmdFlags.remove {
					fmt.Printf("Removing image %s\n", image.ID)
					_, err := cli.ImageRemove(ctx, image.ID, types.ImageRemoveOptions{})
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error removing image: %v\n", err)
					}
				}
			}
		}
	}
}
