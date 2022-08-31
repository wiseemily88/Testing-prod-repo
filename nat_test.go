func TestNetworkLoopbackNat(t *testing.T) {
	skip.If(t, testEnv.GitHubActions) //https://github.com/wiseemily88/Testing-prod-repo/issues/9
	skip.If(t, testEnv.OSType == "windows", "FIXME")
	skip.If(t, testEnv.IsRemoteDaemon)
