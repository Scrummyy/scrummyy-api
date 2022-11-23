package logger_messages

const (
	InfoSuccessViewModeFetch      = "successfully retrieved view modes"
	InfoSuccessContentFetch       = "successfully retrieved contents"
	InfoSuccessContentMetaFetch   = "successfully retrieved the content-metadata"
	InfoSuccessContentSearchFetch = "successfully retrieved search contents"
	InfoSuccessfulCampaignCreated = "successfully added campaign: %s"
	InfoSuccessCampaignList       = "successfully retrieved campaigns"
	InfoSuccessCampaignDetails    = "successfully retrieved campaign details"
	InfoSuccessUpdateCampaign     = "successfully updated the campaign: %s"
	InfoSuccessPatchCampaign      = "successfully patched the campaign: %s"
	InfoSuccessCloneCampaign      = "successfully cloned the campaign: %s"
	InfoSuccessDeleteCampaign     = "successfully deleted the campaign: %s"
	InfoSuccessDeleteCampaigns    = "successfully deleted the campaigns of account id: %s"
	InfoSuccessMergedCampaigns    = "successfully merged campaigns"
	InfoAuthenticationSkipped     = "authentication is excluded on the requested path."
	InfoAuthorizationSkipped      = "authorization is excluded on the requested path."

	InfoSuccessRuleDetail  = "successfully retrieved rule details"
	InfoSuccessRuleList    = "successfully retrieved rules"
	InfoSuccessDeleteRule  = "successfully deleted the rule: %s"
	InfoSuccessDeleteRules = "successfully deleted the rules"

	InfoSuccessPageDetail  = "successfully retrieved page details"
	InfoSuccessDeploySites = "successfully deployed site"
	InfoSuccessAddSlot     = "successfully added slot %s"

	InfoSuccessSitesAnalytics = "successfully fetched sites-analytics for accountID: %s"

	InfoDecisionPreviewFlag = "request has preview flag set, return dummy visitor info"
	InfoNoSlotFound         = "no slot found"

	NoRecordsFound = "no records found"
)
