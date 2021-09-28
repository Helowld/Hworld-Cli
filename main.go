// /*
// Copyright © 2021 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
package main

import "github.com/Helowld/Hworld-Cli/cmd"

func main() {
	cmd.Execute()
}

////// gqgenc approach

// token := os.Getenv("GITHUB_TOKEN")
// ctx := context.Background()
// githubClient := &gen.Client{
// 	Client: clientv2.NewClient(http.DefaultClient, "https://api.github.com/graphql", func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
// 		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

// 		return next(ctx, req, gqlInfo, res)
// 	}),
// }
// getUser, err := githubClient.GetFile(ctx, "hello_world.c")
// if err != nil {
// 	if handledError, ok := err.(*client.ErrorResponse); ok {
// 		fmt.Fprintf(os.Stderr, "handled error: %s\n", handledError.Error())
// 	} else {
// 		fmt.Fprintf(os.Stderr, "unhandled error: %s\n", err.Error())
// 	}
// 	os.Exit(1)
// }

// fmt.Printf("getUser: %v\n", *getUser.Repository.DefaultBranchRef.Target.History.Nodes[0].File.Object.Text)
