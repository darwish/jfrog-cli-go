package commands

import (
    "strings"
    "encoding/json"
    "github.com/JFrogDev/bintray-cli-go/utils"
)

func DownloadVersion(versionDetails *utils.VersionDetails, flags *utils.BintrayDetails) {
    path := flags.ApiUrl + "packages/" + versionDetails.Subject + "/" +
        versionDetails.Repo + "/" + versionDetails.Package + "/versions/" + versionDetails.Version + "/files"
    resp, body := utils.SendGet(path, nil, flags.User, flags.Key)
    if resp.StatusCode != 200 {
        utils.Exit(resp.Status + ". " + utils.ReadBintrayMessage(body))
    }
    var results []VersionFilesResult
    err := json.Unmarshal(body, &results)
    utils.CheckError(err)

    for _, result := range results {
        utils.DownloadBintrayFile(flags, versionDetails, result.Path)
    }
}

func CreateVersionDetailsForDownloadVersion(versionStr string) *utils.VersionDetails {
    parts := strings.Split(versionStr, "/")
    if len(parts) != 4 {
        utils.Exit("Argument format should be subject/repository/package/version. Got " + versionStr)
    }
    return utils.CreateVersionDetails(versionStr)
}

type VersionFilesResult struct {
    Path string
}